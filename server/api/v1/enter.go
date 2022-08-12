package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Test"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	TestApiGroup       Test.ApiGroup
	BasicdataApiGroup  basicdata.ApiGroup
	CoursedataApiGroup lessondata.ApiGroup
	LessondataApiGroup lessondata.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
