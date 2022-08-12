// 自动生成模板QuestionBankSupplyBlank
package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// QuestionBankSupplyBlank 结构体
type QuestionBankSupplyBlank struct {
      global.GVA_MODEL
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:填空题描述;"`
      Is_order  *int `json:"is_order" form:"is_order" gorm:"column:is_order;comment:是否要求有序;"`
      Num  *int `json:"num" form:"num" gorm:"column:num;comment:可填项;"`
}


// TableName QuestionBankSupplyBlank 表名
func (QuestionBankSupplyBlank) TableName() string {
  return "les_questionBank_supply_blank"
}

