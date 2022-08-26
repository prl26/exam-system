package request

import (
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
)

type QuestionBankProgrammSearch struct {
	questionBank.Programm
	request.PageInfo
}
