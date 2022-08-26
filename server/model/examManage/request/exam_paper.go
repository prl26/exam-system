package request

import (
	"exam-system/model/common/request"
	"exam-system/model/examManage"
)

type ExamPaperSearch struct {
	examManage.ExamPaper
	request.PageInfo
}
