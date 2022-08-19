package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
<<<<<<< HEAD
	"github.com/flipped-aurora/gin-vue-admin/server/router/teachplan"
=======
>>>>>>> 32710530ea6d8ea2188de74ad06a8e5a55421c2e
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Basicdata  basicdata.RouterGroup
	Lessondata lessondata.RouterGroup
<<<<<<< HEAD
	Teachplan  teachplan.RouterGroup
=======
>>>>>>> 32710530ea6d8ea2188de74ad06a8e5a55421c2e
}

var RouterGroupApp = new(RouterGroup)
