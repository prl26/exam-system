package request

import (
	"exam-system/model/common/request"
	"exam-system/model/teachplan"
)

type TeachAttendanceRecordSearch struct {
	teachplan.TeachAttendanceRecord
	request.PageInfo
}
