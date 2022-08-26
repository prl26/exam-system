package lessondata

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/lessondata"
	lessondataReq "exam-system/model/lessondata/request"
)

type KnowledgeService struct {
}

// CreateKnowledge 创建Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService) CreateKnowledge(knowledge lessondata.Knowledge) (err error) {
	err = global.GVA_DB.Create(&knowledge).Error
	return err
}

// DeleteKnowledge 删除Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService) DeleteKnowledge(knowledge lessondata.Knowledge) (err error) {
	err = global.GVA_DB.Delete(&knowledge).Error
	return err
}

// DeleteKnowledgeByIds 批量删除Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService) DeleteKnowledgeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.Knowledge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateKnowledge 更新Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService) UpdateKnowledge(knowledge lessondata.Knowledge) (err error) {
	err = global.GVA_DB.Save(&knowledge).Error
	return err
}

// GetKnowledge 根据id获取Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService) GetKnowledge(id uint) (knowledge lessondata.Knowledge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&knowledge).Error
	return
}

// GetKnowledgeInfoList 分页获取Knowledge记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgeService *KnowledgeService) GetKnowledgeInfoList(info lessondataReq.KnowledgeSearch) (list []lessondata.Knowledge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lessondata.Knowledge{})
	var knowledges []lessondata.Knowledge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CourseId != nil {
		db = db.Where("course_id = ?", info.CourseId)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&knowledges).Error
	return knowledges, total, err
}
