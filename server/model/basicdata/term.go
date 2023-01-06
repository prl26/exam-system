// 自动生成模板Term
package basicdata

import (
	"github.com/prl26/exam-system/server/global"
)

// Term 结构体
type Term struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:学期名称;size:255;"`
	StartTime int64  `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间;"`
	EndTime   int64  `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间;"`
}

// TableName Term 表名
func (Term) TableName() string {
	return "bas_term"
}
