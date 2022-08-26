package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
)

type PaperQuestionMergeSearch struct {
	examManage.PaperQuestionMerge
	request.PageInfo
}
