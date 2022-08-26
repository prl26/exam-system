package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
)

type ExamStudentPaperSearch struct {
	examManage.ExamStudentPaper
	request.PageInfo
}
