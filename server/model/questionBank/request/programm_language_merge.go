package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank"
)

type ProgrammLanguageMergeSearch struct {
	questionBank.ProgrammLanguageMerge
	request.PageInfo
}
