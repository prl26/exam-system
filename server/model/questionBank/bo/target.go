package bo

import (
	"github.com/prl26/exam-system/server/global"
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
	CourseSupportPtr
}

type TargetPracticeCriteria struct {
	questionBankPo.CourseSupport
}
