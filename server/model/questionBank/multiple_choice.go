// 自动生成模板QuestionBankMultipleChoice
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MultipleChoice 结构体
type MultipleChoice struct {
	global.GVA_MODEL
	Describe    string `json:"describe" form:"describe" gorm:"column:describe;comment:选择题描述;"`
	MostOptions *int   `json:"mostOptions" form:"mostOptions" gorm:"column:most_options;comment:最多可选项;"`
	Answer      string `json:"answer" form:"answer" gorm:"column:answer;comment:若选择题类型为简单类型,则该字段有效;"`
	Type        *int   `json:"type" form:"type" gorm:"column:type;comment:考虑是简单的就只有ABCD,还是说再经过一次查询选项;"`
}

// TableName MultipleChoice 表名
func (MultipleChoice) TableName() string {
	return "les_questionBank_multiple_choice"
}
