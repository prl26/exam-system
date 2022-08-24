package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
)

type QuestionBankProgrammSearch struct {
	lessondata.QuestionBankProgramm
	request.PageInfo
}
