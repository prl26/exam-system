package questionBank

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
)

type ProgramService struct {
}

func (s ProgramService) Create(p *po.Program) error {
	return global.GVA_DB.Create(p).Error
}

func (p *ProgramService) FindList(criteria questionBankBo.ProgramSearchCriteria, info request.PageInfo) (data []questionBankVoResp.PublicProgramSimple, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&po.Program{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if criteria.LessonId != 0 {
		db = db.Where("lesson_id=?", criteria.LessonId)
	}
	if criteria.IsCheck != nil {
		db = db.Where("is_check=?", *criteria.IsCheck)
	}
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
	err = db.Model(&po.Program{}).Limit(limit).Offset(offset).Find(&data).Error
	return
}

func (p *ProgramService) FindDetail(id int) (result *questionBankBo.ProgramDetail, err error) {
	result = &questionBankBo.ProgramDetail{}
	err = global.GVA_DB.Preload("Chapter").Preload("Knowledge").Model(&questionBankPo.Program{}).First(result, id).Error
	return
}

func (p *ProgramService) Update(t *po.Program) error {
	return global.GVA_DB.Updates(t).Error
}

func (p *ProgramService) Delete(uints []uint) error {
	if len(uints) == 1 {
		return global.GVA_DB.Delete(&questionBankPo.Program{}, uints[0]).Error
	} else {
		return global.GVA_DB.Delete(&questionBankPo.Program{}, uints).Error
	}
}
