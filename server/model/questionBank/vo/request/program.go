package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type ProgramCreate struct {
	questionBankPo.CourseSupport
	questionBankPo.BasicModel
	questionBankBo.ProgramOjSupport
}

type ProgramUpdate struct {
	Id uint `json:"id"`
	questionBankPo.BasicModel
	questionBankPo.CourseSupport
	questionBankBo.ProgramOjSupport
}
type ProgramSearch struct {
	questionBankBo.ProgramSearchCriteria
	request.PageInfo
}
