package response

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
)

type SupplyBlankSimple struct {
	global.GVA_MODEL
	questionBankPo.SimpleModel
}

type SupplyBlankDetail struct {
	global.GVA_MODEL
	Chapter   basicdata.Chapter
	Knowledge basicdata.Knowledge
	questionBankPo.BasicModel
}
type SupplyBlankPractice struct {
	PracticeModel
	Num int `json:"num" form:"num" gorm:"column:num;comment:可填项;"`
}
