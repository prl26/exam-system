// 自动生成模板PaperTemplate
package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// PaperTemplate 结构体
type PaperTemplate struct {
      global.GVA_MODEL
      CourseId  *int `json:"courseId" form:"courseId" gorm:"column:course_id;comment:课程id;size:32;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:数据模板名称;size:64;"`
      Memo  string `json:"memo" form:"memo" gorm:"column:memo;comment:备注;size:255;"`
      PaperTemplateItems []PaperTemplateItem `json:"paper_template_items" gorm:"many2many:exam_paper_template_item"`
}


// TableName PaperTemplate 表名
func (PaperTemplate) TableName() string {
  return "exam_paper_template"
}

