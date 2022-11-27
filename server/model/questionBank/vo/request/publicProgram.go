package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"

	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type PublicProgramCreate struct {
	questionBankPo.BasicModel
	LanguageSupports questionBankBo.LanguageSupports `json:"languageSupports,omitempty"`
	ProgramCases     questionBankBo.ProgramCases     `json:"programCases,omitempty"`
	ReferenceAnswers questionBankBo.ReferenceAnswers `json:"referenceAnswers"`
	DefaultCodes     questionBankBo.DefaultCodes     `json:"defaultCodes"`
}

type PublicProgramUpdate struct {
	Id uint `json:"id"`
	questionBankPo.BasicModel
	LanguageSupports questionBankBo.LanguageSupports `json:"languageSupports,omitempty"`
	ProgramCases     questionBankBo.ProgramCases     `json:"programCases,omitempty"`
}

type PublicProgramSearch struct {
	questionBankBo.PublicProgramSearchCriteria
	request.PageInfo
}
