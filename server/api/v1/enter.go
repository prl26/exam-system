package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Test"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
<<<<<<< HEAD
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/teachplan"
=======
>>>>>>> 32710530ea6d8ea2188de74ad06a8e5a55421c2e
)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	TestApiGroup       Test.ApiGroup
	BasicdataApiGroup  basicdata.ApiGroup
	CoursedataApiGroup lessondata.ApiGroup
	LessondataApiGroup lessondata.ApiGroup
<<<<<<< HEAD
	TeachplanApiGroup  teachplan.ApiGroup
=======
>>>>>>> 32710530ea6d8ea2188de74ad06a8e5a55421c2e
}

var ApiGroupApp = new(ApiGroup)
