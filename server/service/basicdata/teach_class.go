package basicdata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
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
	// 删除教学班级之前先删除 该教学计划班级 与学生的引用关系
	global.GVA_DB.Model(&teachClass).Association("Student").Clear()
	err = global.GVA_DB.Delete(&teachClass).Error
	return err
}

// DeleteTeachClassByIds 批量删除TeachClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassService *TeachClassService) DeleteTeachClassByIds(ids request.IdsReq) (err error) {
	var teachClass basicdata.TeachClass
	Ids := ids.Ids
	for i := 0; i < len(Ids); i++ {
		teachClass.ID = uint(Ids[i])
		global.GVA_DB.Model(&teachClass).Association("Student").Clear()
	}
	err = global.GVA_DB.Delete(&[]basicdata.TeachClass{}, "id in ?", Ids).Error
	return err
}

// UpdateTeachClass 更新TeachClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassService *TeachClassService) UpdateTeachClass(teachClass basicdata.TeachClass) (err error) {
	err = global.GVA_DB.Updates(&teachClass).Error
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
	err = db.Limit(limit).Offset(offset).Find(&teachClasss).Error
	return teachClasss, total, err
}

func (teachClassService *TeachClassService) GetTeachAllClassInfoList(info basicdataReq.TeachClassSearch) (list []basicdata.TeachClass, total int64, err error) {
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

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&teachClasss).Error
	return teachClasss, total, err
}
