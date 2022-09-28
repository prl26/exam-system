// 自动生成模板QuestionBankProgrammCase
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// QuestionBankProgrammCase 结构体
type ProgrammCase struct {
	global.GVA_MODEL
	ProgrammLimit
	ProgrammId uint   `json:"programmId" form:"programmId" gorm:"column:programm_id;comment:;"`
	LanguageId int    `json:"languageId" form:"languageId" gorm:"column:language_id;comment:;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Score      uint   `json:"score" form:"score" gorm:"column:score;comment:;"`
	Input      string `json:"input" form:"input" gorm:"column:input;comment:;"`
	Output     string `json:"output" form:"output" gorm:"column:output;comment:;"`
}

// TableName QuestionBankProgrammCase 表名
func (ProgrammCase) TableName() string {
	return "les_questionBank_programm_case"
}
