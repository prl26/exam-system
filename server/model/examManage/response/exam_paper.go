package response

import (
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
)

type ExamPaperResponse struct {
	PaperId               uint               `json:"paperId"`
	SingleChoiceComponent []ChoiceComponent  `json:"singleChoiceComponent"`
	MultiChoiceComponent  []ChoiceComponent  `json:"multiChoiceComponent"`
	JudgeComponent        []JudgeComponent   `json:"judgeComponent"`
	BlankComponent        []BlankComponent   `json:"blankComponent"`
	ProgramComponent      []ProgramComponent `json:"programComponent"`
}
type ChoiceComponent struct {
	MergeId uint                                  `json:"mergeId"`
	Choice  questionBankVoResp.MultipleChoiceExam `json:"choiceComponent"`
}
type JudgeComponent struct {
	MergeId uint                             `json:"mergeId"`
	Judge   questionBankVoResp.JudgePractice `json:"judgeComponent"`
}
type BlankComponent struct {
	MergeId uint                                   `json:"mergeId"`
	Blank   questionBankVoResp.SupplyBlankPractice `json:"blankComponent"`
}
type ProgramComponent struct {
	MergeId uint                               `json:"mergeId"`
	Program questionBankVoResp.ProgramPractice `json:"programComponent"`
}
