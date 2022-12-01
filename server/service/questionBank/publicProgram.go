package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankEnum "github.com/prl26/exam-system/server/model/questionBank/enum"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
)

type PublicProgramService struct{}

func (p *PublicProgramService) Create(programCase *questionBankPo.PublicProgram) error {
	return global.GVA_DB.Create(programCase).Error
}

func (p *PublicProgramService) FindList(criteria questionBankBo.PublicProgramSearchCriteria, info request.PageInfo) (data []questionBankVoResp.PublicProgramSimple, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBankPo.PublicProgram{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if criteria.Title != "" {
		db = db.Where("title LIKE ?", "%"+criteria.Title+"%")
	}
	if criteria.CanPractice != nil {
		db = db.Where("can_practice = ?", criteria.CanPractice)
	}
	if criteria.CanExam != nil {
		db = db.Where("can_exam = ?", criteria.CanExam)
	}
	if criteria.ProblemType != 0 {
		db = db.Where("problem_type =?", criteria.ProblemType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Model(&questionBankPo.PublicProgram{}).Limit(limit).Offset(offset).Find(&data).Error
	return
}

func (p *PublicProgramService) FindDetail(id int) (result *questionBankPo.PublicProgram, err error) {
	result = &questionBankPo.PublicProgram{}
	err = global.GVA_DB.Where("id=?", id).First(result).Error
	return
}

func (p *PublicProgramService) Update(t *questionBankPo.PublicProgram) error {
	return global.GVA_DB.Updates(t).Error
}

func (p *PublicProgramService) Migrate(ids []uint, migration questionBankBo.PublicProgramMigration) error {
	n := len(ids)
	programs := make([]*questionBankPo.Program, 0, n)
	table := map[questionBankEnum.LanguageType]bool{}
	if len(migration.LanguageIds) != 0 {
		for _, id := range migration.LanguageIds {
			table[id] = true
		}
	}
	if len(ids) == 1 {
		program := &questionBankPo.Program{}
		if err := global.GVA_DB.Model(&questionBankPo.PublicProgram{}).Select("*").First(program, ids[0]).Error; err != nil {
			return err
		}
		programs = append(programs, program)
	} else {
		if err := global.GVA_DB.Model(&questionBankPo.PublicProgram{}).Select("*").Where("id in ?", ids).Find(&programs).Error; err != nil {
			return err
		}
		if len(programs) != len(ids) {

		}
	}
	for _, program := range programs {
		program.GVA_MODEL = global.GVA_MODEL{}
		program.CourseSupport = migration.CourseSupport
		err := p.buildLanguageSupport(program, table)
		if err != nil {
			return err
		}
	}
	if err := global.GVA_DB.Create(&programs).Error; err != nil {
		return err
	}
	return nil
}

func (p *PublicProgramService) buildLanguageSupport(program *questionBankPo.Program, table map[questionBankEnum.LanguageType]bool) error {
	if len(table) != 0 {
		if program.LanguageSupports != "" {
			defaultCode := questionBankBo.LanguageSupports{}
			err := defaultCode.Deserialization(program.LanguageSupports)
			if err != nil {
				global.GVA_LOG.Sugar().Errorf("迁移失败%s", program.ID)
				return err
			}
			defaultCode.Filter(table)
			serialize, err := defaultCode.Serialize()
			if err != nil {
				return err
			}
			program.LanguageSupports = serialize
			program.LanguageSupportsBrief = defaultCode.Brief()
		}
		if program.DefaultCodes != "" {
			defaultCode := questionBankBo.DefaultCodes{}
			err := defaultCode.Deserialization(program.DefaultCodes)
			if err != nil {
				global.GVA_LOG.Sugar().Errorf("迁移失败%s", program.ID)
				return err
			}
			defaultCode.Filter(table)
			serialize, err := defaultCode.Serialize()
			if err != nil {
				return err
			}
			program.DefaultCodes = serialize
		}
		if program.ReferenceAnswers != "" {
			defaultCode := questionBankBo.ReferenceAnswers{}
			err := defaultCode.Deserialization(program.ReferenceAnswers)
			if err != nil {
				global.GVA_LOG.Sugar().Errorf("迁移失败%s", program.ID)
				return err
			}
			defaultCode.Filter(table)
			serialize, err := defaultCode.Serialize()
			if err != nil {
				return err
			}
			program.ReferenceAnswers = serialize
		}

	} else {
		if program.LanguageSupports != "" {
			defaultCode := questionBankBo.LanguageSupports{}
			err := defaultCode.Deserialization(program.LanguageSupports)
			if err != nil {
				global.GVA_LOG.Sugar().Errorf("迁移失败%s", program.ID)
				return err
			}
			serialize, err := defaultCode.Serialize()
			if err != nil {
				return err
			}
			program.LanguageSupports = serialize
			program.LanguageSupportsBrief = defaultCode.Brief()
		}
	}
	return nil
}

func (p *PublicProgramService) Delete(uints []uint) error {
	if len(uints) == 1 {
		return global.GVA_DB.Delete(&questionBankPo.PublicProgram{}, uints[0]).Error
	} else {
		return global.GVA_DB.Delete(&questionBankPo.PublicProgram{}, uints).Error
	}
}
