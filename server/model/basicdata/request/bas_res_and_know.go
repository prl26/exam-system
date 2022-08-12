package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ResandknowSearch struct{
    basicdata.Resandknow
    request.PageInfo
}
