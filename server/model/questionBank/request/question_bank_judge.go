package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank"
)

type QuestionBankJudgeSearch struct {
	questionBank.Judge
	request.PageInfo
}

type JudgeCreate struct {
	questionBank.Judge
	ChapterSupport []uint `json:"chapterSupport"`
}
