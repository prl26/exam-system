package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    lessondataReq "github.com/flipped-aurora/gin-vue-admin/server/model/lessondata/request"
)

type QuestionBankKnowledgeMergeService struct {
}

// CreateQuestionBankKnowledgeMerge 创建QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *QuestionBankKnowledgeMergeService) CreateQuestionBankKnowledgeMerge(questionBankKnowledgeMerge lessondata.QuestionBankKnowledgeMerge) (err error) {
	err = global.GVA_DB.Create(&questionBankKnowledgeMerge).Error
	return err
}

// DeleteQuestionBankKnowledgeMerge 删除QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *QuestionBankKnowledgeMergeService)DeleteQuestionBankKnowledgeMerge(questionBankKnowledgeMerge lessondata.QuestionBankKnowledgeMerge) (err error) {
	err = global.GVA_DB.Delete(&questionBankKnowledgeMerge).Error
	return err
}

// DeleteQuestionBankKnowledgeMergeByIds 批量删除QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *QuestionBankKnowledgeMergeService)DeleteQuestionBankKnowledgeMergeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.QuestionBankKnowledgeMerge{},"id in ?",ids.Ids).Error
	return err
}

// UpdateQuestionBankKnowledgeMerge 更新QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *QuestionBankKnowledgeMergeService)UpdateQuestionBankKnowledgeMerge(questionBankKnowledgeMerge lessondata.QuestionBankKnowledgeMerge) (err error) {
	err = global.GVA_DB.Save(&questionBankKnowledgeMerge).Error
	return err
}

// GetQuestionBankKnowledgeMerge 根据id获取QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *QuestionBankKnowledgeMergeService)GetQuestionBankKnowledgeMerge(id uint) (questionBankKnowledgeMerge lessondata.QuestionBankKnowledgeMerge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBankKnowledgeMerge).Error
	return
}

// GetQuestionBankKnowledgeMergeInfoList 分页获取QuestionBankKnowledgeMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankKnowledgeMergeService *QuestionBankKnowledgeMergeService)GetQuestionBankKnowledgeMergeInfoList(info lessondataReq.QuestionBankKnowledgeMergeSearch) (list []lessondata.QuestionBankKnowledgeMerge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&lessondata.QuestionBankKnowledgeMerge{})
    var questionBankKnowledgeMerges []lessondata.QuestionBankKnowledgeMerge
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Knowledge_id != nil {
        db = db.Where("knowledge_id = ?",info.Knowledge_id)
    }
    if info.Question_id != nil {
        db = db.Where("question_id = ?",info.Question_id)
    }
    if info.Question_type != nil {
        db = db.Where("question_type = ?",info.Question_type)
    }
    if info.Difficulty != nil {
        db = db.Where("difficulty = ?",info.Difficulty)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&questionBankKnowledgeMerges).Error
	return  questionBankKnowledgeMerges, total, err
}
