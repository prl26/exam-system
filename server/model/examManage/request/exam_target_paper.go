package request

type CommitTargetExamPaper struct {
	StudentId       uint              `json:"studentId" form:"studentId"`
	PlanId          uint              `json:"planId" form:"planId"`
	PaperId         uint              `json:"paperId" form:"paperId"`
	TargetComponent []TargetComponent `json:"targetComponent"`
}
type ReExecTargetExamPaper struct {
	StudentId       uint                    `json:"studentId" form:"studentId"`
	PlanId          uint                    `json:"planId" form:"planId"`
	PaperId         uint                    `json:"paperId" form:"paperId"`
	TargetComponent []ReExecTargetComponent `json:"targetComponent"`
}
type ReExecTargetComponent struct {
	Id         uint   `json:"id"`
	QuestionId uint   `json:"questionId" form:"questionId"`
	Answer     string `json:"answer" form:"answer"`
}
type TargetComponent struct {
	MergeId    uint   `json:"mergeId"`
	QuestionId uint   `json:"questionId" form:"questionId"`
	Answer     string `json:"answer" form:"answer"`
}
type TargetInstance struct {
	Id     uint `json:"id" form:"id"`
	PlanId uint `json:"planId" form:"planId"`
}
