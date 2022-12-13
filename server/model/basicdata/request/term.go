package request

import (
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
)

type TermSearch struct {
	basicdata.Term
	request.PageInfo
}
type FrontTermSearch struct {
	basicdata.Term
}
