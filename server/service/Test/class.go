package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Test"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    TestReq "github.com/flipped-aurora/gin-vue-admin/server/model/Test/request"
)

type ClassService struct {
}

// CreateClass 创建Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService) CreateClass(class Test.Class) (err error) {
	err = global.GVA_DB.Create(&class).Error
	return err
}

// DeleteClass 删除Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService)DeleteClass(class Test.Class) (err error) {
	err = global.GVA_DB.Delete(&class).Error
	return err
}

// DeleteClassByIds 批量删除Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService)DeleteClassByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Test.Class{},"id in ?",ids.Ids).Error
	return err
}

// UpdateClass 更新Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService)UpdateClass(class Test.Class) (err error) {
	err = global.GVA_DB.Save(&class).Error
	return err
}

// GetClass 根据id获取Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService)GetClass(id uint) (class Test.Class, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&class).Error
	return
}

// GetClassInfoList 分页获取Class记录
// Author [piexlmax](https://github.com/piexlmax)
func (classService *ClassService)GetClassInfoList(info TestReq.ClassSearch) (list []Test.Class, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Test.Class{})
    var classs []Test.Class
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Id != nil {
        db = db.Where("id = ?",info.Id)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&classs).Error
	return  classs, total, err
}
