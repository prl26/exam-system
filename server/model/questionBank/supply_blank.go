// 自动生成模板QuestionBankSupplyBlank
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// QuestionBankSupplyBlank 结构体
type SupplyBlank struct {
	global.GVA_MODEL
	IsOrder *int   `json:"isOrder" form:"isOrder" gorm:"column:is_order;comment:是否要求有序;"`
	Num     *int   `json:"num" form:"num" gorm:"column:num;comment:可填项;"`
	Answer  string `json:"answer" form:"answer" gorm:"column:answer;comment:答案"`
	BasicModel
}

// TableName QuestionBankSupplyBlank 表名
func (SupplyBlank) TableName() string {
	return "les_questionBank_supply_blank"
}

type SupplyBlankView struct {
	global.GVA_MODEL
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam     *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
}
