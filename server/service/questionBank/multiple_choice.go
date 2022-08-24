package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	questionBankReq "github.com/flipped-aurora/gin-vue-admin/server/model/questionBank/request"
)

type MultipleChoiceService struct {
}

// CreateQuestionBankMultipleChoice 创建QuestionBankMultipleChoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_multiple_choiceService *MultipleChoiceService) CreateQuestionBankMultipleChoice(questionBank_multiple_choice questionBank.MultipleChoice) (err error) {
	err = global.GVA_DB.Create(&questionBank_multiple_choice).Error
	return err
}

// DeleteQuestionBankMultipleChoice 删除QuestionBankMultipleChoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_multiple_choiceService *MultipleChoiceService) DeleteQuestionBankMultipleChoice(questionBank_multiple_choice questionBank.MultipleChoice) (err error) {
	err = global.GVA_DB.Delete(&questionBank_multiple_choice).Error
	return err
}

// DeleteQuestionBankMultipleChoiceByIds 批量删除QuestionBankMultipleChoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_multiple_choiceService *MultipleChoiceService) DeleteQuestionBankMultipleChoiceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]questionBank.MultipleChoice{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuestionBankMultipleChoice 更新QuestionBankMultipleChoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_multiple_choiceService *MultipleChoiceService) UpdateQuestionBankMultipleChoice(questionBank_multiple_choice questionBank.MultipleChoice) (err error) {
	err = global.GVA_DB.Save(&questionBank_multiple_choice).Error
	return err
}

// GetQuestionBankMultipleChoice 根据id获取QuestionBankMultipleChoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_multiple_choiceService *MultipleChoiceService) GetQuestionBankMultipleChoice(id uint) (questionBank_multiple_choice questionBank.MultipleChoice, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBank_multiple_choice).Error
	return
}

// GetQuestionBankMultipleChoiceInfoList 分页获取QuestionBankMultipleChoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_multiple_choiceService *MultipleChoiceService) GetQuestionBankMultipleChoiceInfoList(info questionBankReq.QuestionBankMultipleChoiceSearch) (list []questionBank.MultipleChoice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.MultipleChoice{})
	var questionBank_multiple_choices []questionBank.MultipleChoice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Describe != "" {
		db = db.Where("describe LIKE ?", "%"+info.Describe+"%")
	}
	if info.Answer != "" {
		db = db.Where("answer = ?", info.Answer)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&questionBank_multiple_choices).Error
	return questionBank_multiple_choices, total, err
}
