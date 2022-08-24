package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	examManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/examManage/request"
)

type ExamPaperTemplateService struct {
}

// CreateExamPaperTemplate 创建ExamPaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperTemplateService *ExamPaperTemplateService) CreateExamPaperTemplate(examPaperTemplate examManage.ExamPaperTemplate) (err error) {
	err = global.GVA_DB.Create(&examPaperTemplate).Error
	return err
}

// DeleteExamPaperTemplate 删除ExamPaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperTemplateService *ExamPaperTemplateService) DeleteExamPaperTemplate(examPaperTemplate examManage.ExamPaperTemplate) (err error) {
	err = global.GVA_DB.Delete(&examPaperTemplate).Error
	return err
}

// DeleteExamPaperTemplateByIds 批量删除ExamPaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperTemplateService *ExamPaperTemplateService) DeleteExamPaperTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.ExamPaperTemplate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExamPaperTemplate 更新ExamPaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperTemplateService *ExamPaperTemplateService) UpdateExamPaperTemplate(examPaperTemplate examManage.ExamPaperTemplate) (err error) {
	err = global.GVA_DB.Save(&examPaperTemplate).Error
	return err
}

// GetExamPaperTemplate 根据id获取ExamPaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperTemplateService *ExamPaperTemplateService) GetExamPaperTemplate(id uint) (examPaperTemplate examManage.ExamPaperTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&examPaperTemplate).Error
	return
}

// GetExamPaperTemplateInfoList 分页获取ExamPaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperTemplateService *ExamPaperTemplateService) GetExamPaperTemplateInfoList(info examManageReq.ExamPaperTemplateSearch) (list []examManage.ExamPaperTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamPaperTemplate{})
	var examPaperTemplates []examManage.ExamPaperTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CourseId != nil {
		db = db.Where("course_id = ?", info.CourseId)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&examPaperTemplates).Error
	return examPaperTemplates, total, err
}
