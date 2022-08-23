// 自动生成模板PaperTemplateItem
package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// PaperTemplateItem 结构体
type PaperTemplateItem struct {
      global.GVA_MODEL
      Chapter  string `json:"chapter" form:"chapter" gorm:"column:chapter;comment:章节;size:32;"`
      ProblemType  *int `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:题目类型;size:32;"`
      Difficulty  *int `json:"difficulty" form:"difficulty" gorm:"column:difficulty;comment:难度;size:32;"`
      Num  *int `json:"num" form:"num" gorm:"column:num;comment:数量;size:32;"`
      Score  *int `json:"score" form:"score" gorm:"column:score;comment:分数;size:32;"`
      TemplateId  *int `json:"templateId" form:"templateId" gorm:"column:template_id;comment:试卷模板id;size:32;"`
}


// TableName PaperTemplateItem 表名
func (PaperTemplateItem) TableName() string {
  return "exam_paper_template_item"
}

