// 自动生成模板PaperQuestionMerge
package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PaperQuestionMerge 结构体
type PaperQuestionMerge struct {
	global.GVA_MODEL
	PaperId      *uint `json:"paperId" form:"paperId" gorm:"column:paper_id;comment:试卷id;size:32;"`
	QuestionId   *uint `json:"questionId" form:"paperId" gorm:"column:question_id;comment:题目id;size:32;"`
	Score        *int  `json:"score" form:"score" gorm:"column:score;comment:所占分值;size:8;"`
	QuestionType *int  `json:"questionType" form:"paperId" gorm:"column:question_type;comment:题目类型;size:8;"`
	ProblemType  *int  `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:难度;size:8;"`
}

// TableName PaperQuestionMerge 表名
func (PaperQuestionMerge) TableName() string {
	return "exam_paper_question_merge"
}
