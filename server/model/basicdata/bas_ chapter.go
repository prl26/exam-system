// 自动生成模板Chapter
package basicdata

import (
	"github.com/prl26/exam-system/server/global"
)

// Chapter 结构体
type Chapter struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name" gorm:"column:name;comment:章节名称;size:32;"`
	LessonId string `json:"lessonId" form:"lessonId" gorm:"column:lesson_id;comment:所属课程id;size:32;"`
	Order    *int   `json:"order" form:"order" gorm:"column:order;comment:用于排序的字段;size:32;"`
}

// TableName Chapter 表名
func (Chapter) TableName() string {
	return "bas_chapter"
}
