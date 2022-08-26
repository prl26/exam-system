package questionBank

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
	questionBankReq "exam-system/model/questionBank/request"
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
	if info.ChapterId != nil {
		db = db.Where("chapter_id =?", info.ChapterId)
	}
	if info.ProblemType != nil {
		db = db.Where("problem_type = ?", info.ProblemType)
	}
	if info.Title != "" {
		db = db.Where("title like ?", "%"+info.Title+"%")
	}
	if info.CanExam != nil {
		db = db.Where("can_exam = ?", info.CanExam)
	}
	if info.CanPractice != nil {
		db = db.Where("can_practice = ?", info.CanPractice)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&questionBank_judges).Error
	return questionBank_judges, total, err
}
