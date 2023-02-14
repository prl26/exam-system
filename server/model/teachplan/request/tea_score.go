package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type ScoreSearch struct {
	teachplan.Score
	request.PageInfo
}
type Excel struct {
	FileName string `json:"fileName"`
	PlanId   uint   `json:"planId"`
}
type MultiExcel struct {
	FileName    string `json:"fileName"`
	TeachPlanId uint   `json:"teachPlanId"`
}
