// 自动生成模板QuestionBankKnowledgeMerge
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// QuestionBankKnowledgeMerge 结构体
type ChapterMerge struct {
	global.GVA_MODEL
	ChapterId    *int `json:"chapterId" form:"chapterId" gorm:"column:chapter_id;comment:章节ID;"`
	QuestionId   *int `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目id;"`
	QuestionType *int `json:"questionType" form:"questionType" gorm:"column:question_type;comment:题目类型;"`
	Difficulty   *int `json:"difficulty" form:"difficulty" gorm:"column:difficulty;comment:难度;"`
	CanPractice  *int `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否为训练题目;"`
	CanExam      *int `json:"CanExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
}

// TableName QuestionBankChapterMerge 表名
func (ChapterMerge) TableName() string {
	return "les_questionBank_chapter_merge"
}
