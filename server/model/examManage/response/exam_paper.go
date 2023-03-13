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
type ExamPaperResponse2 struct {
	PaperId               uint                `json:"paperId"`
	SingleChoiceComponent []ChoiceComponent2  `json:"singleChoiceComponent"`
	MultiChoiceComponent  []ChoiceComponent2  `json:"multiChoiceComponent"`
	JudgeComponent        []JudgeComponent2   `json:"judgeComponent"`
	BlankComponent        []BlankComponent2   `json:"blankComponent"`
	ProgramComponent      []ProgramComponent2 `json:"programComponent"`
	TargetComponent       []STargetComponent  `json:"targetComponent"`
}
type STargetComponent struct {
	Order   string `json:"order"`
	MergeId uint   `json:"mergeId"`
	Target  Target `json:"targetComponent"`
	ScoreStruct
}
type ScoreStruct struct {
	Score         *float64 `json:"score" form:"score" gorm:"column:score;comment:本题分值;size:8;"`
	GotScore      *float64 `json:"gotScore" form:"gotScore" gorm:"column:got_score;comment:该生得分"`
	Answer        string   `json:"answer" form:"answer" gorm:"column:answer;comment:该生题目答案;size:16000;"`
	CorrectAnswer string   `json:"correctAnswer" form:"correctAnswer"`
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
type ChoiceComponent2 struct {
	MergeId uint                                  `json:"mergeId"`
	Order   string                                `json:"order"`
	Choice  questionBankVoResp.MultipleChoiceExam `json:"choiceComponent"`
	ScoreStruct
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
type JudgeComponent2 struct {
	MergeId uint                             `json:"mergeId"`
	Order   string                           `json:"order"`
	Judge   questionBankVoResp.JudgePractice `json:"judgeComponent"`
	ScoreStruct
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
type BlankComponent2 struct {
	MergeId uint                                   `json:"mergeId"`
	Order   string                                 `json:"order"`
	Blank   questionBankVoResp.SupplyBlankPractice `json:"blankComponent"`
	ScoreStruct
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
type ProgramComponent2 struct {
	MergeId uint                               `json:"mergeId"`
	Order   string                             `json:"order"`
	Program questionBankVoResp.ProgramPractice `json:"programComponent"`
	ScoreStruct
}
