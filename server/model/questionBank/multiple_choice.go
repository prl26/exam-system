// 自动生成模板QuestionBankMultipleChoice
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// MultipleChoice 结构体
type MultipleChoice struct {
	global.GVA_MODEL
	BasicModel
	Answer      string    `json:"answer" form:"answer" gorm:"column:answer;comment:"`
	MostOptions *int      `json:"mostOptions" form:"mostOptions" gorm:"column:most_options;comment:最多可选项;"`
	Options     []Options `json:"options"`
}

// TableName MultipleChoice 表名
func (MultipleChoice) TableName() string {
	return "les_questionBank_multiple_choice"
}
