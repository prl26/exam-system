package po

import (
	"github.com/prl26/exam-system/server/global"
)

// Program 结构体
type Program struct {
	global.GVA_MODEL
	ProgramModel
	CourseSupport
}

type ProgramModel struct {
	BasicModel
	ProgramCases     string `json:"programCases"`
	LanguageSupports string `json:"languageSupport"`
	ReferenceAnswers string `json:"referenceAnswer"`
	DefaultCodes     string `json:"defaultCodes"`
}

// TableName Program 表名
func (Program) TableName() string {
	return "les_questionBank_programm"
}
