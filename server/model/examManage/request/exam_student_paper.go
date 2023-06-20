package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
)

type ExamStudentPaperSearch struct {
	examManage.ExamStudentPaper
	//request.PageInfo
}
type ExamStudentScore struct {
	ScoreSearch
	request.PageInfo
}

type ScoreSearch struct {
	LessonId *int `json:"lessonId" form:"lessonId" gorm:"column:lesson_id;comment:课程id;size:32;"`
	TermId   *int `json:"termId" form:"termId" gorm:"column:term_id;comment:学期id"`
}
type PaperMultiReview struct {
	TeachPlanId uint `json:"teachPlanId"`
	request.PageInfo
}
type PaperReview struct {
	examManage.ExamScore
	TeachPlanId int `json:"teachPlanId"`
	request.PageInfo
}
type PaperCheating struct {
	StudentId      uint             `json:"studentId"`
	PlanId         uint             `json:"planId"`
	AnswerCheating []AnswerCheating `json:"answerCheating"`
}
type AnswerCheating struct {
	MergeId  uint    `json:"mergeId"`
	Answer   string  `json:"answer"`
	GotScore float64 `json:"gotScore"`
}
type StatusMonitor struct {
	examManage.StudentPaperStatus
	//request.PageInfo
}

type UploadExamPicture struct {
	PlanId uint
}
