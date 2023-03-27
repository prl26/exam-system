package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	"gorm.io/gorm"
	"time"
)

type ExamStudentPaperService struct {
}

// CreateExamStudentPaper 创建ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) CreateExamStudentPaper(examStudentPaper examManageReq.ExamComing) (studentPaper examManage.ExamStudentPaper, err error) {
	var examPlan teachplan.ExamPlan
	err = global.GVA_DB.Where("id = ?", examStudentPaper.PlanId).Find(&examPlan).Error
	if err != nil {
		return
	}
	var examPaper []examManage.ExamPaper
	err = global.GVA_DB.Where("template_id = ?", examPlan.TemplateId).Find(&examPaper).Error
	if err != nil {
		return
	}
	var paperItemsNum int64
	var AllMerge []response.AllPaperMerge
	for i := 0; i < len(examPaper); i++ {
		var paperMerge []examManage.PaperQuestionMerge
		err = global.GVA_DB.Table("exam_paper_question_merge").Where("paper_id = ?", examPaper[0].ID).Find(&paperMerge).Count(&paperItemsNum).Error
		if err != nil {
			return
		}
		AllMerge[i].PaperMerge = paperMerge
	}
	var studentIds []uint
	var numOfStudents int64
	err = global.GVA_DB.Table("bas_student_teach_classes").Select("student_id").Where("teach_class_id = ?", examPlan.TeachClassId).Find(&studentIds).Count(&numOfStudents).Error
	if err != nil {
		return
	}
	order := 0
	for i := 0; i < int(numOfStudents); i++ {
		for j := 0; j < int(paperItemsNum); j++ {
			studentPaper1 := examManage.ExamStudentPaper{
				GVA_MODEL:    global.GVA_MODEL{},
				PaperId:      &examPaper[order].ID,
				QuestionId:   AllMerge[order].PaperMerge[j].QuestionId,
				StudentId:    AllMerge[order].PaperMerge[j].QuestionId,
				Answer:       "",
				PlanId:       &examPlan.ID,
				Score:        nil,
				QuestionType: AllMerge[order].PaperMerge[j].QuestionType,
				ProblemType:  AllMerge[order].PaperMerge[j].ProblemType,
			}
			global.GVA_DB.Create(&studentPaper1)
			studentPaper = studentPaper1
		}
		order++
		if order == len(examPaper)-1 {
			order = 0
		}
	}
	return
}

// DeleteExamStudentPaper 删除ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) DeleteExamStudentPaper(examstudentPaper examManage.ExamStudentPaper) (err error) {
	err = global.GVA_DB.Delete(&examstudentPaper).Error
	return err
}

// DeleteExamStudentPaperByIds 批量删除ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) DeleteExamStudentPaperByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.ExamStudentPaper{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExamStudentPaper 更新ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) UpdateExamStudentPaper(examstudentPaper examManage.ExamStudentPaper) (err error) {
	err = global.GVA_DB.Updates(&examstudentPaper).Error
	return err
}

// GetExamStudentPaper 根据id获取ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) GetExamStudentPaper(id uint) (examstudentPaper examManage.ExamStudentPaper, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&examstudentPaper).Error
	return
}

//恢复学生考试资格
func (examstudentPaperService *ExamStudentPaperService) RecoverStudentPower(sid uint, pid uint) (err error) {
	//var pd teachplan.ExamPlan
	//err = global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", pid).Find(&pd).Error
	//if err != nil {
	//	return
	//}
	err = global.GVA_DB.Table("student_paper_status").Where("student_id = ? and plan_id =?", sid, pid).Updates(examManage.StudentPaperStatus{
		EnterTime: time.Now(),
		EndTime:   time.Now(),
	}).Error
	global.GVA_DB.Model(examManage.StudentPaperStatus{}).Where("student_id = ? and plan_id =?", sid, pid).Update("is_commit", 0)
	return
}

//上报分数
func (examstudentPaperService *ExamStudentPaperService) ReportScore(st teachplan.CoverRqs) (err error) {
	err = global.GVA_DB.Model(examManage.ExamScore{}).Where("plan_id =? and student_id in ?", st.PlanId, st.StudentIds).Update("is_report", 1).Error
	return
}

func (examstudentPaperService *ExamStudentPaperService) CheckIsCommit(st teachplan.CoverRq) (bool bool, err error) {
	err = global.GVA_DB.Model(examManage.StudentPaperStatus{}).Select("is_commit").Where("plan_id =? and student_id =?", st.PlanId, st.StudentId).Scan(&bool).Error
	return
}
func (examstudentPaperService *ExamStudentPaperService) ReportStudentScore(pid uint, sid uint) (err error) {
	err = global.GVA_DB.Model(examManage.ExamScore{}).Omit("updated_at").Where("plan_id =? and student_id =?", pid, sid).Update("is_report", 1).Error
	return
}

