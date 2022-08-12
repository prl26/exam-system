// 自动生成模板QuestionBankKnowledgeMerge
package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// QuestionBankKnowledgeMerge 结构体
type QuestionBankKnowledgeMerge struct {
      global.GVA_MODEL
      Knowledge_id  *int `json:"knowledge_id" form:"knowledge_id" gorm:"column:knowledge_id;comment:知识点id;"`
      Question_id  *int `json:"question_id" form:"question_id" gorm:"column:question_id;comment:题目id;"`
      Question_type  *int `json:"question_type" form:"question_type" gorm:"column:question_type;comment:题目类型;"`
      Difficulty  *int `json:"difficulty" form:"difficulty" gorm:"column:difficulty;comment:难度;"`
}


// TableName QuestionBankKnowledgeMerge 表名
func (QuestionBankKnowledgeMerge) TableName() string {
  return "les_questionBank_knowledge_merge"
}

