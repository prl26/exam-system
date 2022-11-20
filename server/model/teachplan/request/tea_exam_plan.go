package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type ExamPlanSearch struct {
	teachplan.ExamPlan
	request.PageInfo
}
type ExamPlan struct {
	PlanId uint `json:"planId" form:"planId"`
}
