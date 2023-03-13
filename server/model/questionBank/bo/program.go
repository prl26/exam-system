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
	*LanguageSupports `json:"languageSupports"`
	*DefaultCodes     `json:"defaultCodes"`
	*ReferenceAnswers `json:"referenceAnswers"`
	*ProgramCases     `json:"programCases"`
}

type ProgramDetail struct {
	global.GVA_MODEL
	questionBankPo.CourseSupport
	CourseSupportPtr
	questionBankPo.ProgramModel
}

type ProgramPracticeCriteria struct {
	questionBankPo.CourseSupport
}

type ProgramPractice struct {
	PracticeModel
	DefaultCodes          string `json:"defaultCodes"`
	LanguageSupportsBrief string `json:"languageSupportBrief"`
}
type ProgramPractice1 struct {
	PracticeModel
	LanguageSupportsBrief string `json:"languageSupportBrief"`
}
