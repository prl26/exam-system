package teachplan

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/teachplan"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/model/teachplan/response"
	"github.com/prl26/exam-system/server/utils"
	"strconv"
	"strings"
	"time"
)

type ExamPlanService struct {
}

// CreateExamPlan 创建ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) CreateExamPlan(examPlan teachplanReq.ExamPlanRq, userId uint) (err error) {
	startTime := utils.StringToTime(examPlan.StartTime)
	endTime := utils.StringToTime(examPlan.EndTime)
	audit := 1
	state := 1
	time := int64(examPlan.Time)
	time1 := int64(endTime.Sub(startTime).Minutes())
	teachClassIds := examPlan.TeachClassId
	for _, v := range teachClassIds {
		var teachClassdetail basicdata.TeachClass
		global.GVA_DB.Model(basicdata.TeachClass{}).Where("id = ?", v).Scan(&teachClassdetail)
		name := examPlan.Name + "--" + teachClassdetail.Name
		if examPlan.IsLimitTime == false {
			ExamPlan := teachplan.ExamPlan{
				GVA_MODEL:     global.GVA_MODEL{},
				Name:          name,
				TeachClassId:  &v,
				Time:          &time1,
				StartTime:     &startTime,
				EndTime:       &endTime,
				LessonId:      &examPlan.LessonId,
				TemplateId:    &examPlan.TemplateId,
				State:         &state,
				Audit:         &audit,
				Type:          examPlan.Type,
				PassScore:     &examPlan.PassScore,
				TermId:        &examPlan.TermId,
				IsDistributed: false,
				UserId:        &userId,
				PrePlanId:     "0",
				Weight:        &examPlan.Weight,
				IsLimitTime:   examPlan.IsLimitTime,
				IsReady:       false,
			}
			err = global.GVA_DB.Create(&ExamPlan).Error
		} else {
			ExamPlan := teachplan.ExamPlan{
				GVA_MODEL:     global.GVA_MODEL{},
				Name:          name,
				TeachClassId:  &v,
				Time:          &time,
				StartTime:     &startTime,
				EndTime:       &endTime,
				LessonId:      &examPlan.LessonId,
				TemplateId:    &examPlan.TemplateId,
				State:         &state,
				Audit:         &audit,
				Type:          examPlan.Type,
				PassScore:     &examPlan.PassScore,
				TermId:        &examPlan.TermId,
				IsDistributed: false,
				UserId:        &userId,
				PrePlanId:     "0",
				Weight:        &examPlan.Weight,
				IsLimitTime:   examPlan.IsLimitTime,
				IsReady:       false,
			}
			err = global.GVA_DB.Create(&ExamPlan).Error
		}
	}
	return err
}

// DeleteExamPlan 删除ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) DeleteExamPlan(examPlan teachplan.ExamPlan) (err error) {
	err = global.GVA_DB.Delete(&examPlan).Error
	var examPaper []examManage.ExamPaper
	err = global.GVA_DB.Model(&examManage.ExamPaper{}).Delete(&examPaper).Where("plan_id = ?", examPlan.ID).Error
	var examStudentPaper []examManage.ExamStudentPaper
	err = global.GVA_DB.Model(&examManage.ExamStudentPaper{}).Delete(&examStudentPaper).Where("plan_id = ?", examPlan.ID).Error
	var examScore []examManage.ExamScore
	err = global.GVA_DB.Model(&examManage.ExamScore{}).Delete(&examScore).Where("plan_id = ?", examPlan.ID).Error
	return err
}

// DeleteExamPlanByIds 批量删除ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) DeleteExamPlanByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]teachplan.ExamPlan{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExamPlan 更新ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) UpdateExamPlan(examPlanRq teachplanReq.ExamPlanRq1) (err error) {
	startTime := utils.StringToTime(examPlanRq.StartTime)
	endTime := utils.StringToTime(examPlanRq.EndTime)
	Time := int64(examPlanRq.Time)
	examPlanRq.UpdatedAt = time.Now()
	b := examPlanRq.IsLimitTime
	examPlan := teachplan.ExamPlan{
		GVA_MODEL:    examPlanRq.GVA_MODEL,
		Name:         examPlanRq.Name,
		TeachClassId: &examPlanRq.TeachClassId,
		Time:         &Time,
		StartTime:    &startTime,
		EndTime:      &endTime,
		LessonId:     &examPlanRq.LessonId,
		TemplateId:   &examPlanRq.TemplateId,
		State:        &examPlanRq.State,
		Audit:        &examPlanRq.Audit,
		Type:         examPlanRq.Type,
		PassScore:    &examPlanRq.PassScore,
		TermId:       &examPlanRq.TermId,
		Weight:       &examPlanRq.Weight,
	}
	err = global.GVA_DB.Omit("is_distributed", "user_id", "pre_plan_id", "created_at").Updates(&examPlan).Error
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", examPlanRq.ID).Update("is_limit_time", b)
	return err
}
func (examPlanService *ExamPlanService) UpdatePrePlan(info request.PrePlanReq) (err error) {
	result := strings.Join(info.Ids, ",")
	err = global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", info.PlanId).Update("pre_plan_id", result).Error
	return err
}

