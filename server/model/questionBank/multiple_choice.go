// 自动生成模板QuestionBankMultipleChoice
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MultipleChoice 结构体
type MultipleChoice struct {
	global.GVA_MODEL
	Title       string `json:"title" form:"title" gorm:"column:title;comment:标题;`
	Describe    string `json:"describe" form:"describe" gorm:"column:describe;comment:选择题描述;"`
	MostOptions *int   `json:"mostOptions" form:"mostOptions" gorm:"column:most_options;comment:最多可选项;"`
	Answer      string `json:"answer" form:"answer" gorm:"column:answer;comment:"`
}

// TableName MultipleChoice 表名
func (MultipleChoice) TableName() string {
	return "les_questionBank_multiple_choice"
}
