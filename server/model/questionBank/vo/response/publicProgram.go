package response

import (
	"github.com/prl26/exam-system/server/global"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type PublicProgramDetail struct {
	global.GVA_MODEL
	questionBank.BasicModel
	questionBankBo.LanguageSupports
	questionBankBo.ProgramCases
	questionBankBo.DefaultCodes
	questionBankBo.ReferenceAnswers
}

type PublicProgramSimple struct {
	global.GVA_MODEL
	questionBank.SimpleModel
}
