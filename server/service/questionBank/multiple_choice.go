package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	"gorm.io/gorm"
)

type MultipleChoiceService struct {
}

func (a *MultipleChoiceService) Create(multipleChoice *questionBank.MultipleChoice, chapterSupport []uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(multipleChoice).Error; err != nil {
			return err
		}
		for i := 0; i < len(multipleChoice.Options); i++ {
			multipleChoice.Options[i].MultipleChoiceId = multipleChoice.ID
		}
		if len(chapterSupport) != 0 {
			courseSupport := buildCourseSupport(chapterSupport, multipleChoice.ID, questionType.MULTIPLE_CHOICE)
			if err := tx.Create(&courseSupport).Error; err != nil {
				return err
			}
		}
		return tx.Create(&multipleChoice.Options).Error
	})
}

func (a *MultipleChoiceService) Delete(ids request.IdsReq) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&[]questionBank.MultipleChoice{}, "id in ?", ids.Ids).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]questionBank.Options{}, "multiple_choice_id in", ids.Ids).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]questionBank.ChapterMerge{}, "question_id in ? and question_type=?", ids, questionType.MULTIPLE_CHOICE).Error; err != nil {
			return err
		}
		return nil
	})
}

func (a *MultipleChoiceService) Update(multipleChoice questionBank.MultipleChoice) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(multipleChoice).Error; err != nil {
			return err
		}
		if err := tx.Where("multiple_choice_id=?", multipleChoice.ID).Delete(&questionBank.Options{}).Error; err != nil {
			return err
		}
		return tx.Create(multipleChoice.Options).Error
	})
}

func (a *MultipleChoiceService) FindDetail(questionBankMultipleChoice *questionBank.MultipleChoice, id uint) error {
	return global.GVA_DB.Where("id = ?", id).Preload("CreditCards").First(questionBankMultipleChoice).Error
}

func (a *MultipleChoiceService) FindList(info questionBankReq.MultipleChoiceFindList) (list []questionBank.MultipleChoiceView, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.MultipleChoice{})

	if info.ProblemType != 0 {
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
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}
