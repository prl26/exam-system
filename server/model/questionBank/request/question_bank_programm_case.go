package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank"
)

type QuestionBankProgrammCaseSearch struct {
	questionBank.ProgrammCase
	request.PageInfo
}
