package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"time"
)

type ExamScore struct {
	global.GVA_MODEL
	StudentId  *uint      `json:"studentId" form:"studentId"`
	PlanId     *uint      `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:考试计划名称;size:64;"`
	TermId     *uint      `json:"termId" from:"termId"`
	TermName   string     `json:"termName" form:"termName"`
	LessonId   *int       `json:"lessonId" form:"lessonId" gorm:"column:lesson_id;comment:课程Id;size:32;"`
	CourseName string     `json:"courseName" form:"courseName"`
	Score      *float64   `json:"score" form:"score" gorm:"column:score;comment:本题分值;size:8;"`
	ExamType   *int       `json:"examType" form:"examType"`
	StartTime  *time.Time `json:"startTime" form:"startTime"`
	Weight     *int       `json:"weight" form:"weight" gorm:"column:weight;comment:权重;size:8;"`
	IsReport   bool       `json:"isReport" form:"isReport" gorm:"is_report;comment:是否上报"`
}

func (ExamScore) TableName() string {
	return "exam_scores"
}

type AllScore struct {
	Score float64 `json:"score"`
}

type Detail struct {
	TermName   string `json:"termName"`
	CourseName string `json:"courseName"`
}
type ReviewScore struct {
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
	Score     *float64  `json:"score" form:"score"`
	IsReport  bool      `json:"isReport" form:"isReport" gorm:"is_report;comment:是否上报"`
}
type ReviewScore1 struct {
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
	Score     *float64  `json:"score" form:"score"`
	IsReport  bool      `json:"isReport" form:"isReport" gorm:"is_report;comment:是否上报"`
	Message   string    `json:"message"`
}
type ScoreList struct {
	Score    int64 `json:"score" form:"score" gorm:"score"`
	GotScore int64 `json:"gotScore" form:"gotScore" gorm:"got_score"`
}
type TargetExamingScore struct {
	MergeId uint   `json:"mergeId" form:"mergeId"`
	Answer  string `json:"answer" form:"answer"`
	Score   int64  `json:"score" form:"score" gorm:"score"`
}
