package response

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/lessondata"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type ProgramDetail struct {
	global.GVA_MODEL
	Chapter   *basicdata.Chapter
	Knowledge *lessondata.Knowledge
	questionBankBo.ProgramOjSupport
	questionBank.BasicModel
}

type ProgramSimple struct {
	global.GVA_MODEL
	questionBank.SimpleModel
}

type ProgramPractice struct {
	questionBankBo.PracticeModel
	questionBankBo.DefaultCodes `json:"defaultCodes"`
}

func (p *ProgramPractice) Convert(practice *questionBankBo.ProgramPractice) error {
	p.PracticeModel = practice.PracticeModel
	return p.DefaultCodes.DeserializationWithBrief(practice.DefaultCodes, practice.LanguageSupportsBrief)
}
