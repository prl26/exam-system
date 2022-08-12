package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type QuestionBankJudgeSearch struct{
    lessondata.QuestionBankJudge
    request.PageInfo
}
