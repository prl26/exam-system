// 自动生成模板Class
package basicdata

import (
	"exam-system/global"
)

// Class 结构体
type Class struct {
	global.GVA_MODEL
	Name           string `json:"name" form:"name" gorm:"column:name;comment:班级名称;size:255;"`
	ProfessionalId *int   `json:"professionalId" form:"professionalId" gorm:"column:professional_id;comment:专业id;"`
}

// TableName Class 表名
func (Class) TableName() string {
	return "bas_class"
}
