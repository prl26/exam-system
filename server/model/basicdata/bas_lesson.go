// 自动生成模板Lesson
package basicdata

import (
	"github.com/prl26/exam-system/server/global"
)

// Lesson 结构体
type Lesson struct {
	global.GVA_MODEL
	Name     string     `json:"name" form:"name" gorm:"column:name;comment:课程的名称;size:32;"`
	Chapters []*Chapter `json:"chapters"`
}

// TableName Lesson 表名
func (Lesson) TableName() string {
	return "bas_lesson"
}
