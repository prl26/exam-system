package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type TermSearch struct {
	basicdata.Term
	request.PageInfo
}
