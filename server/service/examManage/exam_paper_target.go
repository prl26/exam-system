package examManage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/utils"
	"time"
)

type TargetExamPaperService struct {
}

func (targetExamService *TargetExamPaperService) GetTargetExamPapers(examComing request.ExamComing, ip string) (examPaper response.TargetExamPaperResponse, status examManage.StudentPaperStatus, examScore examManage.ExamScore, err error) {
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
	status, err = targetExamService.CreateStatus(examComing, ip)
	fmt.Println(status)
	var PlanDetail teachplan.ExamPlan
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id =?", examComing.PlanId).Find(&PlanDetail)
	examScore, err = utils.CreateExamScore(PlanDetail, 0, examComing.StudentId)
	_, err = targetExamService.CreateStatusRecord(examComing, ip)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	return
}
func (targetExamService *TargetExamPaperService) GetTargetExamPapersByRedis(examComing request.ExamComing, ip string) (examPaper response.TargetExamPaperResponse, status examManage.StudentPaperStatus, examScore examManage.ExamScore, err error) {
	//examPaper.TargetComponent = make([]response.TargetComponent, 0)
	//var studentPaper []examManage.ExamStudentPaper
	//err = global.GVA_DB.Where("student_id = ? and plan_id = ?", examComing.StudentId, examComing.PlanId).Find(&studentPaper).Error
	//var targetCount uint
	//for i := 0; i < len(studentPaper); i++ {
	//	if *studentPaper[i].QuestionType == questionType.Target {
	//		var Target response.TargetComponent
	//		err = global.GVA_DB.Table("les_questionbank_target").Where("id = ?", studentPaper[i].QuestionId).Find(&Target.Target).Error
	//		if err != nil {
	//			return
	//		}
	//		examPaper.TargetComponent = append(examPaper.TargetComponent, Target)
	//		examPaper.TargetComponent[targetCount].MergeId = studentPaper[i].ID
	//		targetCount++
	//	}
	//}
	examPaper1, err := global.GVA_REDIS.Get(context.Background(), fmt.Sprintf("studentPaper:%d:%d:%d:%d", 05, examComing.StudentId, examComing.PlanId, uint(questionType.Target))).Result()

	err = json.Unmarshal([]byte(examPaper1), &examPaper.TargetComponent)

	err = global.GVA_DB.Table("exam_student_paper").Select("paper_id").Where("student_id = ? and plan_id =?", examComing.StudentId, examComing.PlanId).Scan(&examPaper.PaperId).Error
	if err != nil {
		return
	}
	status, err = targetExamService.CreateStatus(examComing, ip)
	var PlanDetail teachplan.ExamPlan
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id =?", examComing.PlanId).Find(&PlanDetail)
	examScore, err = utils.CreateExamScore(PlanDetail, 0, examComing.StudentId)
	_, err = targetExamService.CreateStatusRecord(examComing, ip)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	return
}
func (targetExamService *TargetExamPaperService) CommitTargetExamPapers(examPaperCommit request.CommitTargetExamPaper) (err error) {
	err = global.GVA_DB.Model(examManage.StudentPaperStatus{}).Where("student_id = ? and plan_id =?", examPaperCommit.StudentId, examPaperCommit.PlanId).Updates(examManage.StudentPaperStatus{IsCommit: true, EndTime: time.Now()}).Error
	if err != nil {
		return
	}
	return
}
func (targetExamService *TargetExamPaperService) GetTargetExamScore(info request.ExamStudentScore, studentId uint) (studentScore []response.ExamScoreResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamScore{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if *info.TermId != 0 {
		db = db.Where("term_id = ?", info.TermId)
	}
	db = db.Where("lesson_id = 25")

	err = db.Where("student_id = ?", studentId).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Where("student_id = ?", studentId).Order("created_at desc,updated_at desc ").Limit(limit).Offset(offset).Find(&studentScore).Error
	return studentScore, total, err
}
func (targetExamService *TargetExamPaperService) CreateStatus(examComing request.ExamComing, IP string) (status examManage.StudentPaperStatus, err error) {
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
			EndTime:   time.Now(),
			IsCommit:  false,
			Ip:        IP,
		}
		global.GVA_DB.Create(&status)
	}
	return
}
func (targetExamService *TargetExamPaperService) CreateStatusRecord(examComing request.ExamComing, IP string) (status examManage.ExamRecord, err error) {
	status = examManage.ExamRecord{
		GVA_MODEL: global.GVA_MODEL{},
		StudentId: examComing.StudentId,
		PlanId:    examComing.PlanId,
		EnterTime: time.Now(),
		EndTime:   time.Now(),
		Ip:        IP,
	}
	err = global.GVA_DB.Create(&status).Error
	return
}
