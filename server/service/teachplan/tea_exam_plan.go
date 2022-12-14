package teachplan

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/teachplan"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/utils"
)

type ExamPlanService struct {
}

// CreateExamPlan 创建ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) CreateExamPlan(examPlan teachplanReq.ExamPlanRq, userId uint) (err error) {
	startTime := utils.StringToTime(examPlan.StartTime)
	endTime := utils.StringToTime(examPlan.EndTime)
	time := int64(endTime.Sub(startTime).Minutes())
	state := 0
	ExamPlan := teachplan.ExamPlan{
		GVA_MODEL:     global.GVA_MODEL{},
		Name:          examPlan.Name,
		TeachClassId:  examPlan.TeachClassId,
		Time:          &time,
		StartTime:     &startTime,
		EndTime:       &endTime,
		CourseId:      examPlan.CourseId,
		TemplateId:    examPlan.TemplateId,
		State:         &state,
		Audit:         examPlan.Audit,
		Type:          *examPlan.Type,
		PassScore:     examPlan.PassScore,
		Weight:        examPlan.Weight,
		TermId:        examPlan.TermId,
		IsDistributed: false,
		UserId:        &userId,
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
func (examPlanService *ExamPlanService) UpdateExamPlan(examPlan teachplan.ExamPlan) (err error) {
	err = global.GVA_DB.Updates(&examPlan).Error
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
func (examPlanService *ExamPlanService) GetExamPlanInfoList(info teachplanReq.ExamPlanSearch, userId uint) (list []teachplan.ExamPlan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&teachplan.ExamPlan{})
	var examPlans []teachplan.ExamPlan
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("user_id = ?", userId)
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.CourseId != nil {
		db = db.Where("course_id = ?", info.CourseId)
	}
	if info.State != nil {
		db = db.Where("state = ?", info.State)
	}
	if info.Audit != nil {
		db = db.Where("audit = ?", info.Audit)
	}
	if &info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.TermId != nil {
		db = db.Where("term_id = ?", info.TermId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc,updated_at desc ").Limit(limit).Offset(offset).Find(&examPlans).Error
	return examPlans, total, err
}
