package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
)

type ExamPaperTemplateSearch struct {
	examManage.ExamPaperTemplate
	request.PageInfo
}
