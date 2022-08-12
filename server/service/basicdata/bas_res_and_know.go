package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
)

type ResandknowService struct {
}

// CreateResandknow 创建Resandknow记录
// Author [piexlmax](https://github.com/piexlmax)
func (resandknowService *ResandknowService) CreateResandknow(resandknow basicdata.Resandknow) (err error) {
	err = global.GVA_DB.Create(&resandknow).Error
	return err
}

// DeleteResandknow 删除Resandknow记录
// Author [piexlmax](https://github.com/piexlmax)
func (resandknowService *ResandknowService)DeleteResandknow(resandknow basicdata.Resandknow) (err error) {
	err = global.GVA_DB.Delete(&resandknow).Error
	return err
}

// DeleteResandknowByIds 批量删除Resandknow记录
// Author [piexlmax](https://github.com/piexlmax)
func (resandknowService *ResandknowService)DeleteResandknowByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.Resandknow{},"id in ?",ids.Ids).Error
	return err
}

// UpdateResandknow 更新Resandknow记录
// Author [piexlmax](https://github.com/piexlmax)
func (resandknowService *ResandknowService)UpdateResandknow(resandknow basicdata.Resandknow) (err error) {
	err = global.GVA_DB.Save(&resandknow).Error
	return err
}

// GetResandknow 根据id获取Resandknow记录
// Author [piexlmax](https://github.com/piexlmax)
func (resandknowService *ResandknowService)GetResandknow(id uint) (resandknow basicdata.Resandknow, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&resandknow).Error
	return
}

// GetResandknowInfoList 分页获取Resandknow记录
// Author [piexlmax](https://github.com/piexlmax)
func (resandknowService *ResandknowService)GetResandknowInfoList(info basicdataReq.ResandknowSearch) (list []basicdata.Resandknow, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&basicdata.Resandknow{})
    var resandknows []basicdata.Resandknow
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.ResourceId != nil {
        db = db.Where("resource_id = ?",info.ResourceId)
    }
    if info.KnowledgeId != nil {
        db = db.Where("knowledge_id = ?",info.KnowledgeId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&resandknows).Error
	return  resandknows, total, err
}