//查看成绩
func (examstudentPaperService *ExamStudentPaperService) ReviewScore(info examManageReq.PaperReview) (score []response.ExamScoreResponse1, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamScore{})
	var scores []examManage.ExamScore
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PlanId != nil {
		db = db.Where("plan_id = ?", info.PlanId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("score desc").Find(&scores).Error
	for _, v := range scores {
		var sName string
		var status examManage.StudentPaperStatus
		global.GVA_DB.Model(basicdata.Student{}).Select("name").Where("id = ?", v.StudentId).Find(&sName)
		global.GVA_DB.Model(examManage.StudentPaperStatus{}).Where("student_id = ? and plan_id = ?", v.StudentId, v.PlanId).Find(&status)
		temp := response.ExamScoreResponse1{
			StudentName: sName,
			ReviewScore: examManage.ReviewScore{
				UpdatedAt: v.UpdatedAt,
				Score:     v.Score,
				IsReport:  v.IsReport,
			},
			StudentPaperStatus: status,
		}
		score = append(score, temp)
	}
	return score, total, err
}
func (examstudentPaperService *ExamStudentPaperService) StudentPaperStatus(info examManageReq.StatusMonitor) (scores []response.PaperStatus, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.StudentPaperStatus{})
	var score []examManage.StudentPaperStatus
	// 如果有条件搜索 下方会自动创建搜索语句
	if &info.PlanId != nil {
		db = db.Where("plan_id = ?", info.PlanId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&score).Error
	for _, v := range score {
		var name string
		var isReport bool
		global.GVA_DB.Model(examManage.ExamScore{}).Select("is_report").Where("student_id = ? and plan_id =?", v.StudentId, info.PlanId).Find(&isReport)
		global.GVA_DB.Model(basicdata.Student{}).Select("name").Where("id = ?", v.StudentId).Find(&name)
		temp := response.PaperStatus{
			Name:               name,
			StudentPaperStatus: v,
			IsReport:           isReport,
		}
		scores = append(scores, temp)
	}
	return scores, total, err
}

// GetExamStudentPaperInfoList 分页获取ExamStudentPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examstudentPaperService *ExamStudentPaperService) GetExamStudentPaperInfoList(info examManageReq.ExamStudentPaperSearch) (list []examManage.ExamStudentPaper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamStudentPaper{})
	var examstudentPapers []examManage.ExamStudentPaper
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PaperId != nil {
		db = db.Where("paper_id = ?", info.PaperId)
	}
	if info.QuestionId != nil {
		db = db.Where("question_id = ?", info.QuestionId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&examstudentPapers).Error
	return examstudentPapers, total, err
}
func (examstudentPaperService *ExamStudentPaperService) PaperCheating(info examManageReq.PaperCheating) (err error) {
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range info.AnswerCheating {
			err = tx.Model(examManage.ExamStudentPaper{}).Where("id = ?", v.MergeId).Updates(examManage.ExamStudentPaper{Answer: v.Answer, GotScore: &v.GotScore}).Error
			if err != nil {
				return err
			}
		}
		var sum int64
		tx.Raw("SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ? and deleted_at is null", info.StudentId, info.PlanId).Scan(&sum)
		err = tx.Model(examManage.ExamScore{}).Where("student_id = ? and plan_id = ?", info.StudentId, info.PlanId).Update("score", sum).Error
		if err != nil {
			return err
		}
		return nil
	})
	return
}
func (examstudentPaperService *ExamStudentPaperService) GetCommitRecord(info examManageReq.RecordRq) (recordList []response.RecordRp, err error) {
	var Recordlist []examManage.ExamRecord
	err = global.GVA_DB.Model(examManage.ExamRecord{}).Where("student_id = ? and plan_id =?", info.Student, info.PlanId).Find(&Recordlist).Error
	if err != nil {
		return
	}
	for _, v := range Recordlist {
		var recordMerge []examManage.ExamRecordMerge
		err = global.GVA_DB.Model(examManage.ExamRecordMerge{}).Where("record_id = ?", v.ID).Find(&recordMerge).Error
		temp := response.RecordRp{
			ExamRecord:  v,
			RecorcMerge: recordMerge,
		}
		recordList = append(recordList, temp)
	}
	return
}

//删除学生考卷
func (examstudentPaperService *ExamStudentPaperService) DeleteAnswer(pid uint, sid uint) (err error) {
	err = global.GVA_DB.Model(examManage.ExamStudentPaper{}).Where("student_id = ? and plan_id = ?", sid, pid).Update("answer", "").Error
	if err != nil {
		return
	}
	var score examManage.ExamScore
	err = global.GVA_DB.Model(examManage.ExamScore{}).Where("student_id = ? and plan_id = ?", sid, pid).Delete(&score).Error
	if err != nil {
		return
	}
	var status examManage.StudentPaperStatus
	err = global.GVA_DB.Model(examManage.StudentPaperStatus{}).Where("student_id = ? and plan_id = ?", sid, pid).Delete(&status).Error
	if err != nil {
		return
	}
	return
}
func (examstudentPaperService *ExamStudentPaperService) RecoverByRecord(pid, sid, rid uint) (err error) {
	err = global.GVA_DB.Raw("UPDATE exam_student_paper as e,exam_record_merge as r SET e.answer = r.answer,e.got_score = r.got_score WHERE e.student_id = r.student_id and e.plan_id = r.plan_id and e.question_type = r.question_type and e.question_id =r.question_id and e.problem_type = r.problem_type and r.record_id = ?", rid).Error
	var sum float64
	global.GVA_DB.Raw("UPDATE exam_scores SET score = (SELECT SUM(got_score) FROM exam_student_paper as e where e.student_id = ? and e.plan_id = ?) WHERE student_id = ? and plan_id = ?", sid, pid, sid, pid).Scan(&sum)
	return
}
func (examstudentPaperService *ExamStudentPaperService) ForceCommitStudent(pid, sid uint) (err error) {
	err = global.GVA_DB.Model(examManage.StudentPaperStatus{}).Where("plan_id = ? and student_id =?", pid, sid).Update("is_commit", 1).Error
	return
}
