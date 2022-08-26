// 自动生成模板TeachAttendance
package teachplan

import (
	"exam-system/global"
	"time"
)

// TeachAttendance 结构体
type TeachAttendance struct {
	global.GVA_MODEL
	Longitude  *float64   `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经度;size:8;"`
	Latitude   *float64   `json:"latitude" form:"latitude" gorm:"column:latitude;comment:维度;size:8;"`
	Expiration *time.Time `json:"expiration" form:"expiration" gorm:"column:expiration;comment:过期时间;"`
	TeachId    *int       `json:"teachId" form:"teachId" gorm:"column:teach_id;comment:教学班id;size:32;"`
}

// TableName TeachAttendance 表名
func (TeachAttendance) TableName() string {
	return "tea_attendance"
}
