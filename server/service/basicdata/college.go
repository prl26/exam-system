package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CollegeService struct {
}

// CreateCollege 创建College记录
// Author [piexlmax](https://github.com/piexlmax)
func (collegeService *CollegeService) CreateCollege(college basicdata.College) (err error) {
	err = global.GVA_DB.Create(&college).Error
	return err
}

// DeleteCollege 删除College记录
// Author [piexlmax](https://github.com/piexlmax)
func (collegeService *CollegeService)DeleteCollege(college basicdata.College) (err error) {
	err = global.GVA_DB.Delete(&college).Error
	return err
}

// DeleteCollegeByIds 批量删除College记录
// Author [piexlmax](https://github.com/piexlmax)
func (collegeService *CollegeService)DeleteCollegeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.College{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCollege 更新College记录
// Author [piexlmax](https://github.com/piexlmax)
func (collegeService *CollegeService)UpdateCollege(college basicdata.College) (err error) {
	err = global.GVA_DB.Save(&college).Error
	return err
}

// GetCollege 根据id获取College记录
// Author [piexlmax](https://github.com/piexlmax)
func (collegeService *CollegeService)GetCollege(id uint) (college basicdata.College, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&college).Error
	return
}

// GetCollegeInfoList 分页获取College记录
// Author [piexlmax](https://github.com/piexlmax)
func (collegeService *CollegeService)GetCollegeInfoList(info basicdataReq.CollegeSearch) (list []basicdata.College, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&basicdata.College{})
    var colleges []basicdata.College
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&colleges).Error
	return  colleges, total, err
}
