package lessondata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/lessondata"
	lessondataReq "github.com/prl26/exam-system/server/model/lessondata/request"
)

type ResourcePracticeService struct {
}

// CreateResourcePractice 创建ResourcePractice记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesPracticeService *ResourcePracticeService) CreateResourcePractice(resourcesPractice lessondata.ResourcePractice) (err error) {
	err = global.GVA_DB.Create(&resourcesPractice).Error
	return err
}

// DeleteResourcePractice 删除ResourcePractice记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesPracticeService *ResourcePracticeService) DeleteResourcePractice(resourcesPractice lessondata.ResourcePractice) (err error) {
	err = global.GVA_DB.Delete(&resourcesPractice).Error
	return err
}

// DeleteResourcePracticeByIds 批量删除ResourcePractice记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesPracticeService *ResourcePracticeService) DeleteResourcePracticeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.ResourcePractice{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateResourcePractice 更新ResourcePractice记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesPracticeService *ResourcePracticeService) UpdateResourcePractice(resourcesPractice lessondata.ResourcePractice) (err error) {
	err = global.GVA_DB.Updates(&resourcesPractice).Error
	return err
}

// GetResourcePractice 根据id获取ResourcePractice记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesPracticeService *ResourcePracticeService) GetResourcePractice(id uint) (resourcesPractice lessondata.ResourcePractice, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&resourcesPractice).Error
	return
}

// GetResourcePracticeInfoList 分页获取ResourcePractice记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourcesPracticeService *ResourcePracticeService) GetResourcePracticeInfoList(info lessondataReq.ResourcePracticeSearch) (list []lessondata.ResourcePractice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lessondata.ResourcePractice{})
	var resourcesPractices []lessondata.ResourcePractice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.QuestionId != nil {
		db = db.Where("qution_id = ?", info.QuestionId)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Orders != nil {
		db = db.Where("orders = ?", info.Orders)
	}
	if info.ResourcesId != nil {
		db = db.Where("resources_id = ?", info.ResourcesId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&resourcesPractices).Error
	return resourcesPractices, total, err
}
