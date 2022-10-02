// 自动生成模板QuestionBankMultipleChoice
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// MultipleChoice 结构体
type MultipleChoice struct {
	global.GVA_MODEL
	BasicModel
	Answer      string    `json:"answer" form:"answer" gorm:"column:answer;comment:"`
	MostOptions int       `json:"mostOptions" form:"mostOptions" gorm:"column:most_options;comment:最多可选项;"`
	Options     []Options `json:"options" gorm:"foreignKey:MultipleChoiceId"`
}

// TableName MultipleChoice 表名
func (MultipleChoice) TableName() string {
	return "les_questionBank_multiple_choice"
}

type MultipleChoiceView struct {
	global.GVA_MODEL
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam     *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
}
