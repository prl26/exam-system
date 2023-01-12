package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
)

type JudgeService struct {
}

// Create 创建QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) Create(judge *questionBank.Judge) error {
	return global.GVA_DB.Create(judge).Error
}

// DeleteQuestionBankJudge 删除QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) DeleteQuestionBankJudge(judge questionBank.Judge) (err error) {
	err = global.GVA_DB.Delete(&judge).Error
	return err
}

// Delete 批量删除QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) Delete(ids request.IdsReq) error {
	return global.GVA_DB.Delete(&[]questionBank.Judge{}, "id in ?", ids.Ids).Error
}

// Update 更新QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) Update(judge questionBank.Judge) (err error) {
	err = global.GVA_DB.Updates(&judge).Error
	return err
}

// GetQuestionBankJudge 根据id获取QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) GetQuestionBankJudge(id uint) (judge questionBank.Judge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&judge).Error
	return
}

// FindJudgeList 分页获取QuestionBankJudge记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *JudgeService) FindJudgeList(criteria questionBankBo.JudgeSearchCriteria, info request.PageInfo) (list []questionBankVoResp.JudgeSimple, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Judge{})
	if criteria.ChapterId != 0 {
		db = db.Where("chapter_id =?", criteria.ChapterId)
	} else {
		if criteria.LessonId != 0 {
			db = db.Where("lesson_id=?", criteria.LessonId)
		}
	}
	if criteria.IsCheck != nil {
		db = db.Where("is_check=?", criteria.IsCheck)
	}
	if criteria.ProblemType != 0 {
		db = db.Where("problem_type = ?", criteria.ProblemType)
	}
	if criteria.Title != "" {
		db = db.Where("title like ?", "%"+criteria.Title+"%")
	}
	if criteria.CanExam != nil {
		db = db.Where("can_exam = ?", criteria.CanExam)
	}
	if criteria.CanPractice != nil {
		db = db.Where("can_practice = ?", criteria.CanPractice)
	}
	if criteria.KnowledgeId != 0 {
		db = db.Where("knowledge_id=?", criteria.KnowledgeId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *JudgeService) FindDetail(id uint) (judge *questionBankBo.JudgeDetail, err error) {
	judge = &questionBankBo.JudgeDetail{}
	err = global.GVA_DB.Preload("Chapter").Preload("Knowledge").Model(&questionBank.Judge{}).First(judge, id).Error
	return
}
