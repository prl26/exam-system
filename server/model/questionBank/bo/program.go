package bo

import (
	"github.com/prl26/exam-system/server/global"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type ProgramSearchCriteria struct {
	questionBankPo.SimpleModel
	questionBankPo.CourseSupport
}

type ProgramOjSupport struct {
	LanguageSupports
	ProgramCases
	ReferenceAnswers
	DefaultCodes
}

type ProgramDetail struct {
	global.GVA_MODEL
	questionBankPo.CourseSupport
	CourseSupportPtr
	questionBankPo.ProgramModel
}
