package Backstage

import (
	"github.com/prl26/exam-system/server/api/Backstage/basicdata"
	"github.com/prl26/exam-system/server/api/Backstage/examManage"
	"github.com/prl26/exam-system/server/api/Backstage/lessondata"
	"github.com/prl26/exam-system/server/api/Backstage/oj"
	"github.com/prl26/exam-system/server/api/Backstage/questionBank"
	"github.com/prl26/exam-system/server/api/Backstage/system"
	"github.com/prl26/exam-system/server/api/Backstage/teachplan"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	BasicdataApiGroup    basicdata.ApiGroup
	CoursedataApiGroup   lessondata.ApiGroup
	LessondataApiGroup   lessondata.ApiGroup
	TeachplanApiGroup    teachplan.ApiGroup
	ExammanageApiGroup   examManage.ApiGroup
	QuestionBankApiGroup questionBank.ApiGroup
	OjApiGroup           oj.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
