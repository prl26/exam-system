package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type TeachAttendanceSearch struct {
	teachplan.TeachAttendance
	request.PageInfo
}

type GenerateQrCode struct {
	AttendanceId uint `json:"attendanceId"`
	Minute       uint `json:"minute"`
}

type Attendance struct {
	Code      string `json:"code"`
	StudentId uint   `json:"studentId"`
}

type AttendanceDetail struct {
	AttendanceId uint `json:"attendanceId" form:"attendanceId"`
	request.PageInfo
}

type Supplement struct {
	AttendanceId uint `json:"attendanceId"`
	StudentId    uint `json:"studentId"`
}
