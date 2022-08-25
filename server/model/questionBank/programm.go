// 自动生成模板QuestionBankProgramm
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Programm 结构体
type Programm struct {
	global.GVA_MODEL
	Describe string `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:;"`
}

// TableName Programm 表名
func (Programm) TableName() string {
	return "les_questionBank_programm"
}
