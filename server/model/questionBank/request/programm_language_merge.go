package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
)

type ProgrammLanguageMergeSearch struct {
	questionBank.ProgrammLanguageMerge
	request.PageInfo
}
