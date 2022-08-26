package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type TeachClassSearch struct {
	basicdata.TeachClass
	request.PageInfo
}
