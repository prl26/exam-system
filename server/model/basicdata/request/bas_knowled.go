package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type KnowledgeSearch struct{
    basicdata.Knowledge
    request.PageInfo
}
