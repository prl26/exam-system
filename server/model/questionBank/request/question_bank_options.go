package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank"
)

type QuestionBankOptionsSearch struct {
	questionBank.Options
	request.PageInfo
}
