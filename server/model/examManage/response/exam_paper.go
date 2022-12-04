package response

import (
	"github.com/prl26/exam-system/server/model/questionBank/po"
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
	MergeId uint                                      `json:"mergeId"`
	Choice  questionBankVoResp.MultipleChoicePractice `json:"question"`
}
type JudgeComponent struct {
	MergeId uint                             `json:"mergeId"`
	Judge   questionBankVoResp.JudgePractice `json:"Judge"`
}
type BlankComponent struct {
	MergeId uint                                   `json:"mergeId"`
	Blank   questionBankVoResp.SupplyBlankPractice `json:"blank"`
}
type ProgramComponent struct {
	MergeId uint        `json:"mergeId"`
	Program ProgramExam `json:"program"`
}
type ProgramExam struct {
	po.BasicModel
	LanguageSupports      string `json:"languageSupport"`
	LanguageSupportsBrief string `json:"languageSupportBrief"`
}
