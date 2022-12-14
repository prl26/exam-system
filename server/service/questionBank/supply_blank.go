package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"

	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type SupplyBlankService struct {
}

func (service *SupplyBlankService) Create(supplyBlank *questionBank.SupplyBlank) error {
	return global.GVA_DB.Create(supplyBlank).Error

}

func (service *SupplyBlankService) Delete(ids request.IdsReq) error {
	return global.GVA_DB.Delete(&[]questionBank.SupplyBlank{}, "id in ?", ids.Ids).Error
}

func (service *SupplyBlankService) UpdateQuestionBankSupplyBlank(questionBank_supply_blank questionBank.SupplyBlank) (err error) {
	return global.GVA_DB.Updates(&questionBank_supply_blank).Error
}

func (service *SupplyBlankService) GetQuestionBankSupplyBlank(id uint) (questionBank_supply_blank questionBank.SupplyBlank, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBank_supply_blank).Error
	return
}

func (service *SupplyBlankService) FindList(info questionBankReq.QuestionBankSupplyBlankSearch) (list []questionBankVoResp.JudgeSimple, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.SupplyBlank{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.CanExam != nil {
		db = db.Where("can_exam= ?", info.CanExam)
	}
	if info.CanPractice != nil {
		db = db.Where("can_practice= ?", info.CanPractice)
	}
	if info.ProblemType != 0 {
		db = db.Where("problem_type= ?", info.ProblemType)
	}
	if info.ChapterId != 0 {
		db = db.Where("chapter_id =?", info.ChapterId)
	}
	if info.KnowledgeId != 0 {
		db = db.Where("knowledge_id=?", info.KnowledgeId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Model(&questionBank.SupplyBlank{}).Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *SupplyBlankService) FindDetail(id uint) (j *questionBankBo.SupplyBlankDetail, err error) {
	j = &questionBankBo.SupplyBlankDetail{}
	err = global.GVA_DB.Preload("Chapter").Preload("Knowledge").Model(&questionBank.SupplyBlank{}).First(j, id).Error
	return
}
