package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
)

type LessonService struct {
}

// CreateLesson 创建Lesson记录
// Author [piexlmax](https://github.com/piexlmax)
func (lessonService *LessonService) CreateLesson(lesson basicdata.Lesson) (err error) {
	err = global.GVA_DB.Create(&lesson).Error
	return err
}

// DeleteLesson 删除Lesson记录
// Author [piexlmax](https://github.com/piexlmax)
func (lessonService *LessonService)DeleteLesson(lesson basicdata.Lesson) (err error) {
	err = global.GVA_DB.Delete(&lesson).Error
	return err
}

// DeleteLessonByIds 批量删除Lesson记录
// Author [piexlmax](https://github.com/piexlmax)
func (lessonService *LessonService)DeleteLessonByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.Lesson{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLesson 更新Lesson记录
// Author [piexlmax](https://github.com/piexlmax)
func (lessonService *LessonService)UpdateLesson(lesson basicdata.Lesson) (err error) {
	err = global.GVA_DB.Save(&lesson).Error
	return err
}

// GetLesson 根据id获取Lesson记录
// Author [piexlmax](https://github.com/piexlmax)
func (lessonService *LessonService)GetLesson(id uint) (lesson basicdata.Lesson, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lesson).Error
	return
}

// GetLessonInfoList 分页获取Lesson记录
// Author [piexlmax](https://github.com/piexlmax)
func (lessonService *LessonService)GetLessonInfoList(info basicdataReq.LessonSearch) (list []basicdata.Lesson, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&basicdata.Lesson{})
    var lessons []basicdata.Lesson
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&lessons).Error
	return  lessons, total, err
}
