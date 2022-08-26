package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type ResourceSearch struct {
	basicdata.Resource
	request.PageInfo
}
