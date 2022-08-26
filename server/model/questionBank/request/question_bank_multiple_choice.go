package request

import (
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
)

type QuestionBankMultipleChoiceSearch struct {
	questionBank.MultipleChoice
	request.PageInfo
}
