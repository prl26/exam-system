// 自动生成模板PaperTemplate
package examManage

import (
	"github.com/prl26/exam-system/server/global"
)

// PaperTemplate 结构体
type PaperTemplate struct {
	global.GVA_MODEL
	CourseId           *int                `json:"course_id" form:"courseId" gorm:"column:course_id;comment:课程id;size:32;"`
	UserId             *int                `json:"user_id" form:"userId" gorm:"column:user_id;comment:创建id;size:32;"`
	Name               string              `json:"name" form:"name" gorm:"column:name;comment:数据模板名称;size:64;"`
	Memo               string              `json:"memo" form:"memo" gorm:"column:memo;comment:备注;size:255;"`
	PaperTemplateItems []PaperTemplateItem `json:"paper_template_items" gorm:"foreignKey:TemplateId"`
}

// TableName PaperTemplate 表名
func (PaperTemplate) TableName() string {
	return "exam_paper_template"
}
