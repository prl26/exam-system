package request

import (
	"exam-system/model/common/request"
	"exam-system/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
