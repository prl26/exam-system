// 自动生成模板College
package basicdata

import (
	"exam-system/global"
)

// College 结构体
type College struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:学院名称;size:255;"`
}

// TableName College 表名
func (College) TableName() string {
	return "bas_college"
}
