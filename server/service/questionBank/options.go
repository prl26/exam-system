package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	questionBankReq "github.com/flipped-aurora/gin-vue-admin/server/model/questionBank/request"
)

type OptionsService struct {
}

// CreateQuestionBankOptions 创建QuestionBankOptions记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_optionsService *OptionsService) CreateQuestionBankOptions(questionBank_options questionBank.Options) (err error) {
	err = global.GVA_DB.Create(&questionBank_options).Error
	return err
}

// DeleteQuestionBankOptions 删除QuestionBankOptions记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_optionsService *OptionsService) DeleteQuestionBankOptions(questionBank_options questionBank.Options) (err error) {
	err = global.GVA_DB.Delete(&questionBank_options).Error
	return err
}

// DeleteQuestionBankOptionsByIds 批量删除QuestionBankOptions记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_optionsService *OptionsService) DeleteQuestionBankOptionsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]questionBank.Options{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuestionBankOptions 更新QuestionBankOptions记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_optionsService *OptionsService) UpdateQuestionBankOptions(questionBank_options questionBank.Options) (err error) {
	err = global.GVA_DB.Save(&questionBank_options).Error
	return err
}

// GetQuestionBankOptions 根据id获取QuestionBankOptions记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_optionsService *OptionsService) GetQuestionBankOptions(id uint) (questionBank_options questionBank.Options, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBank_options).Error
	return
}

// GetQuestionBankOptionsInfoList 分页获取QuestionBankOptions记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_optionsService *OptionsService) GetQuestionBankOptionsInfoList(info questionBankReq.QuestionBankOptionsSearch) (list []questionBank.Options, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Options{})
	var questionBank_optionss []questionBank.Options
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Describe != "" {
		db = db.Where("describe LIKE ?", "%"+info.Describe+"%")
	}
	if info.Multiple_choice_id != nil {
		db = db.Where("multiple_choice_id = ?", info.Multiple_choice_id)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&questionBank_optionss).Error
	return questionBank_optionss, total, err
}
