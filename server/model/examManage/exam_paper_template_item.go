// 自动生成模板PaperTemplateItem
package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
)

// PaperTemplateItem 结构体
type PaperTemplateItem struct {
	global.GVA_MODEL
	ChapterId    *int                       `json:"chapterId" form:"chapterId" gorm:"column:chapter_id;comment:章节id;size:32;"`
	ProblemType  *int                       `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:难度;size:8;"`
	QuestionType *questionType.QuestionType `json:"questionType" form:"paperId" gorm:"column:question_type;comment:题目类型;size:8;"`
	Num          *int                       `json:"num" form:"num" gorm:"column:num;comment:数量;size:32;"`
	Score        *int                       `json:"score" form:"score" gorm:"column:score;comment:分数;size:32;"`
	TemplateId   *int                       `json:"templateId" form:"templateId" gorm:"column:template_id;comment:试卷模板id;size:32;"`
}

// TableName PaperTemplateItem 表名
func (PaperTemplateItem) TableName() string {
	return "exam_paper_template_item"
}
