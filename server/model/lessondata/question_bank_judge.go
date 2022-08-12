// 自动生成模板QuestionBankJudge
package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// QuestionBankJudge 结构体
type QuestionBankJudge struct {
      global.GVA_MODEL
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:描述文本;"`
      Is_right  *int `json:"is_right" form:"is_right" gorm:"column:is_right;comment:是否正确;"`
}


// TableName QuestionBankJudge 表名
func (QuestionBankJudge) TableName() string {
  return "les_questionBank_judge"
}

