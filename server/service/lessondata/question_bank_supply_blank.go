package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    lessondataReq "github.com/flipped-aurora/gin-vue-admin/server/model/lessondata/request"
)

type QuestionBankSupplyBlankService struct {
}

// CreateQuestionBankSupplyBlank 创建QuestionBankSupplyBlank记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_supply_blankService *QuestionBankSupplyBlankService) CreateQuestionBankSupplyBlank(questionBank_supply_blank lessondata.QuestionBankSupplyBlank) (err error) {
	err = global.GVA_DB.Create(&questionBank_supply_blank).Error
	return err
}

// DeleteQuestionBankSupplyBlank 删除QuestionBankSupplyBlank记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_supply_blankService *QuestionBankSupplyBlankService)DeleteQuestionBankSupplyBlank(questionBank_supply_blank lessondata.QuestionBankSupplyBlank) (err error) {
	err = global.GVA_DB.Delete(&questionBank_supply_blank).Error
	return err
}

// DeleteQuestionBankSupplyBlankByIds 批量删除QuestionBankSupplyBlank记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_supply_blankService *QuestionBankSupplyBlankService)DeleteQuestionBankSupplyBlankByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.QuestionBankSupplyBlank{},"id in ?",ids.Ids).Error
	return err
}

// UpdateQuestionBankSupplyBlank 更新QuestionBankSupplyBlank记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_supply_blankService *QuestionBankSupplyBlankService)UpdateQuestionBankSupplyBlank(questionBank_supply_blank lessondata.QuestionBankSupplyBlank) (err error) {
	err = global.GVA_DB.Save(&questionBank_supply_blank).Error
	return err
}

// GetQuestionBankSupplyBlank 根据id获取QuestionBankSupplyBlank记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_supply_blankService *QuestionBankSupplyBlankService)GetQuestionBankSupplyBlank(id uint) (questionBank_supply_blank lessondata.QuestionBankSupplyBlank, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBank_supply_blank).Error
	return
}

// GetQuestionBankSupplyBlankInfoList 分页获取QuestionBankSupplyBlank记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_supply_blankService *QuestionBankSupplyBlankService)GetQuestionBankSupplyBlankInfoList(info lessondataReq.QuestionBankSupplyBlankSearch) (list []lessondata.QuestionBankSupplyBlank, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&lessondata.QuestionBankSupplyBlank{})
    var questionBank_supply_blanks []lessondata.QuestionBankSupplyBlank
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Describe != "" {
        db = db.Where("describe = ?",info.Describe)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&questionBank_supply_blanks).Error
	return  questionBank_supply_blanks, total, err
}
