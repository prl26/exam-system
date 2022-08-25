package request

import (
	"exam-system/model/common/request"
	"exam-system/model/teachplan"
)

type ExamPlanSearch struct {
	teachplan.ExamPlan
	request.PageInfo
}
