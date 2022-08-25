// 自动生成模板QuestionBankJudge
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// QuestionBankJudge 结构体
type Judge struct {
	global.GVA_MODEL
	Describe string `json:"describe" form:"describe" gorm:"column:describe;comment:描述文本;"`
	IsRight  *int   `json:"isRight" form:"isRight" gorm:"column:is_right;comment:是否正确;"`
}

// TableName QuestionBankJudge 表名
func (Judge) TableName() string {
	return "les_questionBank_judge"
}
