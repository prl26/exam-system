package request

type ExamComing struct {
	StudentId uint `json:"studentId" form:"studentId"`
	PlanId    uint `json:"planId" form:"planId"`
}
