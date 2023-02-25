package response

import "github.com/prl26/exam-system/server/model/teachplan"

type GenerateQRCode struct {
	QRCodeURL  string
	ExpireTime string
	Minute     uint
}

type AttendanceDetail struct {
	teachplan.TeachAttendanceRecord
	Name string `json:"Name" `
}
