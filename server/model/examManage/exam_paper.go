// 自动生成模板ExamPaper
package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExamPaper 结构体
type ExamPaper struct {
	global.GVA_MODEL
	PlanId     *int   `json:"planId" form:"planId" gorm:"column:plan_id;comment:考试计划id;size:32;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:试卷名称;size:64;"`
	TemplateId *int   `json:"templateId" form:"templateId" gorm:"column:template_id;comment:试卷模板Id;size:32;"`
}

// TableName ExamPaper 表名
func (ExamPaper) TableName() string {
	return "exam_paper"
}