// GetExamPlan 根据id获取ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) GetExamPlan(id uint) (examPlan teachplan.ExamPlan, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&examPlan).Error
	return
}

// GetExamPlanInfoList 分页获取ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) GetExamPlanInfoList(info teachplanReq.ExamPlanSearch, userId uint, authorityID uint) (list []teachplan.ExamPlan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&teachplan.ExamPlan{})
	var examPlans []teachplan.ExamPlan
	// 如果有条件搜索 下方会自动创建搜索语句
	if authorityID != 888 {
		db = db.Where("user_id = ?", userId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.LessonId != 0 {
		db = db.Where("lesson_id = ?", info.LessonId)
	}
	if info.State != 0 {
		db = db.Where("state = ?", info.State)
	}
	if info.Audit != 0 {
		db = db.Where("audit = ?", info.Audit)
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.TermId != 0 {
		db = db.Where("term_id = ?", info.TermId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc,updated_at desc ").Limit(limit).Offset(offset).Find(&examPlans).Error
	return examPlans, total, err
}
func (examPlanService *ExamPlanService) GetExamPlanDetail(list []teachplan.ExamPlan) (detail []response.ExamPlanRp) {
	len := len(list)
	result := make([]response.ExamPlanRp, 0, len)
	for i := 0; i < len; i++ {
		lessonName := utils.GetLessonName(*list[i].LessonId)
		TermName := utils.GetTermName(*list[i].TermId)
		TeachClassName := utils.GetTeachPlanName(*list[i].TeachClassId)
		temp := response.ExamPlanRp{
			ExamPlan: list[i],
			ItemName: response.ItemName{
				LessonName:     lessonName,
				TermName:       TermName,
				TeachClassName: TeachClassName,
			},
		}
		result = append(result, temp)
	}
	return result
}
func (examPlanService *ExamPlanService) ChangeStatus(planId uint) (err error) {
	var examPlan teachplan.ExamPlan
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", planId).Find(&examPlan)
	if *examPlan.State == 1 {
		err = global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", planId).Update("state", 2).Error
	} else {
		err = global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", planId).Update("state", 1).Error
	}
	return
}
func (examPlanService *ExamPlanService) ChangeAudit(planId uint, value uint) (err error) {
	err = global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", planId).Update("audit", value).Error
	return
}
func (examPlanService *ExamPlanService) IsFinishPreExam(planId uint, studentId uint) (result bool, err error, preExamNames []string) {
	var examPlan teachplan.ExamPlan
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", planId).Find(&examPlan)
	preExamIds := strings.Split(examPlan.PrePlanId, ",")
	for _, v := range preExamIds {
		preExamId, _ := strconv.Atoi(v)
		if preExamId == 0 {
			continue
		}
		var examRecords examManage.ExamScore
		var count int64
		var prePlanDetail teachplan.ExamPlan
		err = global.GVA_DB.Where("id = ?", preExamId).Find(&prePlanDetail).Error
		err = global.GVA_DB.Where("plan_id = ? and student_id = ?", preExamId, studentId).Find(&examRecords).Count(&count).Error
		preExamNames = append(preExamNames, prePlanDetail.Name)
		if err != nil {
			return false, err, preExamIds
		}
		if count == 0 {
			return false, nil, preExamIds
		} else if *examRecords.Score < *prePlanDetail.PassScore {
			return false, nil, preExamIds
		}
	}
	return true, err, preExamIds
}
func (examPlanService *ExamPlanService) CheckIsExamSt(planId uint, studentId uint) (result bool, err error) {
	var count int64
	err = global.GVA_DB.Model(examManage.ExamStudentPaper{}).Where("student_id = ? and plan_id =?", studentId, planId).Count(&count).Error
	if count > 0 {
		return true, err
	} else {
		return false, err
	}
}
