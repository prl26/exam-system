package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"time"
)

type ExamScore struct {
	global.GVA_MODEL
	StudentId  *uint      `json:"studentId" form:"studentId"`
	PlanId     *uint      `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	TermId     *uint      `json:"termId" from:"termId"`
	TermName   string     `json:"termName" form:"termName"`
	CourseId   *int       `json:"courseId" form:"courseId"`
	CourseName string     `json:"courseName" form:"courseName"`
	Score      *float64   `json:"score" form:"score"`
	ExamType   *int       `json:"examType" form:"examType"`
	StartTime  *time.Time `json:"startTime" form:"startTime"`
}

func (ExamScore) TableName() string {
	return "exam_scores"
}

type Detail struct {
	TermName   string `json:"termName"`
	CourseName string `json:"courseName"`
}
