package request

import (
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
)

type CollegeSearch struct {
	basicdata.College
	request.PageInfo
}
