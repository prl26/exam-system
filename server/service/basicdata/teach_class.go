package basicdata

import (
	"exam-system/global"
	"exam-system/model/basicdata"
	basicdataReq "exam-system/model/basicdata/request"
	"exam-system/model/common/request"
)

type TeachClassService struct {
}

// CreateTeachClass 创建TeachClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassService *TeachClassService) CreateTeachClass(teachClass basicdata.TeachClass) (err error) {
	err = global.GVA_DB.Create(&teachClass).Error
	return err
}

// DeleteTeachClass 删除TeachClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassService *TeachClassService) DeleteTeachClass(teachClass basicdata.TeachClass) (err error) {
	err = global.GVA_DB.Delete(&teachClass).Error
	return err
}

// DeleteTeachClassByIds 批量删除TeachClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassService *TeachClassService) DeleteTeachClassByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.TeachClass{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeachClass 更新TeachClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassService *TeachClassService) UpdateTeachClass(teachClass basicdata.TeachClass) (err error) {
	err = global.GVA_DB.Save(&teachClass).Error
	return err
}

// GetTeachClass 根据id获取TeachClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassService *TeachClassService) GetTeachClass(id uint) (teachClass basicdata.TeachClass, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teachClass).Error
	return
}

// GetTeachClassInfoList 分页获取TeachClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassService *TeachClassService) GetTeachClassInfoList(info basicdataReq.TeachClassSearch) (list []basicdata.TeachClass, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&basicdata.TeachClass{})
	var teachClasss []basicdata.TeachClass
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CourseId != nil {
		db = db.Where("course_id = ?", info.CourseId)
	}
	if info.TermId != nil {
		db = db.Where("term_id = ?", info.TermId)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.TeacherId != nil {
		db = db.Where("teacher_id = ?", info.TeacherId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&teachClasss).Error
	return teachClasss, total, err
}
