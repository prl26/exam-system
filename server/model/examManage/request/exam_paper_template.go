package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
)

type ExamPaperTemplateSearch struct {
	examManage.ExamPaperTemplate
	request.PageInfo
}
