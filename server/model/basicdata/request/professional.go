package request

import (
	"exam-system/model/basicdata"
	"exam-system/model/common/request"
)

type ProfessionalSearch struct {
	basicdata.Professional
	request.PageInfo
}
