package request

import (
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
)

type ResourceSearch struct {
	basicdata.Resource
	request.PageInfo
}
