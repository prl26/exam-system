package teachplan

import (
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	"time"
)

// PracticeRecord 练习记录表
type PracticeItem struct {
	ID           uint `json:"id" gorm:"primarykey" form:"id"` // 主键ID
	StudentId    uint `json:"studentId"`
	LessonId     uint `json:"lessonId"`
	QuestionType questionType.QuestionType
	QuestionId   uint
	Score        uint
	RecordId     uint
	CommitTime   time.Time
}

func (PracticeItem) TableName() string {
	return "tea_practice_item"
}
