package response

import (
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	"time"
)

type QuestionSituation struct {
	StudentId     uint
	RecordCount   uint
	QuestionCount uint
}

type SituationDetail struct {
	*questionBank.BasicModelWith `gorm:"-"`
	Id                           uint
	QuestionType                 questionType.QuestionType
	QuestionId                   uint
	Score                        uint
	CommitTime                   time.Time
}
