package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	questionBankReq "github.com/flipped-aurora/gin-vue-admin/server/model/questionBank/request"
)

type KnowledgeMergeService struct {
}

// CreateQuestionBankKnowledgeMerge 创建QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *KnowledgeMergeService) CreateQuestionBankKnowledgeMerge(questionBankKnowledgeMerge questionBank.KnowledgeMerge) (err error) {
	err = global.GVA_DB.Create(&questionBankKnowledgeMerge).Error
	return err
}

// DeleteQuestionBankKnowledgeMerge 删除QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *KnowledgeMergeService) DeleteQuestionBankKnowledgeMerge(questionBankKnowledgeMerge questionBank.KnowledgeMerge) (err error) {
	err = global.GVA_DB.Delete(&questionBankKnowledgeMerge).Error
	return err
}

// DeleteQuestionBankKnowledgeMergeByIds 批量删除QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *KnowledgeMergeService) DeleteQuestionBankKnowledgeMergeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]questionBank.KnowledgeMerge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuestionBankKnowledgeMerge 更新QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *KnowledgeMergeService) UpdateQuestionBankKnowledgeMerge(questionBankKnowledgeMerge questionBank.KnowledgeMerge) (err error) {
	err = global.GVA_DB.Save(&questionBankKnowledgeMerge).Error
	return err
}

// GetQuestionBankKnowledgeMerge 根据id获取QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *KnowledgeMergeService) GetQuestionBankKnowledgeMerge(id uint) (questionBankKnowledgeMerge questionBank.KnowledgeMerge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBankKnowledgeMerge).Error
	return
}

// GetQuestionBankKnowledgeMergeInfoList 分页获取QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *KnowledgeMergeService) GetQuestionBankKnowledgeMergeInfoList(info questionBankReq.QuestionBankKnowledgeMergeSearch) (list []questionBank.KnowledgeMerge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.KnowledgeMerge{})
	var questionBankKnowledgeMerges []questionBank.KnowledgeMerge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.KnowledgeId != nil {
		db = db.Where("knowledge_id = ?", info.KnowledgeId)
	}
	if info.QuestionId != nil {
		db = db.Where("question_id = ?", info.QuestionId)
	}
	if info.QuestionType != nil {
		db = db.Where("question_type = ?", info.QuestionType)
	}
	if info.Difficulty != nil {
		db = db.Where("difficulty = ?", info.Difficulty)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&questionBankKnowledgeMerges).Error
	return questionBankKnowledgeMerges, total, err
}
