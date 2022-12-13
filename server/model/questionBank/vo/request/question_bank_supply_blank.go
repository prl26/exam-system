package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type QuestionBankSupplyBlankSearch struct {
	questionBankBo.SupplyBlankSearchCriteria
	request.PageInfo
}

type SupplyBlankCreate struct {
	questionBankPo.SupplyBlankModel
	questionBankPo.CourseSupport
	Answers questionBankBo.SupplyBlankAnswers `json:"answers"`
}

type SupplyBlankUpdate struct {
	Id uint `json:"id"`
	questionBankPo.SupplyBlankModel
	IsOrder int                               `json:"isOrder" form:"isOrder" gorm:"column:is_order;comment:是否要求有序;"`
	Answers questionBankBo.SupplyBlankAnswers `json:"answers"`
}
type QuestionBankSupplyBlankPracticeSearch struct {
	questionBankBo.SupplyBlankPracticeCriteria
	request.PageInfo
}
