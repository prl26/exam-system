package examManage

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	"time"
)

type TargetExamPaperService struct {
}

func (targetExamService *TargetExamPaperService) GetTargetExamPapers(examComing request.ExamComing) (examPaper response.TargetExamPaperResponse, status examManage.StudentPaperStatus, err error) {
	examPaper.TargetComponent = make([]response.TargetComponent, 0)
	var studentPaper []examManage.ExamStudentPaper
	err = global.GVA_DB.Where("student_id = ? and plan_id = ?", examComing.StudentId, examComing.PlanId).Find(&studentPaper).Error
	var targetCount uint
	for i := 0; i < len(studentPaper); i++ {
		if *studentPaper[i].QuestionType == questionType.Target {
			var Target response.TargetComponent
			err = global.GVA_DB.Table("les_questionbank_target").Where("id = ?", studentPaper[i].QuestionId).Find(&Target.Target).Error
			if err != nil {
				return
			}
			examPaper.TargetComponent = append(examPaper.TargetComponent, Target)
			examPaper.TargetComponent[targetCount].MergeId = studentPaper[i].ID
			targetCount++
		}
	}
	var PaperId int64
	err = global.GVA_DB.Table("exam_student_paper").Select("paper_id").Where("student_id = ? and plan_id =?", examComing.StudentId, examComing.PlanId).Scan(&PaperId).Error
	if err != nil {
		return
	}
	examPaper.PaperId = uint(PaperId)
	status, err = targetExamService.CreateStatus(examComing)
	fmt.Println(status)
	if err != nil {
		return
	}
	return
}

func (targetExamService *TargetExamPaperService) CommitTargetExamPapers(examPaperCommit request.CommitTargetExamPaper) (err error) {
	err = global.GVA_DB.Table("student_paper_status").Where("student_id = ? and plan_id =?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("is_commit", 1).Error
	if err != nil {
		return
	}
	return
}
func (targetExamService *TargetExamPaperService) CreateStatus(examComing request.ExamComing) (status examManage.StudentPaperStatus, err error) {
	var num int64
	err = global.GVA_DB.Table("student_paper_status").Where("student_id = ? and plan_id = ?", examComing.StudentId, examComing.PlanId).Find(&status).Count(&num).Error
	if err != nil {
		return
	} else if num == 0 {
		status = examManage.StudentPaperStatus{
			GVA_MODEL: global.GVA_MODEL{},
			StudentId: examComing.StudentId,
			PlanId:    examComing.PlanId,
			EnterTime: time.Now(),
			IsCommit:  false,
		}
		global.GVA_DB.Create(&status)
	}
	return
}
