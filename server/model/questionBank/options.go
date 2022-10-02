// 自动生成模板QuestionBankOptions
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// Options 结构体
type Options struct {
	global.GVA_MODEL
	Describe         string `json:"describe" form:"describe" gorm:"column:describe;comment:描述;"`
	Orders           uint   `json:"orders" form:"Order" gorm:"orders"`
	MultipleChoiceId uint   `json:"multipleChoiceId" form:"multipleChoiceId" gorm:"column:multiple_choice_id;comment:选择题id;"`
}

// TableName Options 表名
func (Options) TableName() string {
	return "les_questionBank_options"
}
