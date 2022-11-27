package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/po"
)

type MultipleChoiceList struct {
	questionBankBo.MultipleCriteria
	request.PageInfo
}

type MultipleChoiceCreate struct {
	po.MultipleChoice
}

type MultipleChoiceUpdate struct {
	po.MultipleChoice
}
type MultipleChoicePracticeList struct {
	questionBankBo.MultiplePracticeCriteria
	request.PageInfo
}
