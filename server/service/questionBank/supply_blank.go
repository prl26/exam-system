package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	"gorm.io/gorm"
)

type SupplyBlankService struct {
}

func (service *SupplyBlankService) Create(supplyBlank *questionBank.SupplyBlank, courseSupport []uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(supplyBlank).Error; err != nil {
			return err
		}
		courseSupport := buildCourseSupport(courseSupport, supplyBlank.ID, questionType.SUPPLY_BLANK)
		if err := tx.Create(&courseSupport).Error; err != nil {
			return err
		}
		return nil
	})
}

func (service *SupplyBlankService) Delete(ids request.IdsReq) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&[]questionBank.SupplyBlank{}, "id in ?", ids.Ids).Error; err != nil {
			global.GVA_LOG.Sugar().Errorf("SupplyBlankService.Delete:%s", err.Error())
			return err
		}
		if err := tx.Delete(&[]questionBank.ChapterMerge{}, "question_id in ? and question_type =?", ids, questionType.SUPPLY_BLANK).Error; err != nil {
			return err
		}
		return nil
	})
}

func (service *SupplyBlankService) UpdateQuestionBankSupplyBlank(questionBank_supply_blank questionBank.SupplyBlank) (err error) {
	return global.GVA_DB.Updates(&questionBank_supply_blank).Error
}

func (service *SupplyBlankService) GetQuestionBankSupplyBlank(id uint) (questionBank_supply_blank questionBank.SupplyBlank, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBank_supply_blank).Error
	return
}

func (service *SupplyBlankService) FindList(info questionBankReq.QuestionBankSupplyBlankSearch) (list []questionBank.SupplyBlankView, total int64, err error) {
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
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Model(&questionBank.SupplyBlankView{}).Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (service *SupplyBlankService) FindDetail(j *questionBank.SupplyBlank, id uint) error {
	return global.GVA_DB.Where("id = ?", id).Find(j).Error
}
