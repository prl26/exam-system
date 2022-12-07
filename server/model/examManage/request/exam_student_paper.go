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
	TeachClassId *uint `json:"teachClassId" form:"teachClassId" gorm:"column:teach_class_id;comment:教学班id;size:32;"`
	TermId       *uint `json:"termId" form:"termId" gorm:"column:term_id;comment:学期id"`
}
