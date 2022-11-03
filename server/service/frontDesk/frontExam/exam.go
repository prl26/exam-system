package frontExam

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/utils"
)

type ExamService struct {
}

func (examService *ExamService) FindExamPlans(teachClassId uint) (examPlans []teachplan.ExamPlan, err error) {
	err = global.GVA_DB.Where("teach_class_id = ?", teachClassId).Find(&examPlans).Error
	return
}

func (examService *ExamService) GetExamPapers(examComing request.ExamComing) (examPaper response.ExamPaperResponse, err error) {
	var studentPaper []examManage.ExamStudentPaper
	err = global.GVA_DB.Where("student_id = ? and plan_id = ?", examComing.StudentId, examComing.PlanId).Find(&studentPaper).Error
	for i := 0; i < len(studentPaper); i++ {
		if *studentPaper[i].QuestionType == questionType.MULTIPLE_CHOICE {
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Where("id = ?", studentPaper[i].QuestionId).Find(&examPaper.ChoiceComponent).Error
			if err != nil {
				return
			}
		} else if *studentPaper[i].QuestionType == questionType.JUDGE {
			err = global.GVA_DB.Table("les_questionBank_judge").Where("id = ?", studentPaper[i].QuestionId).Find(&examPaper.JudgeComponent).Error
			if err != nil {
				return
			}
		} else if *studentPaper[i].QuestionType == questionType.SUPPLY_BLANK {
			err = global.GVA_DB.Table("les_questionBank_supply_blank").Where("id = ?", studentPaper[i].QuestionId).Find(&examPaper.BlankComponent).Error
			if err != nil {
				return
			}
		} else if *studentPaper[i].QuestionType == questionType.PROGRAM {
			err = global.GVA_DB.Table("les_questionBank_programm").Where("id = ?", studentPaper[i].QuestionId).Find(&examPaper.ProgramComponent).Error
			if err != nil {
				return
			}
		}
	}
	PaperId, err := examService.GetStudentPaperId(examComing)
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
			Where("student_id = ? and paper_id = ? and plan_id = ? and question_id = ?", examPaperCommit.StudentId, examPaperCommit.PaperId, examPaperCommit.PlanId, optionCommit[j].QuestionId).
			Updates(examManage.ExamStudentPaper{Answer: answers}).
			Error
		if err != nil {
			return
		}
	}
	for j := 0; j < len(JudgeCommit); j++ {
		err = global.GVA_DB.Table("exam_student_paper").Select("answer").
			Where("student_id = ? and paper_id = ? and plan_id = ? and question_id = ?", examPaperCommit.StudentId, examPaperCommit.PaperId, examPaperCommit.PlanId, JudgeCommit[j].QuestionId).
			Updates(examManage.ExamStudentPaper{Answer: JudgeCommit[j].Answer}).
			Error
		if err != nil {
			return
		}
	}
	for j := 0; j < len(JudgeCommit); j++ {
		err = global.GVA_DB.Table("exam_student_paper").Select("answer").
			Where("student_id = ? and paper_id = ? and plan_id = ? and question_id = ?", examPaperCommit.StudentId, examPaperCommit.PaperId, examPaperCommit.PlanId, BlankCommit[j].QuestionId).
			Updates(examManage.ExamStudentPaper{Answer: BlankCommit[j].Answer}).
			Error
		if err != nil {
			return
		}
	}
	return
}
