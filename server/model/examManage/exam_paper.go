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
	CourseId   uint   `json:"courseId" form:"courseId"`
	UserId     *int   `json:"user_id" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
}

// TableName ExamPaper 表名
func (ExamPaper) TableName() string {
	return "exam_paper"
}
