package questionBank

import "github.com/prl26/exam-system/server/service/questionBank/oj"

type ServiceGroup struct {
	PublicProgramService
	ProgramService
	JudgeService
	SupplyBlankService
	MultipleChoiceService
	QuestionBankService
	oj.OjService
}
