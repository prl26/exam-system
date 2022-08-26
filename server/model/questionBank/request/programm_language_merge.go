package request

import (
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
)

type ProgrammLanguageMergeSearch struct {
	questionBank.ProgrammLanguageMerge
	request.PageInfo
}
