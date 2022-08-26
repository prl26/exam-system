// 自动生成模板QuestionBankProgramm
package questionBank

import (
	"exam-system/global"
)

// Programm 结构体
type Programm struct {
	global.GVA_MODEL
	BasicModel
}

// TableName Programm 表名
func (Programm) TableName() string {
	return "les_questionBank_programm"
}
