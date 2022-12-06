package bo

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type MultipleCriteria struct {
	IsIndefinite int `form:"isIndefinite"`
	questionBankPo.SimpleModel
	questionBankPo.CourseSupport
}

type MultipleDetail struct {
	global.GVA_MODEL
	questionBankPo.CourseSupport
	questionBankPo.MultipleChoiceModel
	Chapter   *basicdata.Chapter
	Knowledge *basicdata.Knowledge
}

type MultiplePracticeCriteria struct {
	questionBankPo.CourseSupport
}
