package request

import (
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
)

type QuestionBankProgrammCaseSearch struct {
	questionBank.ProgrammCase
	request.PageInfo
}
