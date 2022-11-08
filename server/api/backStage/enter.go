package backStage

import (
	"github.com/prl26/exam-system/server/api/backStage/basicdata"
	"github.com/prl26/exam-system/server/api/backStage/examManage"
	"github.com/prl26/exam-system/server/api/backStage/lessondata"
	"github.com/prl26/exam-system/server/api/backStage/questionBank"
	"github.com/prl26/exam-system/server/api/backStage/system"
	"github.com/prl26/exam-system/server/api/backStage/teachplan"
)

type BackStage struct {
	BasicDataApiGroup    basicdata.ApiGroup
	ExamManageApiGroup   examManage.ApiGroup
	LessondataApiGroup   lessondata.ApiGroup
	QuestionBankApiGroup questionBank.ApiGroup
	SystemApiGroup       system.ApiGroup
	TeachPlanApiGroup    teachplan.ApiGroup
}
