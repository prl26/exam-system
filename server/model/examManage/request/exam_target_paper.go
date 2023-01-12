package request

type CommitTargetExamPaper struct {
	StudentId       uint              `json:"studentId" form:"studentId"`
	PlanId          uint              `json:"planId" form:"planId"`
	PaperId         uint              `json:"paperId" form:"paperId"`
	TargetComponent []TargetComponent `json:"targetComponent"`
}
type TargetComponent struct {
	MergeId    uint   `json:"mergeId"`
	QuestionId uint   `json:"questionId" form:"questionId"`
	Answer     string `json:"answer" form:"answer"`
}
