package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Test"
	"github.com/flipped-aurora/gin-vue-admin/server/service/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/lesson"
	"github.com/flipped-aurora/gin-vue-admin/server/service/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
<<<<<<< HEAD
	"github.com/flipped-aurora/gin-vue-admin/server/service/teachplan"
=======
>>>>>>> 32710530ea6d8ea2188de74ad06a8e5a55421c2e
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	TestServiceGroup       Test.ServiceGroup
	LessonServiceGroup     lesson.ServiceGroup
<<<<<<< HEAD
	BasicdataApiGroup      basicdata.ServiceGroup
	LessondataServiceGroup lessondata.ServiceGroup
	TeachplanServiceGroup  teachplan.ServiceGroup
=======
	BasicdataApiGroup  basicdata.ServiceGroup
	LessondataServiceGroup lessondata.ServiceGroup
>>>>>>> 32710530ea6d8ea2188de74ad06a8e5a55421c2e
}

var ServiceGroupApp = new(ServiceGroup)
