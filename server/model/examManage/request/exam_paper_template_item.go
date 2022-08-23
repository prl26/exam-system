package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PaperTemplateItemSearch struct{
    examManage.PaperTemplateItem
    request.PageInfo
}
