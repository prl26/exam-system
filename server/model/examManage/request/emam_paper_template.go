package request

import (
	"exam-system/model/common/request"
	"exam-system/model/examManage"
)

type PaperTemplateSearch struct {
	examManage.PaperTemplate
	request.PageInfo
}
