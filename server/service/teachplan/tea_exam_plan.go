package teachplan

import (
	"github.com/prl26/exam-system/server/global"
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
	time := int64(endTime.Sub(startTime).Minutes())
	audit := 1
	state := 1
	ExamPlan := teachplan.ExamPlan{
		GVA_MODEL:     global.GVA_MODEL{},
		Name:          examPlan.Name,
		TeachClassId:  &examPlan.TeachClassId,
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
	}
	err = global.GVA_DB.Create(&ExamPlan).Error
	return err
}

// DeleteExamPlan 删除ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) DeleteExamPlan(examPlan teachplan.ExamPlan) (err error) {
	err = global.GVA_DB.Delete(&examPlan).Error
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
	Time := int64(endTime.Sub(startTime).Minutes())
	examPlanRq.UpdatedAt = time.Now()
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
	}
	err = global.GVA_DB.Omit("is_distributed", "user_id", "pre_plan_id", "created_at", "updated_at").Updates(&examPlan).Error

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
func (examPlanService *ExamPlanService) IsFinishPreExam(planId uint, studentId uint) (result bool, err error, preExamIds []string) {
	var examPlan teachplan.ExamPlan
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", planId).Find(&examPlan)
	preExamIds = strings.Split(examPlan.PrePlanId, ",")
	for _, v := range preExamIds {
		preExamId, _ := strconv.Atoi(v)
		var examRecords examManage.ExamScore
		var count int64
		err = global.GVA_DB.Where("plan_id = ? and student_id = ?", preExamId, studentId).Find(&examRecords).Count(&count).Error
		if err != nil {
			return false, err, preExamIds
		}
		if count == 0 {
			return false, nil, preExamIds
		}
	}
	return true, err, preExamIds
}
