package response

import (
	questionBankVoResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
)

type ExamPaperResponse1 struct {
	PaperId               uint                `json:"paperId"`
	SingleChoiceComponent []ChoiceComponent1  `json:"singleChoiceComponent"`
	MultiChoiceComponent  []ChoiceComponent1  `json:"multiChoiceComponent"`
	JudgeComponent        []JudgeComponent1   `json:"judgeComponent"`
	BlankComponent        []BlankComponent1   `json:"blankComponent"`
	ProgramComponent      []ProgramComponent1 `json:"programComponent"`
	TargetComponent       []TargetComponent1  `json:"targetComponent"`
}
type ExamPaperResponse struct {
	PaperId               uint               `json:"paperId"`
	SingleChoiceComponent []ChoiceComponent  `json:"singleChoiceComponent"`
	MultiChoiceComponent  []ChoiceComponent  `json:"multiChoiceComponent"`
	JudgeComponent        []JudgeComponent   `json:"judgeComponent"`
	BlankComponent        []BlankComponent   `json:"blankComponent"`
	ProgramComponent      []ProgramComponent `json:"programComponent"`
	TargetComponent       []TargetComponent  `json:"targetComponent"`
}
type ChoiceComponent struct {
	MergeId uint                                  `json:"mergeId"`
	Choice  questionBankVoResp.MultipleChoiceExam `json:"choiceComponent"`
}
type ChoiceComponent1 struct {
	MergeId uint                                  `json:"mergeId"`
	Choice  questionBankVoResp.MultipleChoiceExam `json:"choiceComponent"`
	Score   *int                                  `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
}
type JudgeComponent struct {
	MergeId uint                             `json:"mergeId"`
	Judge   questionBankVoResp.JudgePractice `json:"judgeComponent"`
}
type JudgeComponent1 struct {
	MergeId uint                             `json:"mergeId"`
	Judge   questionBankVoResp.JudgePractice `json:"judgeComponent"`
	Score   *int                             `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
}
type BlankComponent struct {
	MergeId uint                                   `json:"mergeId"`
	Blank   questionBankVoResp.SupplyBlankPractice `json:"blankComponent"`
}
type BlankComponent1 struct {
	MergeId uint                                   `json:"mergeId"`
	Blank   questionBankVoResp.SupplyBlankPractice `json:"blankComponent"`
	Score   *int                                   `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
}
type ProgramComponent struct {
	MergeId uint                               `json:"mergeId"`
	Program questionBankVoResp.ProgramPractice `json:"programComponent"`
}
type ProgramComponent1 struct {
	MergeId uint                               `json:"mergeId"`
	Program questionBankVoResp.ProgramPractice `json:"programComponent"`
	Score   *int                               `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
}
