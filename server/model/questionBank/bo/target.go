package bo

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type TargetSearchCriteria struct {
	questionBankPo.SimpleModel
	questionBankPo.CourseSupport
}

type TargetDetail struct {
	global.GVA_MODEL
	questionBankPo.CourseSupport
	questionBankPo.TargetModel
	Chapter   *basicdata.Chapter
	Knowledge *basicdata.Knowledge
}
