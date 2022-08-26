// 自动生成模板TeachAttendanceRecord
package teachplan

import (
	"exam-system/global"
)

// TeachAttendanceRecord 结构体
type TeachAttendanceRecord struct {
	global.GVA_MODEL
	StudentId    *int     `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;size:32;"`
	Longitute    *float64 `json:"longitute" form:"longitute" gorm:"column:longitute;comment:经度;size:8;"`
	Latitude     *float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:维度;size:8;"`
	AttendanceId *int     `json:"attendanceId" form:"attendanceId" gorm:"column:attendance_id;comment:考勤Id;size:32;"`
}

// TableName TeachAttendanceRecord 表名
func (TeachAttendanceRecord) TableName() string {
	return "tea_attendance_record"
}
