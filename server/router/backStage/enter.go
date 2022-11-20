package backStage

import (
	"github.com/prl26/exam-system/server/router/backStage/basicdata"
	"github.com/prl26/exam-system/server/router/backStage/examManage"
	"github.com/prl26/exam-system/server/router/backStage/lessondata"
	"github.com/prl26/exam-system/server/router/backStage/questionBank"
	"github.com/prl26/exam-system/server/router/backStage/system"
	"github.com/prl26/exam-system/server/router/backStage/teachplan"
)

type BackStage struct {
	System       system.RouterGroup
	Basicdata    basicdata.RouterGroup
	Lessondata   lessondata.RouterGroup
	Teachplan    teachplan.RouterGroup
	Exammanage   examManage.RouterGroup
	QuestionBank questionBank.RouterGroup

}
