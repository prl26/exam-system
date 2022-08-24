// 自动生成模板QuestionBankProgramm
package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// QuestionBankProgramm 结构体
type QuestionBankProgramm struct {
	global.GVA_MODEL
	Describe string `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:;"`
}

// TableName QuestionBankProgramm 表名
func (QuestionBankProgramm) TableName() string {
	return "les_questionBank_programm"
}
