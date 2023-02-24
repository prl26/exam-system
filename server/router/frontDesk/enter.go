package frontDesk

import (
	"github.com/prl26/exam-system/server/router/frontDesk/basic"
	"github.com/prl26/exam-system/server/router/frontDesk/exam"
	"github.com/prl26/exam-system/server/router/frontDesk/questionBank"
	"github.com/prl26/exam-system/server/router/frontDesk/system"
	"github.com/prl26/exam-system/server/router/frontDesk/teachplan"
)

type RouterGroup struct {
	BasicRouterGroup        basic.RouterGroup
	SystemRouterGroup       system.RouterGroup
	ExamRouterGroup         exam.RouterGroup
	QuestionBankRouterGroup questionBank.RouterGroup
	TeachPlanRouterGroup    teachplan.RouterGroup
}
