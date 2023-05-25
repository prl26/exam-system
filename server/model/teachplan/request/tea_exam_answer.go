package request

import "github.com/prl26/exam-system/server/model/questionBank/enum/questionType"

type History struct {
	QuestionType questionType.QuestionType `json:"questionType"`
	Ids          []uint                    `json:"ids"`
}

type Answer struct {
	QuestionType questionType.QuestionType `json:"questionType"`
	Id           uint                      `json:"id"`
}
