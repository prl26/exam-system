package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
)

type PaperTemplateService struct {
}

// CreatePaperTemplate 创建PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService) CreatePaperTemplate(Papertemplate examManage.PaperTemplate) (err error) {
	err = global.GVA_DB.Create(&Papertemplate).Error
	return err
}

// DeletePaperTemplate 删除PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService) DeletePaperTemplate(Id uint) (err error) {
	err = global.GVA_DB.Where("id = ?", Id).Delete(&examManage.PaperTemplate{}).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Where("template_id = ?", Id).Delete(&examManage.PaperTemplateItem{}).Error
	return
}

// DeletePaperTemplateByIds 批量删除PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService) DeletePaperTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.PaperTemplate{}, "id in ?", ids.Ids).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Delete(&[]examManage.PaperTemplateItem{}, "template_id in ?", ids.Ids).Error
	return err
}

// UpdatePaperTemplate 更新PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService) UpdatePaperTemplate(Papertemplate examManage.PaperTemplate) (err error) {
	//global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	//	err = global.GVA_DB.Table("exam_paper_template").Where("id = ?", Papertemplate.ID).Updates(&Papertemplate).Error
	//	err = tx.Clauses(clause.OnConflict{
	//		Columns:   []clause.Column{{Name: "id"}},
	//		UpdateAll: true,
	//	}).Create(&paperTemplateItem).Error
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//err = global.GVA_DB.Create(&Papertemplate).Error
	err = global.GVA_DB.Updates(&Papertemplate).Error
	for i := 0; i < len(Papertemplate.PaperTemplateItems); i++ {
		global.GVA_DB.Save(&Papertemplate.PaperTemplateItems[i])
	}
	var IdOfItems []uint
	global.GVA_DB.Model(&examManage.PaperTemplateItem{}).Select("id").Where("template_id  = ?", Papertemplate.ID).Find(&IdOfItems)
	set := make(map[uint]bool)
	for _, v := range Papertemplate.PaperTemplateItems {
		set[v.ID] = true
	}
	for _, v := range IdOfItems {
		_, ok := set[v]
		if !ok {
			global.GVA_DB.Where("id = ?", v).Delete(&examManage.PaperTemplateItem{})
		}
	}
	return err
}

// GetPaperTemplate 根据id获取PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService) GetPaperTemplate(id uint) (Papertemplate examManage.PaperTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&Papertemplate).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Where("template_id = ?", Papertemplate.ID).Find(&Papertemplate.PaperTemplateItems).Error
	return
}

