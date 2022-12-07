package response

import (
	"github.com/prl26/exam-system/server/global"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type MultipleChoiceSimple struct {
	global.GVA_MODEL
	questionBank.SimpleModel
}

type MultipleChoicePractice struct {
	questionBankBo.PracticeModel
	MostOptions int `json:"mostOptions" form:"mostOptions" gorm:"column:most_options;comment:最多可选项;"`
}
type MultipleChoiceExam struct {
	questionBankBo.PracticeModel
	MostOptions  int `json:"mostOptions" form:"mostOptions" gorm:"column:most_options;comment:最多可选项;"`
	IsIndefinite int `json:"isIndefinite" form:"isIndefinite" gorm:"column:is_indefinite"`
}
