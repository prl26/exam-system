package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank"
)

type QuestionBankMultipleChoiceSearch struct {
	questionBank.MultipleChoice
	request.PageInfo
}
