// 自动生成模板QuestionBankJudge
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// QuestionBankJudge 结构体
type Judge struct {
	global.GVA_MODEL
	IsRight *int `json:"isRight" form:"isRight" gorm:"column:is_right;comment:是否正确;"`
	BasicModel
}

// TableName QuestionBankJudge 表名
func (Judge) TableName() string {
	return "les_questionBank_judge"
}
