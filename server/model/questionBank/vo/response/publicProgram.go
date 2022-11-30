package response

import (
	"github.com/prl26/exam-system/server/global"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type PublicProgramDetail struct {
	global.GVA_MODEL
	questionBank.BasicModel
	questionBankBo.LanguageSupports `json:"languageSupports"`
	questionBankBo.ProgramCases     `json:"programCases"`
	questionBankBo.DefaultCodes     `json:"defaultCodes"`
	questionBankBo.ReferenceAnswers `json:"referenceAnswers"`
}

type PublicProgramSimple struct {
	global.GVA_MODEL
	questionBank.SimpleModel
}
