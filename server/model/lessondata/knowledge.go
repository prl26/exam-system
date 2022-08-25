// 自动生成模板Knowledge
package lessondata

import (
	"exam-system/global"
)

// Knowledge 结构体
type Knowledge struct {
	global.GVA_MODEL
	CourseId *int   `json:"courseId" form:"courseId" gorm:"column:course_id;comment:课程id;"`
	Name     string `json:"name" form:"name" gorm:"column:name;comment:知识点名称;"`
}

// TableName Knowledge 表名
func (Knowledge) TableName() string {
	return "les_knowledge"
}
