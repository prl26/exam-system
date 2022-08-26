package request

import (
	"exam-system/model/common/request"
	"exam-system/model/questionBank"
)

type QuestionBankJudgeSearch struct {
	questionBank.Judge
	request.PageInfo
}
