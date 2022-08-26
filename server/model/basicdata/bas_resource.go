// 自动生成模板Resource
package basicdata

import (
	"exam-system/global"
)

// Resource 结构体
type Resource struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name" gorm:"column:name;comment:课程资源名称;size:32;"`
	Type     *int   `json:"type" form:"type" gorm:"column:type;comment:课程资源类型;size:32;"`
	DetailId string `json:"detailId" form:"detailId" gorm:"column:detail_id;comment:课程资源详情id;size:32;"`
}

// TableName Resource 表名
func (Resource) TableName() string {
	return "bas_resource"
}
