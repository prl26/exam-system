package response

import (
	"github.com/prl26/exam-system/server/global"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type SupplyBlankSimple struct {
	global.GVA_MODEL
	questionBankPo.SimpleModel
}

type SupplyBlankDetail struct {
	global.GVA_MODEL
	questionBankBo.CourseSupportPtr
	questionBankPo.SupplyBlankModel
	Answers questionBankBo.SupplyBlankAnswers `json:"answers"`
}
type SupplyBlankPractice struct {
	questionBankBo.PracticeModel
	Num int `json:"num" form:"num" gorm:"column:num;comment:可填项;"`
}
