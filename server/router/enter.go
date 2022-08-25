package router

import (
	"exam-system/router/basicdata"
	"exam-system/router/examManage"
	"exam-system/router/lessondata"
	"exam-system/router/questionBank"
	"exam-system/router/system"
	"exam-system/router/teachplan"
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
