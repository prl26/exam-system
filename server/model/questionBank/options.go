// 自动生成模板QuestionBankOptions
package questionBank

import (
	"exam-system/global"
)

// Options 结构体
type Options struct {
	global.GVA_MODEL
	Describe         string `json:"describe" form:"describe" gorm:"column:describe;comment:描述;"`
	MultipleChoiceId *int   `json:"multipleChoiceId" form:"multipleChoiceId" gorm:"column:multiple_choice_id;comment:选择题id;"`
	Orders           *int   `json:"orders" form:"Order" gorm:"orders"`
}

// TableName Options 表名
func (Options) TableName() string {
	return "les_questionBank_options"
}
