package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Test"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ClassSearch struct{
    Test.Class
    request.PageInfo
}
