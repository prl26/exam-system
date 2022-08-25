package request

import (
	"exam-system/model/common/request"
	"exam-system/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
