package request

type RecordRq struct {
	Student uint `json:"studentId" form:"studentId"`
	PlanId  uint `json:"planId" form:"planId"`
}
type RecordRq1 struct {
	Student  uint `json:"studentId" form:"studentId"`
	PlanId   uint `json:"planId" form:"planId"`
	RecordId uint `json:"recordId" form:"recordId"`
}
