package request

import (
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
)

type StudentSearch struct {
	basicdata.Student
	request.PageInfo
}
