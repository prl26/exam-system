// 自动生成模板ExamPaper
package examManage

import (
	"github.com/prl26/exam-system/server/global"
)

// ExamPaper 结构体
type ExamPaper struct {
	global.GVA_MODEL
	PlanId     *int   `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	TemplateId *int   `json:"templateId" form:"templateId" gorm:"column:template_id;comment:试卷模板Id;size:32;"`
	TermId     uint   `json:"termId" from:"termId"`
	LessonId   uint   `json:"lessonId" form:"lessonId"`
	UserId     *uint  `json:"user_id" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
}

type ExamPaper1 struct {
	global.GVA_MODEL
	PlanId     *int                 `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name       string               `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	TemplateId *int                 `json:"templateId" form:"templateId" gorm:"column:template_id;comment:试卷模板Id;size:32;"`
	TermId     uint                 `json:"termId" from:"termId"`
	LessonId   uint                 `json:"lessonId" form:"lessonId"`
	UserId     *uint                `json:"user_id" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
	PaperItem  []PaperQuestionMerge `json:"paperItem" gorm:"foreignKey:paper_id"`
}
type Product struct {
	Id   int64  `json:"id"` //字段一定要大写不然各种问题
	Name string `json:"name"`
}

// TableName ExamPaper 表名
func (ExamPaper) TableName() string {
	return "exam_paper"
}
func (ExamPaper1) TableName() string {
	return "exam_paper"
}
