package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type StudentSearch struct {
	basicdata.Student
	request.PageInfo
}
