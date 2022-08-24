package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	questionBankReq "github.com/flipped-aurora/gin-vue-admin/server/model/questionBank/request"
)

type JudgeService struct {
}

// CreateQuestionBankJudge 创建QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_judgeService *JudgeService) CreateQuestionBankJudge(questionBank_judge questionBank.Judge) (err error) {
	err = global.GVA_DB.Create(&questionBank_judge).Error
	return err
}

// DeleteQuestionBankJudge 删除QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_judgeService *JudgeService) DeleteQuestionBankJudge(questionBank_judge questionBank.Judge) (err error) {
	err = global.GVA_DB.Delete(&questionBank_judge).Error
	return err
}

// DeleteQuestionBankJudgeByIds 批量删除QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_judgeService *JudgeService) DeleteQuestionBankJudgeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]questionBank.Judge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuestionBankJudge 更新QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_judgeService *JudgeService) UpdateQuestionBankJudge(questionBank_judge questionBank.Judge) (err error) {
	err = global.GVA_DB.Save(&questionBank_judge).Error
	return err
}

// GetQuestionBankJudge 根据id获取QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_judgeService *JudgeService) GetQuestionBankJudge(id uint) (questionBank_judge questionBank.Judge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBank_judge).Error
	return
}

// GetQuestionBankJudgeInfoList 分页获取QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBank_judgeService *JudgeService) GetQuestionBankJudgeInfoList(info questionBankReq.QuestionBankJudgeSearch) (list []questionBank.Judge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Judge{})
	var questionBank_judges []questionBank.Judge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Describe != "" {
		db = db.Where("describe LIKE ?", "%"+info.Describe+"%")
	}
	if info.Is_right != nil {
		db = db.Where("is_right = ?", info.Is_right)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&questionBank_judges).Error
	return questionBank_judges, total, err
}
