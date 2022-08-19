package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type QuestionBankSupplyBlankSearch struct{
    lessondata.QuestionBankSupplyBlank
    request.PageInfo
}
