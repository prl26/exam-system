package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
)

type TargetService struct {
}

// Create 创建QuestionBankRangTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *TargetService) Create(RangTopic *questionBank.Target) error {
	return global.GVA_DB.Create(RangTopic).Error
}

// DeleteQuestionBankRangTopic 删除QuestionBankRangTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *TargetService) DeleteQuestionBankRangTopic(RangTopic questionBank.Target) (err error) {
	err = global.GVA_DB.Delete(&RangTopic).Error
	return err
}

// Delete 批量删除QuestionBankRangTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *TargetService) Delete(ids request.IdsReq) error {
	return global.GVA_DB.Delete(&[]questionBank.Target{}, "id in ?", ids.Ids).Error
}

// Update 更新QuestionBankRangTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *TargetService) Update(RangTopic questionBank.Target) (err error) {
	err = global.GVA_DB.Updates(&RangTopic).Error
	return err
}

// GetQuestionBankRangTopic 根据id获取QuestionBankRangTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *TargetService) GetQuestionBankRangTopic(id uint) (RangTopic questionBank.Target, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&RangTopic).Error
	return
}

// FindTargetList 分页获取QuestionBankRangTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *TargetService) FindTargetList(criteria questionBankBo.TargetSearchCriteria, info request.PageInfo) (list []questionBankVoResp.TargetSimple, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Target{})
	if criteria.LessonId != 0 {
		db = db.Where("lesson_id=?", criteria.LessonId)
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
	if criteria.ChapterId != 0 {
		db = db.Where("chapter_id =?", criteria.ChapterId)
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

func (service *TargetService) FindDetail(id uint) (RangTopic *questionBankBo.TargetDetail, err error) {
	RangTopic = &questionBankBo.TargetDetail{}
	err = global.GVA_DB.Preload("Chapter").Preload("Knowledge").Model(&questionBank.Target{}).First(RangTopic, id).Error
	return
}

func (service *TargetService) FindTargetPracticeList(knowledge questionBankBo.TargetPracticeCriteria, info request.PageInfo) (list []*questionBankVoResp.TargetSimplePractice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&questionBank.Target{})
	db = db.Where("is_check=?", 1)
	db = db.Where("can_practice=?", 1)
	if knowledge.KnowledgeId != 0 {
		db = db.Where("knowledge_id=?", knowledge.KnowledgeId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}
