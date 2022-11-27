package response

import (
	"github.com/prl26/exam-system/server/global"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type MultipleChoiceSimple struct {
	global.GVA_MODEL
	questionBank.SimpleModel
}

type MultipleChoicePractice struct {
	PracticeModel
	MostOptions int `json:"mostOptions" form:"mostOptions" gorm:"column:most_options;comment:最多可选项;"`
}
