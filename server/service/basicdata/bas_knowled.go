package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
)

type KnowledgeService struct {
}

// CreateKnowledge 创建Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService) CreateKnowledge(knowledge basicdata.Knowledge) (err error) {
	err = global.GVA_DB.Create(&knowledge).Error
	return err
}

// DeleteKnowledge 删除Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService)DeleteKnowledge(knowledge basicdata.Knowledge) (err error) {
	err = global.GVA_DB.Delete(&knowledge).Error
	return err
}

// DeleteKnowledgeByIds 批量删除Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService)DeleteKnowledgeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.Knowledge{},"id in ?",ids.Ids).Error
	return err
}

// UpdateKnowledge 更新Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService)UpdateKnowledge(knowledge basicdata.Knowledge) (err error) {
	err = global.GVA_DB.Save(&knowledge).Error
	return err
}

// GetKnowledge 根据id获取Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService)GetKnowledge(id uint) (knowledge basicdata.Knowledge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&knowledge).Error
	return
}

// GetKnowledgeInfoList 分页获取Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService)GetKnowledgeInfoList(info basicdataReq.KnowledgeSearch) (list []basicdata.Knowledge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&basicdata.Knowledge{})
    var knowledges []basicdata.Knowledge
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.ChapterId != "" {
        db = db.Where("chapter_id = ?",info.ChapterId)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&knowledges).Error
	return  knowledges, total, err
}
