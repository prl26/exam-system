package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
)

type QuestionBankProgrammSearch struct {
	questionBank.Programm
	request.PageInfo
}
