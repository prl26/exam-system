package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	questionBankReq "github.com/flipped-aurora/gin-vue-admin/server/model/questionBank/request"
)

type ProgrammCaseService struct {
}

// CreateQuestionBankProgrammCase 创建QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *ProgrammCaseService) CreateQuestionBankProgrammCase(questionBankProgrammCase questionBank.ProgrammCase) (err error) {
	err = global.GVA_DB.Create(&questionBankProgrammCase).Error
	return err
}

// DeleteQuestionBankProgrammCase 删除QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *ProgrammCaseService) DeleteQuestionBankProgrammCase(questionBankProgrammCase questionBank.ProgrammCase) (err error) {
	err = global.GVA_DB.Delete(&questionBankProgrammCase).Error
	return err
}

// DeleteQuestionBankProgrammCaseByIds 批量删除QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *ProgrammCaseService) DeleteQuestionBankProgrammCaseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]questionBank.ProgrammCase{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuestionBankProgrammCase 更新QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *ProgrammCaseService) UpdateQuestionBankProgrammCase(questionBankProgrammCase questionBank.ProgrammCase) (err error) {
	err = global.GVA_DB.Save(&questionBankProgrammCase).Error
	return err
}

// GetQuestionBankProgrammCase 根据id获取QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *ProgrammCaseService) GetQuestionBankProgrammCase(id uint) (questionBankProgrammCase questionBank.ProgrammCase, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBankProgrammCase).Error
	return
}

// GetQuestionBankProgrammCaseInfoList 分页获取QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *ProgrammCaseService) GetQuestionBankProgrammCaseInfoList(info questionBankReq.QuestionBankProgrammCaseSearch) (list []questionBank.ProgrammCase, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.ProgrammCase{})
	var questionBankProgrammCases []questionBank.ProgrammCase
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ProgrammId != "" {
		db = db.Where("programm_id = ?", info.ProgrammId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&questionBankProgrammCases).Error
	return questionBankProgrammCases, total, err
}
