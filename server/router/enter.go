package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/router/examManage"
	"github.com/flipped-aurora/gin-vue-admin/server/router/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/router/questionBank"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/teachplan"
)

type RouterGroup struct {
	System       system.RouterGroup
	Basicdata    basicdata.RouterGroup
	Lessondata   lessondata.RouterGroup
	Teachplan    teachplan.RouterGroup
	Exammanage   examManage.RouterGroup
	QuestionBank questionBank.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
