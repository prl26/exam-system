package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type CollegeSearch struct {
	basicdata.College
	request.PageInfo
}
