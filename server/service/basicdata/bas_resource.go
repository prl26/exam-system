package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
)

type ResourceService struct {
}

// CreateResource 创建Resource记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService) CreateResource(resource basicdata.Resource) (err error) {
	err = global.GVA_DB.Create(&resource).Error
	return err
}

// DeleteResource 删除Resource记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)DeleteResource(resource basicdata.Resource) (err error) {
	err = global.GVA_DB.Delete(&resource).Error
	return err
}

// DeleteResourceByIds 批量删除Resource记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)DeleteResourceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.Resource{},"id in ?",ids.Ids).Error
	return err
}

// UpdateResource 更新Resource记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)UpdateResource(resource basicdata.Resource) (err error) {
	err = global.GVA_DB.Save(&resource).Error
	return err
}

// GetResource 根据id获取Resource记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)GetResource(id uint) (resource basicdata.Resource, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&resource).Error
	return
}

// GetResourceInfoList 分页获取Resource记录
// Author [piexlmax](https://github.com/piexlmax)
func (resourceService *ResourceService)GetResourceInfoList(info basicdataReq.ResourceSearch) (list []basicdata.Resource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&basicdata.Resource{})
    var resources []basicdata.Resource
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.DetailId != "" {
        db = db.Where("detail_id = ?",info.DetailId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&resources).Error
	return  resources, total, err
}
