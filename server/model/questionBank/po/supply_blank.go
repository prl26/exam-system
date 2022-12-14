// 自动生成模板QuestionBankSupplyBlank
package po

import (
	"github.com/prl26/exam-system/server/global"
)

// QuestionBankSupplyBlank 结构体
type SupplyBlank struct {
	global.GVA_MODEL
	SupplyBlankModel
	CourseSupport
	Answer     string `json:"answer" form:"answer" gorm:"column:answer;comment:答案"`
	Proportion string `json:"proportion"`
}

type SupplyBlankModel struct {
	BasicModel
	IsOrder *bool `json:"isOrder" form:"isOrder" gorm:"column:is_order;comment:是否要求有序;"`
	Num     *int  `json:"num" form:"num" gorm:"column:num;comment:可填项;"`
}

// TableName QuestionBankSupplyBlank 表名
func (SupplyBlank) TableName() string {
	return "les_questionBank_supply_blank"
}
