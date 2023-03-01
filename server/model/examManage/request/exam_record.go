package request

type RecordRq struct {
	Student uint `json:"studentId" form:"studentId"`
	PlanId  uint `json:"planId" form:"planId"`
}
