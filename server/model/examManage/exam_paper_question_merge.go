// 自动生成模板PaperQuestionMerge
package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PaperQuestionMerge 结构体
type PaperQuestionMerge struct {
	global.GVA_MODEL
	PaperId      *int `json:"paperId" form:"paperId" gorm:"column:paper_id;comment:试卷id;size:32;"`
	QuestionId   *int `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目id;size:32;"`
	Score        *int `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
	QuestionType *int `json:"questionType" form:"questionType" gorm:"column:question_type;comment:题目类型;size:8;"`
}

// TableName PaperQuestionMerge 表名
func (PaperQuestionMerge) TableName() string {
	return "exam_paper_question_merge"
}
