package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
)

type ExamPaperSearch struct {
	examManage.ExamPaper
	request.PageInfo
}
