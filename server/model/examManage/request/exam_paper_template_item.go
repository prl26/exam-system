package request

import (
	"exam-system/model/common/request"
	"exam-system/model/examManage"
)

type PaperTemplateItemSearch struct {
	examManage.PaperTemplateItem
	request.PageInfo
}
