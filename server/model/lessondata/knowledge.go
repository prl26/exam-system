// 自动生成模板Knowledge
package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Knowledge 结构体
type Knowledge struct {
      global.GVA_MODEL
      Chapter_id  *int `json:"chapter_id" form:"chapter_id" gorm:"column:chapter_id;comment:章节id;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:知识点名称;"`
}


// TableName Knowledge 表名
func (Knowledge) TableName() string {
  return "les_knowledge"
}

