package request

import (
	"exam-system/model/common/request"
	"exam-system/model/examManage"
)

type ExamPaperTemplateSearch struct {
	examManage.ExamPaperTemplate
	request.PageInfo
}
