package teachplan

import (
	"gorm.io/gorm"
	"time"
)

// PracticeRecord 练习记录表
type PracticeRecord struct {
	ID            uint           `json:"id" gorm:"primarykey" form:"id"` // 主键ID
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`                 // 删除时间
	LessonId      uint           `json:"lessonId"`
	StudentId     uint           `json:"studentId"`
	BeginIp       string         `json:"beginIp"`
	BeginTime     time.Time      `json:"beginTime"`
	QuestionCount *uint          //做题数
	IsValid       *bool          //是否有效
}

func (PracticeRecord) TableName() string {
	return "tea_practice_record"
}
