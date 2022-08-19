// 自动生成模板Term
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Term 结构体
type Term struct {
	global.GVA_MODEL
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:学期名称;size:255;"`
	Start_time *time.Time `json:"start_time" form:"start_time" gorm:"column:start_time;comment:开始时间;"`
	End_time   *time.Time `json:"end_time" form:"end_time" gorm:"column:end_time;comment:结束时间;"`
}

// TableName Term 表名
func (Term) TableName() string {
	return "bas_term"
}
