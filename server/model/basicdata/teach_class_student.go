// 自动生成模板TeachClassStudent
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeachClassStudent 结构体
type TeachClassStudent struct {
	global.GVA_MODEL
	StudentId    *int `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;foreignKey:studentid"`
	TeachClassId *int `json:"teachClassId" form:"teachClassId" gorm:"column:teach_class_id;comment:班级id;"`
}

// TableName TeachClassStudent 表名
func (TeachClassStudent) TableName() string {
	return "teach_class_student"
}
