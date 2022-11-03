package response

import "github.com/prl26/exam-system/server/model/questionBank"

type ExamPaperResponse struct {
	PaperId          uint `json:"paperId"`
	ChoiceComponent  []ChoiceComponent
	JudgeComponent   []JudgeComponent   `json:"judgeComponent"`
	BlankComponent   []BlankComponent   `json:"blankComponent"`
	ProgramComponent []ProgramComponent `json:"programComponent"`
}
type ChoiceComponent struct {
	MergeId uint                          `json:"mergeId"`
	Choice  []questionBank.MultipleChoice `json:"questionComponent"`
}
type JudgeComponent struct {
	MergeId uint                       `json:"mergeId"`
	Judge   []questionBank.SupplyBlank `json:"blankComponent"`
}
type BlankComponent struct {
	MergeId uint                       `json:"mergeId"`
	Blank   []questionBank.SupplyBlank `json:"blankComponent"`
}
type ProgramComponent struct {
	MergeId uint                    `json:"mergeId"`
	Program []questionBank.Programm `json:"programComponent"`
}
