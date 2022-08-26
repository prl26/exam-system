package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type ClassSearch struct {
	basicdata.Class
	request.PageInfo
}
