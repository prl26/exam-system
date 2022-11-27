package response

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type PublicProgramDetail struct {
	global.GVA_MODEL
	Chapter   *basicdata.Chapter
	Knowledge *basicdata.Knowledge
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
