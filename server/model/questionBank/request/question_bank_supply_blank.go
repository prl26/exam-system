package request

import (
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
)

type QuestionBankSupplyBlankSearch struct {
	questionBank.SupplyBlank
	request.PageInfo
}
