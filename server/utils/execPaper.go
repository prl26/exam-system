package utils

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/examType"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service/questionBank/oj"
	"gorm.io/gorm"
)

var ojService oj.OjService

func ExecPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	//判断题处理
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(examPaperCommit.JudgeCommit); i++ {
			if Bool, err := ojService.JudgeService.Check(examPaperCommit.JudgeCommit[i].QuestionId, examPaperCommit.JudgeCommit[i].Answer); err != nil {
				return err
			} else {
				if Bool == true {
					var result examManage.ExamStudentPaper
					err = tx.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score where id = ?", examPaperCommit.JudgeCommit[i].MergeId).Scan(&result).Error
					if err != nil {
						return err
					}
				}
			}
		}
		//选择题处理
		for i := 0; i < len(examPaperCommit.MultipleChoiceCommit); i++ {
			if Bool, err := ojService.MultipleChoiceService.Check(examPaperCommit.MultipleChoiceCommit[i].QuestionId, examPaperCommit.MultipleChoiceCommit[i].Answer); err != nil {
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
			if _, num, err := ojService.SupplyBlankService.Check(examPaperCommit.BlankCommit[i].QuestionId, examPaperCommit.BlankCommit[i].Answer); err != nil {
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
		fmt.Println("进入统分")
		var PlanDetail teachplan.ExamPlan
		tx.Model(teachplan.ExamPlan{}).Where("id =?", examPaperCommit.PlanId).Find(&PlanDetail)
		if *PlanDetail.Type == examType.FinalExam {
			tx.Raw("UPDATE tea_score as s SET s.exam_score = (SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ?),s.final_exam_name = ?,s.final_exam_id = ? "+
				"WHERE s.teach_class_id = ? and student_id = ? ", examPaperCommit.StudentId, examPaperCommit.PlanId, PlanDetail.Name, PlanDetail.ID, PlanDetail.TeachClassId, examPaperCommit.StudentId)
		} else if *PlanDetail.Type == examType.ProceduralExam {
			tx.Raw("UPDATE tea_score as s SET s.procedure_score = s.procedure_score+(SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ?)"+
				"WHERE s.teach_class_id = ? and student_id = ? ", examPaperCommit.StudentId, examPaperCommit.PlanId, PlanDetail.TeachClassId, examPaperCommit.StudentId)
		}
		var sum int
		tx.Raw("SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ?", examPaperCommit.StudentId, examPaperCommit.PlanId).Scan(&sum)
		var detail examManage.Detail
		tx.Raw("SELECT b.`name` ,c.`name` from bas_term as b,bas_lesson as c where b.id = 1 and  c.id =1 ").Scan(&detail)
		tx.Create(&examManage.ExamScore{
			StudentId:  &examPaperCommit.StudentId,
			PlanId:     &PlanDetail.ID,
			Name:       PlanDetail.Name,
			TermId:     PlanDetail.TermId,
			TermName:   detail.TermName,
			CourseId:   PlanDetail.CourseId,
			CourseName: detail.CourseName,
			Score:      &sum,
			ExamType:   PlanDetail.Type,
			StartTime:  PlanDetail.StartTime,
		})
		return nil
	})
	return
}

func ExecProgram(program examManage.CommitProgram, score uint) (err error) {
	var result examManage.ExamStudentPaper
	if score != 0 {
		err = global.GVA_DB.Raw(fmt.Sprintf("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score*100/%d where id = ?", score), program.MergeId).Scan(&result).Error
	}
	return err
}
