// 自动生成模板QuestionBankOptions
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Options 结构体
type Options struct {
	global.GVA_MODEL
	Describe         string `json:"describe" form:"describe" gorm:"column:describe;comment:描述;"`
	MultipleChoiceId *int   `json:"multipleChoiceId" form:"multipleChoiceId" gorm:"column:multiple_choice_id;comment:选择题id;"`
	Orders           *int   `json:"orders" form:"Order" gorm:"orders"`
	QuestionType     *int   `json:"questionType" form:"paperId" gorm:"column:question_type;comment:题目类型;size:8;"`
}

// TableName Options 表名
func (Options) TableName() string {
	return "les_questionBank_options"
}
