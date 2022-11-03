package frontExam

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"strconv"
	"time"
)

type ExamService struct {
}

var ojService = service.ServiceGroupApp.OjServiceServiceGroup

func (examService *ExamService) FindExamPlans(teachClassId uint) (examPlans []teachplan.ExamPlan, err error) {
	err = global.GVA_DB.Where("teach_class_id = ?", teachClassId).Find(&examPlans).Error
	return
}

func (examService *ExamService) GetExamPapers(examComing request.ExamComing) (examPaper response.ExamPaperResponse, err error) {
	var studentPaper []examManage.ExamStudentPaper
	err = global.GVA_DB.Where("student_id = ? and plan_id = ?", examComing.StudentId, examComing.PlanId).Find(&studentPaper).Error
	var choiceCount, judgeCount, blankCount, programCount uint
	for i := 0; i < len(studentPaper); i++ {
		if *studentPaper[i].QuestionType == questionType.MULTIPLE_CHOICE {
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Where("id = ?", studentPaper[i].QuestionId).Find(&examPaper.ChoiceComponent[choiceCount].Choice).Error
			if err != nil {
				return
			}
			choiceCount++
			examPaper.ChoiceComponent[i].MergeId = studentPaper[i].ID
		} else if *studentPaper[i].QuestionType == questionType.JUDGE {
			err = global.GVA_DB.Table("les_questionBank_judge").Where("id = ?", studentPaper[i].QuestionId).Find(&examPaper.JudgeComponent[judgeCount].Judge).Error
			if err != nil {
				return
			}
			judgeCount++
			examPaper.JudgeComponent[i].MergeId = studentPaper[i].ID
		} else if *studentPaper[i].QuestionType == questionType.SUPPLY_BLANK {
			err = global.GVA_DB.Table("les_questionBank_supply_blank").Where("id = ?", studentPaper[i].QuestionId).Find(&examPaper.BlankComponent[blankCount].MergeId).Error
			if err != nil {
				return
			}
			blankCount++
			examPaper.BlankComponent[i].MergeId = studentPaper[i].ID
		} else if *studentPaper[i].QuestionType == questionType.PROGRAM {
			err = global.GVA_DB.Table("les_questionBank_programm").Where("id = ?", studentPaper[i].QuestionId).Find(&examPaper.ProgramComponent[programCount].Program).Error
			if err != nil {
				return
			}
			programCount++
			examPaper.ProgramComponent[i].MergeId = studentPaper[i].ID
		}
	}
	var PaperId int64
	err = global.GVA_DB.Table("exam_student_paper").Select("paper_id").Where("student_id = ? and plan_id =?", examComing.StudentId, examComing.PlanId).Scan(&PaperId).Error
	//PaperId, err := examService.GetStudentPaperId(examComing)
	if err != nil {
		return
	}
	examPaper.PaperId = uint(PaperId)
	return
}
func (examService *ExamService) GetStudentPaperId(examComing request.ExamComing) (Id int64, err error) {
	err = global.GVA_DB.Table("exam_student_paper").Select("paper_id").Where("student_id = ? and plan_id =?", examComing.StudentId, examComing.PlanId).First(&Id).Error
	return
}

func (examService *ExamService) CommitExamPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	var optionCommit = examPaperCommit.MultipleChoiceCommit
	var JudgeCommit = examPaperCommit.JudgeCommit
	var BlankCommit = examPaperCommit.BlankCommit
	for j := 0; j < len(optionCommit); j++ {
		answers := utils.IntArrayToString(optionCommit[j].Answers)
		err = global.GVA_DB.Table("exam_student_paper").Select("answer").
			Where("id = ?", optionCommit[j].MergeId).
			Updates(examManage.ExamStudentPaper{Answer: answers}).
			Error
		if err != nil {
			return
		}
	}
	for j := 0; j < len(JudgeCommit); j++ {
		err = global.GVA_DB.Table("exam_student_paper").Select("answer").
			Where("id = ?", BlankCommit[j].MergeId).
			Updates(examManage.ExamStudentPaper{Answer: JudgeCommit[j].Answer}).
			Error
		if err != nil {
			return
		}
	}
	for j := 0; j < len(JudgeCommit); j++ {
		err = global.GVA_DB.Table("exam_student_paper").Select("answer").
			Where("id = ?", JudgeCommit[j].MergeId).
			Updates(examManage.ExamStudentPaper{Answer: BlankCommit[j].Answer}).
			Error
		if err != nil {
			return
		}
	}
	go func() {
		time.Sleep(time.Minute * 15)
		examService.ExecPapers(examPaperCommit)
	}()
	return
}

func (examService *ExamService) ExecPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	for i := 0; i < len(examPaperCommit.JudgeCommit); i++ {
		s, _ := strconv.ParseBool(examPaperCommit.JudgeCommit[0].Answer)
		if Bool, err := ojService.JudgeService.Check(examPaperCommit.JudgeCommit[0].QuestionId, s); err != nil {
			return err
		} else {
			if Bool == true {
				err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ?", examPaperCommit.JudgeCommit[0].MergeId).Error
				if err != nil {
					return err
				}
			}
		}
	}

	for i := 0; i < len(examPaperCommit.MultipleChoiceCommit); i++ {
		if Bool, err := ojService.MultipleChoiceService.Check(examPaperCommit.MultipleChoiceCommit[0].QuestionId, examPaperCommit.MultipleChoiceCommit[0].Answers); err != nil {
			return err
		} else {
			if Bool == true {
				err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ?", examPaperCommit.MultipleChoiceCommit[0].MergeId).Error
				if err != nil {
					return err
				}
			}
		}
	}
	//此处判题接口应该修改参数[]string为string
	//for i := 0; i < len(examPaperCommit.MultipleChoiceCommit); i++ {
	//	if Bool, err := ojService.SupplyBlankService.Check(examPaperCommit.MultipleChoiceCommit[0].QuestionId, examPaperCommit.JudgeCommit[0].Answer); err != nil {
	//		return
	//	} else {
	//		if Bool == true {
	//			err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ?", examPaperCommit.JudgeCommit[0].MergeId).Error
	//			if err != nil {
	//				return
	//			}
	//		}
	//	}
	//}
	return
}
