package teachplan

import (
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
)

// PracticeAnswer 练习答案表
type PracticeAnswer struct {
	ID           uint `json:"id" gorm:"primarykey" form:"id"` // 主键ID
	StudentId    uint `json:"studentId"`
	QuestionType questionType.QuestionType
	QuestionId   uint
	LessonId     uint
	Score        uint
	Answer       string
}

func (PracticeAnswer) TableName() string {
	return "tea_practice_answer"
}
