// 自动生成模板ResourcePractice
package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ResourcePractice 结构体
type ResourcePractice struct {
	global.GVA_MODEL
	QuestionId  *int `json:"questionId" form:"questionId" gorm:"column:question_id;comment:问题Id;size:64;"`
	Type        *int `json:"type" form:"type" gorm:"column:type;comment:题目类型;size:8;"`
	Orders      *int `json:"orders" form:"orders" gorm:"column:orders;comment:问题排序;size:32;"`
	ResourcesId *int `json:"resourcesId" form:"resourcesId" gorm:"column:resources_id;comment:关联资源Id;size:32;"`
}

// TableName ResourcePractice 表名
func (ResourcePractice) TableName() string {
	return "les_resources_practice"
}
