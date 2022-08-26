package v1

import (
	"exam-system/api/v1/basicdata"
	"exam-system/api/v1/examManage"
	"exam-system/api/v1/lessondata"
	"exam-system/api/v1/questionBank"
	"exam-system/api/v1/system"
	"exam-system/api/v1/teachplan"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	BasicdataApiGroup    basicdata.ApiGroup
	CoursedataApiGroup   lessondata.ApiGroup
	LessondataApiGroup   lessondata.ApiGroup
	TeachplanApiGroup    teachplan.ApiGroup
	ExammanageApiGroup   examManage.ApiGroup
	QuestionBankApiGroup questionBank.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
