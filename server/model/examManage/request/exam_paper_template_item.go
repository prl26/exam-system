package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
)

type PaperTemplateItemSearch struct {
	examManage.PaperTemplateItem
	request.PageInfo
}
