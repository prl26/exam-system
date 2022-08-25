// 自动生成模板ResourcesTest
package lessondata

import (
	"exam-system/global"
)

// ResourcesTest 结构体
type ResourcesTest struct {
	global.GVA_MODEL
	ResourcesId *int     `json:"resourcesId" form:"resourcesId" gorm:"column:resources_id;comment:关联资源Id;size:32;"`
	QuestionId  *int     `json:"questionId" form:"questionId" gorm:"column:question_id;comment:问题Id;size:32;"`
	Type        *int     `json:"type" form:"type" gorm:"column:type;comment:题目类型;size:8;"`
	Weight      *float64 `json:"weight" form:"weight" gorm:"column:weight;comment:分值比重;size:32;"`
	Orders      *int     `json:"orders" form:"orders" gorm:"column:orders;comment:问题排序;size:8;"`
}

// TableName ResourcesTest 表名
func (ResourcesTest) TableName() string {
	return "les_resources_test"
}
