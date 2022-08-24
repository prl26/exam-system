package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/examManage"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/lessondata"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/questionBank"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/teachplan"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	ExampleApiGroup      example.ApiGroup
	BasicdataApiGroup    basicdata.ApiGroup
	CoursedataApiGroup   lessondata.ApiGroup
	LessondataApiGroup   lessondata.ApiGroup
	TeachplanApiGroup    teachplan.ApiGroup
	ExammanageApiGroup   examManage.ApiGroup
	QuestionBankApiGroup questionBank.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
