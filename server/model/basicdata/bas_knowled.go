// 自动生成模板Knowledge
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Knowledge 结构体
type Knowledge struct {
      global.GVA_MODEL
      ChapterId  string `json:"chapterId" form:"chapterId" gorm:"column:chapter_id;comment:改知识点属于哪个章节;size:32;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:知识点名称(简述);size:255;"`
}


// TableName Knowledge 表名
func (Knowledge) TableName() string {
  return "knowledge"
}

