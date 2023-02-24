// 自动生成模板TeachAttendance
package teachplan

import (
	"github.com/prl26/exam-system/server/global"
)

// TeachAttendance 结构体
type TeachAttendance struct {
	global.GVA_MODEL
	TeachClassId uint `json:"teachClassId" form:"teachClassId" gorm:"column:teach_class_id;comment:教学班id;size:32;"`
}

// TableName TeachAttendance 表名
func (TeachAttendance) TableName() string {
	return "tea_attendance"
}
