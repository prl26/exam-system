package examManage

import "github.com/prl26/exam-system/server/global"

type StudentPaperStatus struct {
	global.GVA_MODEL
	StudentId uint `json:"studentId,omitempty"`
	PlanId    uint `json:"planId"`
}

// TableName ExamStudentPaper 表名
func (StudentPaperStatus) TableName() string {
	return "student_paper_status"
}
