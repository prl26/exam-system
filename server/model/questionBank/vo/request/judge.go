package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type QuestionBankJudgeSearch struct {
	questionBankBo.JudgeSearchCriteria
	request.PageInfo
}

type JudgeCreate struct {
	questionBankPo.Judge
}

type JudgeUpdate struct {
	Id uint
	questionBankPo.JudgeModel
}

type JudgePracticeSearch struct {
	questionBankBo.JudgePracticeCriteria
	request.PageInfo
}
