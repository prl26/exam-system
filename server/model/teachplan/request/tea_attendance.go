package request

import (
	"exam-system/model/common/request"
	"exam-system/model/teachplan"
)

type TeachAttendanceSearch struct {
	teachplan.TeachAttendance
	request.PageInfo
}
