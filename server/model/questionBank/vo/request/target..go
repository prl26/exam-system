package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type TargetCreate struct {
	questionBankPo.Target
}

type TargetSearch struct {
	questionBankBo.TargetSearchCriteria
	request.PageInfo
}
