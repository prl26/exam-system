package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	lessondataReq "github.com/flipped-aurora/gin-vue-admin/server/model/lessondata/request"
)

type QuestionBankProgrammService struct {
}

// CreateQuestionBankProgramm 创建QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *QuestionBankProgrammService) CreateQuestionBankProgramm(questionBankProgramm lessondata.QuestionBankProgramm) (err error) {
	err = global.GVA_DB.Create(&questionBankProgramm).Error
	return err
}

// DeleteQuestionBankProgramm 删除QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *QuestionBankProgrammService) DeleteQuestionBankProgramm(questionBankProgramm lessondata.QuestionBankProgramm) (err error) {
	err = global.GVA_DB.Delete(&questionBankProgramm).Error
	return err
}

// DeleteQuestionBankProgrammByIds 批量删除QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *QuestionBankProgrammService) DeleteQuestionBankProgrammByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.QuestionBankProgramm{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuestionBankProgramm 更新QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *QuestionBankProgrammService) UpdateQuestionBankProgramm(questionBankProgramm lessondata.QuestionBankProgramm) (err error) {
	err = global.GVA_DB.Save(&questionBankProgramm).Error
	return err
}

// GetQuestionBankProgramm 根据id获取QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *QuestionBankProgrammService) GetQuestionBankProgramm(id uint) (questionBankProgramm lessondata.QuestionBankProgramm, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBankProgramm).Error
	return
}

// GetQuestionBankProgrammInfoList 分页获取QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *QuestionBankProgrammService) GetQuestionBankProgrammInfoList(info lessondataReq.QuestionBankProgrammSearch) (list []lessondata.QuestionBankProgramm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lessondata.QuestionBankProgramm{})
	var questionBankProgramms []lessondata.QuestionBankProgramm
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&questionBankProgramms).Error
	return questionBankProgramms, total, err
}
