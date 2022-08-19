package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Test"
	"github.com/flipped-aurora/gin-vue-admin/server/service/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/lesson"
	"github.com/flipped-aurora/gin-vue-admin/server/service/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/teachplan"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	TestServiceGroup       Test.ServiceGroup
	LessonServiceGroup     lesson.ServiceGroup
	BasicdataApiGroup      basicdata.ServiceGroup
	LessondataServiceGroup lessondata.ServiceGroup
	TeachplanServiceGroup  teachplan.ServiceGroup

}

var ServiceGroupApp = new(ServiceGroup)
