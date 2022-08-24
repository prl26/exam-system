package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	lessondataReq "github.com/flipped-aurora/gin-vue-admin/server/model/lessondata/request"
)

type QuestionBankProgrammCaseService struct {
}

// CreateQuestionBankProgrammCase 创建QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *QuestionBankProgrammCaseService) CreateQuestionBankProgrammCase(questionBankProgrammCase lessondata.QuestionBankProgrammCase) (err error) {
	err = global.GVA_DB.Create(&questionBankProgrammCase).Error
	return err
}

// DeleteQuestionBankProgrammCase 删除QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *QuestionBankProgrammCaseService) DeleteQuestionBankProgrammCase(questionBankProgrammCase lessondata.QuestionBankProgrammCase) (err error) {
	err = global.GVA_DB.Delete(&questionBankProgrammCase).Error
	return err
}

// DeleteQuestionBankProgrammCaseByIds 批量删除QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *QuestionBankProgrammCaseService) DeleteQuestionBankProgrammCaseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.QuestionBankProgrammCase{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuestionBankProgrammCase 更新QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *QuestionBankProgrammCaseService) UpdateQuestionBankProgrammCase(questionBankProgrammCase lessondata.QuestionBankProgrammCase) (err error) {
	err = global.GVA_DB.Save(&questionBankProgrammCase).Error
	return err
}

// GetQuestionBankProgrammCase 根据id获取QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *QuestionBankProgrammCaseService) GetQuestionBankProgrammCase(id uint) (questionBankProgrammCase lessondata.QuestionBankProgrammCase, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBankProgrammCase).Error
	return
}

// GetQuestionBankProgrammCaseInfoList 分页获取QuestionBankProgrammCase记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammCaseService *QuestionBankProgrammCaseService) GetQuestionBankProgrammCaseInfoList(info lessondataReq.QuestionBankProgrammCaseSearch) (list []lessondata.QuestionBankProgrammCase, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lessondata.QuestionBankProgrammCase{})
	var questionBankProgrammCases []lessondata.QuestionBankProgrammCase
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
