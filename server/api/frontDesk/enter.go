package frontDesk

import (
	"github.com/prl26/exam-system/server/api/frontDesk/basic"
	"github.com/prl26/exam-system/server/api/frontDesk/exam"
	"github.com/prl26/exam-system/server/api/frontDesk/questionBank"
	"github.com/prl26/exam-system/server/api/frontDesk/system"
	"github.com/prl26/exam-system/server/api/frontDesk/teachplan"
)

type FrontDesk struct {
	BasicApiGroup     basic.ApiGroup
	ExamApiGroup      exam.ApiGroup
	QuestionBankGroup questionBank.ApiGroup
	SystemApiGroup    system.ApiGroup
	TeachplanApiGroup teachplan.ApiGroup
}
