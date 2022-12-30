package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
)

type MultipleChoiceService struct {
}

func (a *MultipleChoiceService) Create(multipleChoice *questionBank.MultipleChoice) (err error) {
	return global.GVA_DB.Create(multipleChoice).Error
}

func (a *MultipleChoiceService) Delete(ids request.IdsReq) error {
	return global.GVA_DB.Delete(&[]questionBank.MultipleChoice{}, "id in ?", ids.Ids).Error
}

func (a *MultipleChoiceService) Update(multipleChoice questionBank.MultipleChoice) error {
	return global.GVA_DB.Updates(multipleChoice).Error
}

func (a *MultipleChoiceService) FindDetail(id uint) (result *questionBankBo.MultipleDetail, err error) {
	result = &questionBankBo.MultipleDetail{}
	err = global.GVA_DB.Preload("Chapter").Preload("Knowledge").Model(&questionBank.MultipleChoice{}).First(result, id).Error
	return
}

func (a *MultipleChoiceService) FindList(criteria questionBankBo.MultipleCriteria, info request.PageInfo) (list []questionBankResp.MultipleChoiceSimple, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.MultipleChoice{})
	if criteria.IsCheck != nil {
		db = db.Where("is_check=?", *criteria.IsCheck)
	}
	if criteria.LessonId != 0 {
		db = db.Where("lesson_id=?", criteria.LessonId)
	}
	if criteria.IsIndefinite == 1 {
		db = db.Where("is_indefinite = 1")
	} else {
		db = db.Where("is_indefinite = 0")
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
