// 自动生成模板QuestionBankKnowledgeMerge
package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// QuestionBankKnowledgeMerge 结构体
type KnowledgeMerge struct {
	global.GVA_MODEL
	KnowledgeId  *int `json:"knowledgeId" form:"knowledgeId" gorm:"column:knowledge_id;comment:知识点id;"`
	QuestionId   *int `json:"questionId" form:"questionId" gorm:"column:question_id;comment:题目id;"`
	QuestionType *int `json:"questionType" form:"questionType" gorm:"column:question_type;comment:题目类型;"`
	Difficulty   *int `json:"difficulty" form:"difficulty" gorm:"column:difficulty;comment:难度;"`
	CanPractice  *int `json:"canPractice" form:"canPractice"`
}

// TableName QuestionBankKnowledgeMerge 表名
func (KnowledgeMerge) TableName() string {
	return "les_questionBank_knowledge_merge"
}
