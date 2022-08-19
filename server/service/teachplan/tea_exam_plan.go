package teachplan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/teachplan"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    teachplanReq "github.com/flipped-aurora/gin-vue-admin/server/model/teachplan/request"
)

type ExamPlanService struct {
}

// CreateExamPlan 创建ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService) CreateExamPlan(examPlan teachplan.ExamPlan) (err error) {
	err = global.GVA_DB.Create(&examPlan).Error
	return err
}

// DeleteExamPlan 删除ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService)DeleteExamPlan(examPlan teachplan.ExamPlan) (err error) {
	err = global.GVA_DB.Delete(&examPlan).Error
	return err
}

// DeleteExamPlanByIds 批量删除ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService)DeleteExamPlanByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]teachplan.ExamPlan{},"id in ?",ids.Ids).Error
	return err
}

// UpdateExamPlan 更新ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService)UpdateExamPlan(examPlan teachplan.ExamPlan) (err error) {
	err = global.GVA_DB.Save(&examPlan).Error
	return err
}

// GetExamPlan 根据id获取ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService)GetExamPlan(id uint) (examPlan teachplan.ExamPlan, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&examPlan).Error
	return
}

// GetExamPlanInfoList 分页获取ExamPlan记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPlanService *ExamPlanService)GetExamPlanInfoList(info teachplanReq.ExamPlanSearch) (list []teachplan.ExamPlan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&teachplan.ExamPlan{})
    var examPlans []teachplan.ExamPlan
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.TeachClassId != nil {
        db = db.Where("teach_class_id = ?",info.TeachClassId)
    }
    if info.CourseId != nil {
        db = db.Where("course_id = ?",info.CourseId)
    }
    if info.State != nil {
        db = db.Where("state = ?",info.State)
    }
    if info.Audit != nil {
        db = db.Where("audit = ?",info.Audit)
    }
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&examPlans).Error
	return  examPlans, total, err
}
