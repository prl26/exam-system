package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/router/coursedata"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Basicdata  basicdata.RouterGroup
	Coursedata coursedata.RouterGroup
	Lessondata lessondata.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
