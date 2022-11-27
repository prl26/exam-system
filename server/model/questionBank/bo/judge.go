package bo

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type JudgeSearchCriteria struct {
	questionBankPo.SimpleModel
	questionBankPo.CourseSupport
}

type JudgePracticeCriteria struct {
	questionBankPo.CourseSupport
}

type JudgeDetail struct {
	global.GVA_MODEL
	questionBankPo.CourseSupport
	questionBankPo.JudgeModel
	Chapter   *basicdata.Chapter
	Knowledge *basicdata.Knowledge
}
