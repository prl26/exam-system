// 自动生成模板QuestionBankProgramm
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// Programm 结构体
type Programm struct {
	global.GVA_MODEL
	BasicModel
}

// TableName Programm 表名
func (Programm) TableName() string {
	return "les_questionBank_programm"
}

type ProgrammView struct {
	global.GVA_MODEL
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam     *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
}
