package response

import "github.com/prl26/exam-system/server/global"

type TargetExamPaperResponse struct {
	PaperId         uint              `json:"paperId"`
	TargetComponent []TargetComponent `json:"targetComponent"`
}

type TargetComponent struct {
	MergeId uint   `json:"mergeId"`
	Target  Target `json:"target"`
}
type TargetComponent1 struct {
	MergeId uint   `json:"mergeId"`
	Target  Target `json:"target"`
	Score   *int   `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
}
type Target struct {
	global.GVA_MODEL
	Title    string `json:"title" form:"title" gorm:"column:title;comment:;"`
	Describe string `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
	Code     string `json:"code" form:"code" gorm:"code"`
	ByteCode string `json:"byteCode" form:"byteCode"`
}
