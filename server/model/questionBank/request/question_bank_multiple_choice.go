package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
)

type QuestionBankMultipleChoiceSearch struct {
	questionBank.MultipleChoice
	request.PageInfo
}