// GetPaperTemplateInfoList 分页获取PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService) GetPaperTemplateInfoList(info examManageReq.PaperTemplateSearch, userId int, authorityId uint) (list []examManage.PaperTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.PaperTemplate{})
	var Papertemplates []examManage.PaperTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.LessonId != 0 {
		db = db.Where("lesson_id = ?", info.LessonId)
	}
	if authorityId != 888 {
		db = db.Where("user_id = ?", userId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc,updated_at desc ").Limit(limit).Offset(offset).Find(&Papertemplates).Error
	return Papertemplates, total, err
}

//查找该课程下有哪些章节,章节下面各题目难度的题目数目
func (PapertemplateService *PaperTemplateService) GetDetails(lessonId uint) (templates response.Template, err error) {
	err = global.GVA_DB.Raw("select b.id as chapter_id,b.`name` as chapter_name,problem_type,count(j.id) as Num\nFROM bas_chapter as b,les_questionbank_multiple_choice as j\nWHERE  b.lesson_id = ? and b.id = j.chapter_id and j.deleted_at is null\ngroup by b.id,b.`name`,problem_type\nORDER BY b.`name`\n", lessonId).
		Scan(&templates.Choice).Error
	if err != nil {
		return
	}

	err = global.GVA_DB.Raw("select b.id as chapter_id,b.`name` as chapter_name,problem_type,count(j.id) as Num\nFROM bas_chapter as b,les_questionbank_judge as j\nWHERE  b.lesson_id = ? and b.id = j.chapter_id and j.deleted_at is null\ngroup by b.id,b.`name`,problem_type\nORDER BY b.`name`\n", lessonId).
		Scan(&templates.Judge).Error
	if err != nil {
		return
	}

	err = global.GVA_DB.Raw("select b.id as chapter_id,b.`name` as chapter_name,problem_type,count(j.id) as Num\nFROM bas_chapter as b,les_questionbank_supply_blank as j\nWHERE  b.lesson_id = ? and b.id = j.chapter_id and j.deleted_at is null\ngroup by b.id,b.`name`,problem_type\nORDER BY b.`name`\n", lessonId).
		Scan(&templates.Blank).Error
	if err != nil {
		return
	}

	err = global.GVA_DB.Raw("select b.id as chapter_id,b.`name` as chapter_name,problem_type,count(j.id) as Num\nFROM bas_chapter as b,les_questionbank_programm as j\nWHERE  b.lesson_id = ? and b.id = j.chapter_id and j.deleted_at is null\ngroup by b.id,b.`name`,problem_type\nORDER BY b.`name`\n", lessonId).
		Scan(&templates.Program).Error
	if err != nil {
		return
	}
	return
}
func (PapertemplateService *PaperTemplateService) CheckPaperTemplate(info []examManage.PaperTemplateItem) (IsOk bool, err error) {
	for _, v := range info {
		if *v.QuestionType == questionType.JUDGE {
			if IsOk, err = PapertemplateService.CheckPaperJudgeTemplate(v); err != nil {
				return false, err
			} else if IsOk != true {
				return false, err
			}
		}
		if *v.QuestionType == questionType.SUPPLY_BLANK {
			if IsOk, err = PapertemplateService.CheckPaperBlankTemplate(v); err != nil {
				return false, err
			} else if IsOk != true {
				return false, err
			}
		}
		if *v.QuestionType == questionType.PROGRAM {
			if IsOk, err = PapertemplateService.CheckPaperProgramTemplate(v); err != nil {
				return false, err
			} else if IsOk != true {
				return false, err
			}
		}
		if *v.QuestionType == questionType.SINGLE_CHOICE {
			if IsOk, err = PapertemplateService.CheckPaperChoiceTemplate(v); err != nil {
				return false, err
			} else if IsOk != true {
				return false, err
			}
		}
	}
	return true, err
}
func (PapertemplateService *PaperTemplateService) CheckPaperJudgeTemplate(info examManage.PaperTemplateItem) (IsOk bool, err error) {
	var count int64
	err = global.GVA_DB.Table("les_questionbank_judge").Where("problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null ", info.ProblemType, 1, info.ChapterId).Count(&count).Error
	temp := int(count)
	if temp < *info.Num {
		IsOk = false
		return
	}
	return true, err
}
func (PapertemplateService *PaperTemplateService) CheckPaperBlankTemplate(info examManage.PaperTemplateItem) (IsOk bool, err error) {
	var count int64
	err = global.GVA_DB.Table("les_questionbank_supply_blank").Where("problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null ", info.ProblemType, 1, info.ChapterId).Count(&count).Error
	temp := int(count)
	if temp < *info.Num {
		IsOk = false
		return
	}
	return true, err
}
func (PapertemplateService *PaperTemplateService) CheckPaperProgramTemplate(info examManage.PaperTemplateItem) (IsOk bool, err error) {
	var count int64
	err = global.GVA_DB.Table("les_questionbank_programm").Where("problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null", info.ProblemType, 1, info.ChapterId).Count(&count).Error
	temp := int(count)
	if temp < *info.Num {
		IsOk = false
		return
	}
	return true, err
}
func (PapertemplateService *PaperTemplateService) CheckPaperChoiceTemplate(info examManage.PaperTemplateItem) (IsOk bool, err error) {
	var count int64
	err = global.GVA_DB.Table("les_questionbank_multiple_choice").Where("problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null", info.ProblemType, 1, info.ChapterId).Count(&count).Error
	temp := int(count)
	if temp < *info.Num {
		IsOk = false
		return
	}
	return true, err
}
