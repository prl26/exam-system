package response

import "github.com/prl26/exam-system/server/model/questionBank"

type ExamPaperResponse struct {
	PaperId          uint                          `json:"paperId"`
	ChoiceComponent  []questionBank.MultipleChoice `json:"questionComponent"`
	JudgeComponent   []questionBank.Judge          `json:"judgeComponent"`
	BlankComponent   []questionBank.SupplyBlank    `json:"blankComponent"`
	ProgramComponent []questionBank.Programm       `json:"programComponent"`
}
