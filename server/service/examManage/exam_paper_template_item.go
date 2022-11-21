package examManage

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/questionBank"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaperTemplateItemService struct {
}

// CreatePaperTemplateItem 创建PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) CreatePaperTemplateItem(paperTemplateItem examManage.PaperTemplateItem) (err error) {
	err = global.GVA_DB.Create(&paperTemplateItem).Error
	return err
}

// DeletePaperTemplateItem 删除PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) DeletePaperTemplateItem(paperTemplateItem examManage.PaperTemplateItem) (err error) {
	err = global.GVA_DB.Delete(&paperTemplateItem).Error
	return err
}

// DeletePaperTemplateItemByIds 批量删除PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) DeletePaperTemplateItemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.PaperTemplateItem{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePaperTemplateItem 更新PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) UpdatePaperTemplateItem(paperTemplateItem []examManage.PaperTemplateItem) (err error) {
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(&paperTemplateItem).Error
		if err != nil {
			return err
		}
		return err
	})
	return err
}

// GetPaperTemplateItem 根据id获取PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) GetPaperTemplateItem(id uint) (paperTemplateItem examManage.PaperTemplateItem, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&paperTemplateItem).Error
	return
}

// GetPaperTemplateItemInfoList 分页获取PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) GetPaperTemplateItemInfoList(info examManageReq.PaperTemplateItemSearch) (list []examManage.PaperTemplateItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.PaperTemplateItem{})
	var paperTemplateItems []examManage.PaperTemplateItem
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ChapterId != nil {
		db = db.Where("chapter = ?", info.ChapterId)
	}
	if info.ProblemType != nil {
		db = db.Where("problem_type = ?", info.ProblemType)
	}
	if info.QuestionType != nil {
		db = db.Where("difficulty = ?", info.QuestionType)
	}
	if info.Num != nil {
		db = db.Where("num = ?", info.Num)
	}
	if info.Score != nil {
		db = db.Where("score = ?", info.Score)
	}
	if info.TemplateId != nil {
		db = db.Where("template_id = ?", info.TemplateId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&paperTemplateItems).Error
	return paperTemplateItems, total, err
}
func (paperTemplateItemService *PaperTemplateItemService) GetPaperId(info examManage.PaperTemplateItem) (paperId uint) {
	var paper examManage.ExamPaper
	global.GVA_DB.Where("id = ?", info.TemplateId).Find(&paper)
	return paper.ID
}

func (paperTemplateItemService *PaperTemplateItemService) SetPaperChoiceQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.MultipleChoice
	num := info.Num
	err = global.GVA_DB.Raw("SELECT * FROM les_questionbank_supply_blank ORDER BY RAND()").Where("problem_type = ? and can_exam = ?", info.ProblemType, 1).Limit(*num).Find(&list).Error
	if err != nil {
		return
	} else {
		if len(list) > 0 {
			for j := 0; j < *num; j++ {
				questionMerge := examManage.PaperQuestionMerge{
					GVA_MODEL:    global.GVA_MODEL{},
					PaperId:      &Id,
					QuestionId:   &list[j].ID,
					Score:        info.Score,
					QuestionType: info.QuestionType,
					ProblemType:  info.ProblemType,
				}
				err = global.GVA_DB.Create(&questionMerge).Error
				if err != nil {
					return
				}
			}
		}
	}
	return
}
func (paperTemplateItemService *PaperTemplateItemService) SetPaperJudgeQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.Judge
	num := info.Num
	err = global.GVA_DB.Raw("SELECT * FROM les_questionbank_supply_blank ORDER BY RAND()").Where("problem_type = ? and can_exam = ?", info.ProblemType, 1).Limit(*num).Find(&list).Error
	if err != nil {
		return
	} else {
		if len(list) > 0 {
			for j := 0; j < *num; j++ {
				questionMerge := examManage.PaperQuestionMerge{
					GVA_MODEL:    global.GVA_MODEL{},
					PaperId:      &Id,
					QuestionId:   &list[j].ID,
					Score:        info.Score,
					QuestionType: info.QuestionType,
					ProblemType:  info.ProblemType,
				}
				err = global.GVA_DB.Create(&questionMerge).Error
				if err != nil {
					return
				}
			}
		}
	}
	return
}
func (paperTemplateItemService *PaperTemplateItemService) SetPaperBlankQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.SupplyBlank
	num := info.Num
	err = global.GVA_DB.Raw("SELECT * FROM les_questionbank_supply_blank ORDER BY RAND()").Where("problem_type = ? and can_exam = ?", info.ProblemType, 1).Limit(*num).Find(&list).Error
	if err != nil {
		return
	} else {
		if len(list) > 0 {
			for j := 0; j < *num; j++ {
				fmt.Println(j)
				questionMerge := examManage.PaperQuestionMerge{
					GVA_MODEL:    global.GVA_MODEL{},
					PaperId:      &Id,
					QuestionId:   &list[j].ID,
					Score:        info.Score,
					QuestionType: info.QuestionType,
					ProblemType:  info.ProblemType,
				}
				err = global.GVA_DB.Create(&questionMerge).Error
				if err != nil {
					return
				}
			}
		}
	}
	return
}
func (paperTemplateItemService *PaperTemplateItemService) SetPaperProgramQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.Programm
	num := info.Num
	err = global.GVA_DB.Raw("SELECT * FROM les_questionbank_supply_blank ORDER BY RAND()").Where("problem_type = ? and can_exam = ?", info.ProblemType, 1).Limit(*num).Find(&list).Error
	if err != nil {
		return
	} else {
		if len(list) > 0 {
			for j := 0; j < *num; j++ {
				questionMerge := examManage.PaperQuestionMerge{
					GVA_MODEL:    global.GVA_MODEL{},
					PaperId:      &Id,
					QuestionId:   &list[j].ID,
					Score:        info.Score,
					QuestionType: info.QuestionType,
					ProblemType:  info.ProblemType,
				}
				err = global.GVA_DB.Create(&questionMerge).Error
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func (paperTemplateItemService *PaperTemplateItemService) SetPaperQuestion(info []examManage.PaperTemplateItem, Id uint) (err error) {
	fmt.Println("jinru")
	for _, v := range info {
		if *v.QuestionType == questionType.MultipleChoice {
			go func() {
				err = paperTemplateItemService.SetPaperChoiceQuestion(v, Id)
				if err != nil {
					return
				}
			}()
		} else if *v.QuestionType == questionType.JUDGE {
			go func() {
				err = paperTemplateItemService.SetPaperJudgeQuestion(v, Id)
				if err != nil {
					return
				}
			}()
		} else if *v.QuestionType == questionType.SUPPLY_BLANK {
			err = paperTemplateItemService.SetPaperBlankQuestion(v, Id)
			if err != nil {
				return
			}
		} else if *v.QuestionType == questionType.PROGRAM {
			go func() {
				err = paperTemplateItemService.SetPaperProgramQuestion(v, Id)
				if err != nil {
					return
				}
			}()
		}
	}
	return
}
