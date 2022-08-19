// 自动生成模板TeachClassStudent
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// TeachClassStudent 结构体
type TeachClassStudent struct {
      global.GVA_MODEL
      Student_id  *int `json:"student_id" form:"student_id" gorm:"column:student_id;comment:;"`
      Teach_class_id  *int `json:"teach_class_id" form:"teach_class_id" gorm:"column:teach_class_id;comment:班级id;"`
}


// TableName TeachClassStudent 表名
func (TeachClassStudent) TableName() string {
  return "bas_teach_class_student"
}

