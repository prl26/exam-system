package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"time"
)

type StudentPaperStatus struct {
	global.GVA_MODEL
	StudentId uint `json:"studentId"`
	PlanId    uint `json:"planId" form:"planId"`
	EnterTime time.Time
	EndTime   time.Time
	IsCommit  bool   `json:"isCommit"` //是否提交
	Ip        string `json:"ip"`       //学生ip
}

// TableName ExamStudentPaper 表名
func (StudentPaperStatus) TableName() string {
	return "student_paper_status"
}
