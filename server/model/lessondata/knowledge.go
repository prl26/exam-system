// 自动生成模板Knowledge
package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Knowledge 结构体
type Knowledge struct {
	global.GVA_MODEL
	Course_id *int   `json:"course_id" form:"course_id" gorm:"column:course_id;comment:课程id;"`
	Name      string `json:"name" form:"name" gorm:"column:name;comment:知识点名称;"`
}

// TableName Knowledge 表名
func (Knowledge) TableName() string {
	return "les_knowledge"
}
