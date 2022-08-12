// 自动生成模板Resandknow
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Resandknow 结构体
type Resandknow struct {
      global.GVA_MODEL
      ResourceId  *int `json:"resourceId" form:"resourceId" gorm:"column:resource_id;comment:课程资源id;size:32;"`
      KnowledgeId  *int `json:"knowledgeId" form:"knowledgeId" gorm:"column:knowledge_id;comment:资源对应知识点id;size:32;"`
}


// TableName Resandknow 表名
func (Resandknow) TableName() string {
  return "bas_res_and_know"
}

