package utils

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/examType"
	"github.com/prl26/exam-system/server/model/questionBank/enum/languageType"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service/questionBank/oj"
	"gorm.io/gorm"
)

var ojService = oj.OjService{}

func ExecPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	//判断题处理
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(examPaperCommit.JudgeCommit); i++ {
			if Bool, err := ojService.JudgeService.ExamCheck(examPaperCommit.JudgeCommit[i].QuestionId, examPaperCommit.JudgeCommit[i].Answer); err != nil {
				return err
			} else {
				if Bool == true {
					var result examManage.ExamStudentPaper
					err = tx.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score where id = ?", examPaperCommit.JudgeCommit[i].MergeId).Scan(result).Error
					if err != nil {
						return err
					}
				}
			}
		}
		//选择题处理
		for i := 0; i < len(examPaperCommit.MultipleChoiceCommit); i++ {
			if Bool, err := ojService.MultipleChoiceService.ExamCheck(examPaperCommit.MultipleChoiceCommit[i].QuestionId, examPaperCommit.MultipleChoiceCommit[i].Answer); err != nil {
				return err
			} else {
				if Bool == true {
					var result examManage.ExamStudentPaper
					err = tx.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ?", examPaperCommit.MultipleChoiceCommit[i].MergeId).Scan(&result).Error
					if err != nil {
						return err
					}
				}
			}
		}
		//填空题处理
		for i := 0; i < len(examPaperCommit.BlankCommit); i++ {
			if _, num, err := ojService.SupplyBlankService.ExamCheck(examPaperCommit.BlankCommit[i].QuestionId, examPaperCommit.BlankCommit[i].Answer); err != nil {
				return err
			} else {
				if num != 0 {
					var result examManage.ExamStudentPaper
					err = tx.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score*"+fmt.Sprintf("%f", float64(num)/100.0)+" where id = ?", examPaperCommit.BlankCommit[i].MergeId).Scan(&result).Error
					if err != nil {
						return err
					}
				}
			}
		}
		//总分
		global.GVA_LOG.Info("进入统分")
		var sum examManage.AllScore
		tx.Raw("SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ?", examPaperCommit.StudentId, examPaperCommit.PlanId).Scan(&sum.Score)
		var PlanDetail teachplan.ExamPlan
		tx.Model(teachplan.ExamPlan{}).Where("id =?", examPaperCommit.PlanId).Find(&PlanDetail)
		planId := int(PlanDetail.ID)
		if PlanDetail.Type == examType.FinalExam {
			global.GVA_LOG.Info("期末统分")
			tx.Model(teachplan.Score{}).Where("student_id = ? and teach_class_id = ?", examPaperCommit.StudentId, PlanDetail.TeachClassId).Updates(teachplan.Score{
				ExamScrore:    &sum.Score,
				FinalExamName: PlanDetail.Name,
				FinalExamId:   &planId,
			})
		} else if PlanDetail.Type == examType.ProceduralExam {
			global.GVA_LOG.Info("过程化统分")
			global.GVA_DB.Raw("UPDATE tea_score SET procedure_score = procedure_score+procedure_proportion/100*?)", sum.Score).Where("student_id = ? and teach_class_id = ?", examPaperCommit.StudentId, PlanDetail.TeachClassId)
		}
		err = tx.Model(examManage.ExamScore{}).Where("student_id = ? and plan_id = ?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("score", sum.Score).Error
		if err != nil {
			return err
		}
		var recordId uint
		var recoreMerge []examManage.ExamRecordMerge
		tx.Model(examManage.ExamRecord{}).Select("id").Where("student_id =? and plan_id =?", examPaperCommit.StudentId, examPaperCommit.PlanId).Order("created_at desc").First(&recordId)
		err = tx.Raw("INSERT INTO exam_record_merge(created_at,updated_at,paper_id,question_id,student_id,answer,plan_id,score,question_type,problem_type,got_score,record_id) SELECT created_at,updated_at,paper_id,question_id,student_id,answer,plan_id,score,question_type,problem_type,got_score,"+fmt.Sprintf("%d", recordId)+" FROM exam_student_paper WHERE student_id = ? AND plan_id = ? and deleted_at is null ", examPaperCommit.StudentId, examPaperCommit.PlanId).Scan(&recoreMerge).Error
		if err != nil {
			return err
		}
		//err = tx.Model(examManage.ExamRecordMerge{}).Where("student_id =? and plan_id =?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("record_id", recordId).Error
		//CreateExamScore(PlanDetail,sum,examPaperCommit.StudentId)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// 试卷重新批阅
func ReExecPapers(sp teachplan.CoverRq) (err error) {
	var examPaperCommit examManage.ReExecExamPaper
	examPaperCommit.StudentId = sp.StudentId
	examPaperCommit.PlanId = sp.PlanId

	choiceType := uint(questionType.SINGLE_CHOICE)
	judgeType := uint(questionType.JUDGE)
	blankType := uint(questionType.SUPPLY_BLANK)
	//examManage.
	programType := uint(questionType.PROGRAM)
	//判断题处理
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		global.GVA_DB.Model(examManage.ExamStudentPaper{}).Where("student_id = ? and plan_id = ?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("got_score", 0)
		tx.Model(examManage.ExamStudentPaper{}).Where("student_id = ? and plan_id = ? and question_type = ?", sp.StudentId, sp.PlanId, judgeType).Find(&examPaperCommit.JudgeCommit)
		for i := 0; i < len(examPaperCommit.JudgeCommit); i++ {
			if Bool, err := ojService.JudgeService.ExamCheck(examPaperCommit.JudgeCommit[i].QuestionId, examPaperCommit.JudgeCommit[i].Answer); err != nil {
				return err
			} else {
				if Bool == true {
					var result examManage.ExamStudentPaper
					err = tx.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score where id = ? and deleted_at is null", examPaperCommit.JudgeCommit[i].Id).Scan(&result).Error
					if err != nil {
						return err
					}
				}
			}
		}
		tx.Model(examManage.ExamStudentPaper{}).Where("student_id = ? and plan_id = ? and question_type = ?", sp.StudentId, sp.PlanId, choiceType).Find(&examPaperCommit.MultipleChoiceCommit)
		//选择题处理
		for i := 0; i < len(examPaperCommit.MultipleChoiceCommit); i++ {
			answer := StringToStringArray(examPaperCommit.MultipleChoiceCommit[i].Answer, ",")
			if Bool, err := ojService.MultipleChoiceService.ExamCheck(examPaperCommit.MultipleChoiceCommit[i].QuestionId, answer); err != nil {
				return err
			} else {
				if Bool == true {
					var result examManage.ExamStudentPaper
					err = tx.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ? and deleted_at is null", examPaperCommit.MultipleChoiceCommit[i].Id).Scan(&result).Error
					if err != nil {
						return err
					}
				} else {
					var result examManage.ExamStudentPaper
					err = tx.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = 0  where id = ? and deleted_at is null", examPaperCommit.MultipleChoiceCommit[i].Id).Scan(&result).Error
					if err != nil {
						return err
					}
				}
			}
		}
		//填空题处理
		tx.Model(examManage.ExamStudentPaper{}).Where("student_id = ? and plan_id = ? and question_type = ?", sp.StudentId, sp.PlanId, blankType).Find(&examPaperCommit.BlankCommit)
		for i := 0; i < len(examPaperCommit.BlankCommit); i++ {
			answer := StringToStringArray(examPaperCommit.BlankCommit[i].Answer, ",")
			if len(answer) != 0 {
				if _, num, err := ojService.SupplyBlankService.ExamCheck(examPaperCommit.BlankCommit[i].QuestionId, answer); err != nil {
					return err
				} else {
					if num != 0 {
						var result examManage.ExamStudentPaper
						sql := "UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score*" + fmt.Sprintf("%f", float64(num)/100.0) + " where id = ? and deleted_at is null"
						err = tx.Raw(sql, examPaperCommit.BlankCommit[i].Id).Scan(&result).Error
						if err != nil {
							return err
						}
					}
				}
			}
		}

		// 编程题的处理
		tx.Model(examManage.ExamStudentPaper{}).Where("student_id = ? and plan_id = ? and question_type = ?", sp.StudentId, sp.PlanId, programType).Find(&examPaperCommit.ProgramCommit)
		for i := 0; i < len(examPaperCommit.ProgramCommit); i++ {
			if examPaperCommit.ProgramCommit[i].Answer != "" {
				answer := examManage.ProgramAnswer{}
				answer.Decode(examPaperCommit.ProgramCommit[i].Answer)
				l := languageType.LanguageType(0)
				l.ToLanguageId(answer.LanguageType)
				_, score, _, err := ojService.ProgramService.CheckProgram(examPaperCommit.ProgramCommit[i].QuestionId, answer.Code, l)
				if err != nil {
					global.GVA_LOG.Error("ReExecPapers:" + err.Error())
					continue
				}
				if score != 0 {
					var result examManage.ExamStudentPaper
					err = tx.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score*"+fmt.Sprintf("%f", float64(score)/100.0)+" where id = ? and deleted_at is null", examPaperCommit.ProgramCommit[i].Id).Scan(&result).Error
					if err != nil {
						return err
					}
				}
			}
		}
		//总分
		global.GVA_LOG.Info("进入统分")
		var sum examManage.AllScore
		tx.Raw("SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ? and deleted_at is null", examPaperCommit.StudentId, examPaperCommit.PlanId).Scan(&sum.Score)
		var PlanDetail teachplan.ExamPlan
		tx.Model(teachplan.ExamPlan{}).Where("id =?", examPaperCommit.PlanId).Find(&PlanDetail)
		planId := int(PlanDetail.ID)
		if PlanDetail.Type == examType.FinalExam {
			global.GVA_LOG.Info("期末统分")
			tx.Model(teachplan.Score{}).Where("student_id = ? and teach_class_id = ?", examPaperCommit.StudentId, PlanDetail.TeachClassId).Updates(teachplan.Score{
				ExamScrore:    &sum.Score,
				FinalExamName: PlanDetail.Name,
				FinalExamId:   &planId,
			})
		} else if PlanDetail.Type == examType.ProceduralExam {
			global.GVA_LOG.Info("过程化统分")
			global.GVA_DB.Raw("UPDATE tea_score SET procedure_score = procedure_score+procedure_proportion/100*?)", sum.Score).Where("student_id = ? and teach_class_id = ? and deleted_at is null", examPaperCommit.StudentId, PlanDetail.TeachClassId)
		}
		err = tx.Model(examManage.ExamScore{}).Where("student_id = ? and plan_id = ?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("score", sum.Score).Error
		if err != nil {
			return err
		}
		return nil
	})
	return
}
func ExecProgram(program examManage.CommitProgram, score uint) (err error) {
	var result examManage.ExamStudentPaper
	if score != 0 {
		err = global.GVA_DB.Raw(fmt.Sprintf("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score*"+fmt.Sprintf("%f", float64(score)/100.0)+" where id = ? and deleted_at is null"), program.MergeId).Scan(&result).Error
	}
	var sum float64
	global.GVA_DB.Raw("SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ? and deleted_at is null", program.StudentId, program.PlanId).Scan(&sum)
	err = global.GVA_DB.Model(examManage.ExamScore{}).Where("student_id = ? and plan_id = ?", program.StudentId, program.PlanId).Update("score", sum).Error
	return err
}
func CreateExamScore(PlanDetail teachplan.ExamPlan, sum float64, studentId uint) (examScore examManage.ExamScore, err error) {
	var num int64
	err = global.GVA_DB.Model(examManage.ExamScore{}).Where("student_id = ? and plan_id = ?", studentId, PlanDetail.ID).Find(&examScore).Count(&num).Error
	if err != nil {
		return
	} else if num == 0 {
		global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			var term basicdata.Term
			var lesson basicdata.Lesson
			tx.Model(&basicdata.Term{}).Where("id = ?", PlanDetail.TermId).Find(&term)
			tx.Model(&basicdata.Lesson{}).Where("id = ?", PlanDetail.LessonId).Find(&lesson)
			examScore = examManage.ExamScore{
				StudentId:  &studentId,
				PlanId:     &PlanDetail.ID,
				Name:       PlanDetail.Name,
				TermId:     PlanDetail.TermId,
				TermName:   term.Name,
				LessonId:   PlanDetail.LessonId,
				CourseName: lesson.Name,
				Score:      &sum,
				ExamType:   &PlanDetail.Type,
				StartTime:  PlanDetail.StartTime,
				Weight:     PlanDetail.Weight,
				IsReport:   false,
			}
			tx.Create(&examScore)
			return nil
		})
	}
	return
}
