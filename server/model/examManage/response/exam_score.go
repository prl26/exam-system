package response

import (
	"github.com/prl26/exam-system/server/global"
	"time"
)

type ExamScoreResponse struct {
	global.GVA_MODEL
	PlanId     *uint      `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:考试名称;size:64;"`
	CourseName string     `json:"courseName" form:"courseName"`
	TermName   string     `json:"termName" form:"termName"`
	Score      *int       `json:"score" form:"score"`
	ExamType   *int       `json:"examType" form:"examType"`
	StartTime  *time.Time `json:"startTime" form:"startTime"`
}
