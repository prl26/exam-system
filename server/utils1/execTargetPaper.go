package utils1

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/examType"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service"
	"gorm.io/gorm"
)

var targetService = service.ServiceGroupApp.QuestionBankServiceGroup.TargetService
var targetOjService = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.TargetService

func ExecTarget(examPaperCommit request.CommitTargetExamPaper) (err error) {
	//global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	Target := examPaperCommit.TargetComponent
	for _, v := range Target {
		address, _ := targetService.QueryExamRecord(examPaperCommit.StudentId, v.QuestionId, examPaperCommit.PlanId)
		//if !isGenerateAddress {
		//	return fmt.Errorf("暂未生成实例地址", err.Error())
		//}
		score, err := targetOjService.QueryScore(address)
		if err != nil {
			return fmt.Errorf("获取分数错误，请联系管理员或重新生成实例")
		}
		var result examManage.ExamStudentPaper
		err = global.GVA_DB.Raw("UPDATE exam_student_paper SET answer = ?,exam_student_paper.got_score = exam_student_paper.score*"+fmt.Sprintf("%f", float64(score)/100.0)+" where id = ?", address, v.MergeId).Scan(&result).Error
		if err != nil {
			return err
		}
	}
	//总分
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var sum float64
		tx.Raw("SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ?", examPaperCommit.StudentId, examPaperCommit.PlanId).Scan(&sum)
		global.GVA_LOG.Info(fmt.Sprintf("进入统分 %v", sum))
		var PlanDetail teachplan.ExamPlan
		tx.Model(teachplan.ExamPlan{}).Where("id =?", examPaperCommit.PlanId).Find(&PlanDetail)
		planId := int(PlanDetail.ID)
		if PlanDetail.Type == examType.FinalExam {
			tx.Select("exam_score", "final_exam_name", "final_exam_id").Where("student_id = ? and teach_class_id = ?", examPaperCommit.StudentId, PlanDetail.TeachClassId).Updates(teachplan.Score{
				ExamScrore:    &sum,
				FinalExamName: PlanDetail.Name,
				FinalExamId:   &planId,
			})
		} else if PlanDetail.Type == examType.ProceduralExam {
			fmt.Println("过程化统分统分")
			tx.Raw("UPDATE tea_score SET procedure_score = procedure_score+procedure_proportion/100*?)", sum).Where("student_id = ? and teach_class_id = ?", examPaperCommit.StudentId, PlanDetail.TeachClassId)
		}
		err = tx.Model(examManage.ExamScore{}).Where("student_id = ? and plan_id = ?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("score", sum).Error
		//var term basicdata.Term
		//var lesson basicdata.Lesson
		//tx.Model(&basicdata.Term{}).Where("id = ?", PlanDetail.TermId).Find(&term)
		//tx.Model(&basicdata.Lesson{}).Where("id = ?", PlanDetail.LessonId).Find(&lesson)
		//tx.Create(&examManage.ExamScore{
		//	StudentId:  &examPaperCommit.StudentId,
		//	PlanId:     &PlanDetail.ID,
		//	Name:       PlanDetail.Name,
		//	TermId:     PlanDetail.TermId,
		//	TermName:   term.Name,
		//	LessonId:   PlanDetail.LessonId,
		//	CourseName: lesson.Name,
		//	Score:      &sum,
		//	ExamType:   &PlanDetail.Type,
		//	StartTime:  PlanDetail.StartTime,
		//	Weight:     PlanDetail.Weight,
		//})
		return err
	})
	return
}
