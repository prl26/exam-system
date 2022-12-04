package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"time"
)

type StudentPaperStatus struct {
	global.GVA_MODEL
	StudentId uint `json:"studentId"`
	PlanId    uint `json:"planId"`
	EnterTime time.Time
	IsCommit  bool
}

// TableName ExamStudentPaper 表名
func (StudentPaperStatus) TableName() string {
	return "student_paper_status"
}
