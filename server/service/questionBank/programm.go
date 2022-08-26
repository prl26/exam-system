package questionBank

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
	questionBankReq "exam-system/model/questionBank/request"
	"gorm.io/gorm"
)

type ProgrammService struct {
}

// CreateQuestionBankProgramm 创建QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *ProgrammService) CreateQuestionBankProgramm(questionBankProgramm questionBank.Programm) (err error) {
	err = global.GVA_DB.Create(&questionBankProgramm).Error
	return err
}

// DeleteQuestionBankProgramm 删除QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *ProgrammService) DeleteQuestionBankProgramm(questionBankProgramm questionBank.Programm) (err error) {
	err = global.GVA_DB.Delete(&questionBankProgramm).Error
	return err
}

// DeleteQuestionBankProgrammByIds 批量删除QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *ProgrammService) DeleteQuestionBankProgrammByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]questionBank.Programm{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateQuestionBankProgramm 更新QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *ProgrammService) UpdateQuestionBankProgramm(questionBankProgramm questionBank.Programm) (err error) {
	err = global.GVA_DB.Save(&questionBankProgramm).Error
	return err
}

// GetQuestionBankProgramm 根据id获取QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *ProgrammService) GetQuestionBankProgramm(id uint) (questionBankProgramm questionBank.Programm, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&questionBankProgramm).Error
	return
}

// GetQuestionBankProgrammInfoList 分页获取QuestionBankProgramm记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionBankProgrammService *ProgrammService) GetQuestionBankProgrammInfoList(info questionBankReq.QuestionBankProgrammSearch) (list []questionBank.Programm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.Programm{})
	var questionBankProgramms []questionBank.Programm
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
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&questionBankProgramms).Error
	return questionBankProgramms, total, err
}

func (questionBankProgrammService *ProgrammService) AddLanguageSupport(languages []*questionBankReq.SupportLanguage, programmsId uint) error {
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		n := len(languages)
		id := int(programmsId)
		for i := 0; i < n; i++ {
			languages[i].ProgrammId = &id
			for _, programmCase := range languages[i].Cases {
				programmCase.ProgrammId = &id
				programmCase.ID = 0
				programmCase.LanguageId = languages[i].LanguageId
			}
			create := global.GVA_DB.Create(languages[i].Cases)
			if create.Error != nil {
				return create.Error
			}
		}
		create := global.GVA_DB.Create(languages)
		if create.Error != nil {
			return create.Error
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
