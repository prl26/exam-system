package request

import (
	"exam-system/model/common/request"
	"exam-system/model/examManage"
)

type PaperQuestionMergeSearch struct {
	examManage.PaperQuestionMerge
	request.PageInfo
}
