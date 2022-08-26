package request

import (
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
)

type QuestionBankOptionsSearch struct {
	questionBank.Options
	request.PageInfo
}
