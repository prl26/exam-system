package questionBank

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	responese "github.com/prl26/exam-system/server/model/questionBank/response"
	"gorm.io/gorm"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 13:31

 * @Note:

 **/
type ProgrammService struct {
}

//
//// FindProgramms 分页获取 程序题
//func (p *ProgrammService) FindProgramms(info request.PageInfo)  {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//
//}
var lessonSupportSql = `
	select c.*,d.lesson_id,d.chapter_name,d.lesson_name
	from les_questionbank_chapter_merge c
	left join(SELECT
			b.id AS 'chapter_id',
			c.id AS 'lesson_id',
			b.name AS 'chapter_name',
			c.name AS 'lesson_name' 
		FROM
		bas_chapter b
		LEFT JOIN bas_lesson c ON b.lesson_id = c.id
		where ISNULL(b.deleted_at)
		) as d
	on c.chapter_id =d.chapter_id
	where c.question_id= ?
	and c.question_type= ?
	and ISNULL(c.deleted_at)
`

// FindProgramDetail 获取 编程题的详细
func (p *ProgrammService) FindProgramDetail(prgramm *questionBank.Programm, programId uint) error {
	if err := global.GVA_DB.Where("id", programId).Find(&prgramm).Error; err != nil {
		return err
	}
	return nil
}
func (p *ProgrammService) EditProgrammDetail(prgramm *questionBank.Programm) error {
	if err := global.GVA_DB.Updates(prgramm).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProgrammService) DeleteProgramm(ints []int) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := global.GVA_DB.Where("id in ?", ints).Delete(&questionBank.Programm{}).Error; err != nil {
			return err
		}
		if err := global.GVA_DB.Where("programm_id in ? ", ints).Delete(&questionBank.ProgrammLanguageMerge{}).Error; err != nil {
			return err
		}
		if err := global.GVA_DB.Where("question_id in ? and question_type=?", ints, questionType.PROGRAM).Delete(&questionBank.ChapterMerge{}).Error; err != nil {
			return err
		}
		if err := global.GVA_DB.Where("programm_id in ?", ints).Delete(&questionBank.ProgrammCase{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//FindLanguageSupport 获取 编程题的语言支持
func (p *ProgrammService) FindLanguageSupport(support *[]responese.LanguageSupport, programId uint) error {
	if err := global.GVA_DB.Raw(languageSupport, programId).Find(&support).Error; err != nil {
		return fmt.Errorf("无法找到语言支持")
	}
	return nil
}

func (p *ProgrammService) FindProgrammCases(cases *[]questionBank.ProgrammCase, programId uint, languageId int) error {
	if err := global.GVA_DB.Where("programm_id=? and language_id=?", programId, languageId).Find(cases).Error; err != nil {
		return fmt.Errorf("该编程题不支持该语言")
	}
	return nil
}

func (p *ProgrammService) AddProgrammCase(cases *[]questionBank.ProgrammCase) error {
	if err := global.GVA_DB.Create(cases).Error; err != nil {
		return fmt.Errorf("编程题添加失败")
	}
	return nil
}

func (p *ProgrammService) EditProgrammCases(cases []questionBank.ProgrammCase) error {
	err := global.GVA_DB.Transaction(
		func(tx *gorm.DB) error {
			for i := 0; i < len(cases); i++ {
				if err := tx.Model(cases[i]).Updates(cases[i]).Error; err != nil {
					return fmt.Errorf("更新失败")
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

func (p *ProgrammService) DeleteProgrammCases(ints []int) error {
	if err := global.GVA_DB.Delete(&questionBank.ProgrammCase{}, ints).Error; err != nil {
		return fmt.Errorf("编程题用例删除失败")
	}
	return nil
}

func (p *ProgrammService) AddLanguageSupport(merge *questionBank.ProgrammLanguageMerge, cases *[]questionBank.ProgrammCase) error {
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var result int64
		if err := tx.Model(&questionBank.Programm{}).Where("id", merge.ProgrammId).Count(&result).Error; err != nil {
			return fmt.Errorf("查找编程题失败")
		}
		if result == 0 {
			return fmt.Errorf("该编程题不存在")
		}
		if err := tx.Create(merge).Error; err != nil {
			return fmt.Errorf("创建语言支持失败")
		}
		if len(*cases) != 0 {
			if err := tx.Create(cases).Error; err != nil {
				return fmt.Errorf("创建编程用例失败")
			}
		}
		return nil
	})
	return err
}

func (*ProgrammService) EditLanguageSupport(merge *questionBank.ProgrammLanguageMerge) error {
	if err := global.GVA_DB.Updates(merge).Error; err != nil {
		return fmt.Errorf("更新失败")
	}
	return nil
}

func (*ProgrammService) DeleteLanguageSupport(prgrammId uint, ints []int) error {
	err := global.GVA_DB.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Where("prgramm_id =? and language_id in ?", prgrammId, ints).Error; err != nil {
				return fmt.Errorf("删除失败")
			}
			if err := tx.Where("prgramm_id =? and language_id in ?", prgrammId, ints).Error; err != nil {
				return fmt.Errorf("删除失败")
			}
			return nil
		})
	return err
}

func (p *ProgrammService) FindList(info questionBankReq.ProgramFindList) (list []questionBank.ProgrammView, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&questionBank.Programm{})
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
