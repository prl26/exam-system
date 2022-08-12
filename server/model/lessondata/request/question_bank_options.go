package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type QuestionBankOptionsSearch struct{
    lessondata.QuestionBankOptions
    request.PageInfo
}
