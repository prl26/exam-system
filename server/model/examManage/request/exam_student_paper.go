package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
)

type ExamStudentPaperSearch struct {
	examManage.ExamStudentPaper
	request.PageInfo
}
type ExamStudentScore struct {
	ScoreSearch
	request.PageInfo
}
type ScoreSearch struct {
	LessonId *int `json:"lessonId" form:"lessonId" gorm:"column:lesson_id;comment:课程id;size:32;"`
	TermId   *int `json:"termId" form:"termId" gorm:"column:term_id;comment:学期id"`
}
