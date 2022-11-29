package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
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

func (p *PublicProgramService) Update(t *questionBankPo.Program) error {
	return global.GVA_DB.Updates(t).Error
}
