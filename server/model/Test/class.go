// 自动生成模板Class
package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Class 结构体
type Class struct {
      global.GVA_MODEL
      Id  *int `json:"id" form:"id" gorm:"column:id;comment:;size:255;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`
}


// TableName Class 表名
func (Class) TableName() string {
  return "class"
}

