package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PaperTemplateSearch struct{
    examManage.PaperTemplate
    request.PageInfo
}
