// 自动生成模板QuestionBankJudge
package questionBank

import (
	"github.com/prl26/exam-system/server/global"
)

// QuestionBankJudge 结构体
type Judge struct {
	global.GVA_MODEL
	IsRight *int `json:"isRight" form:"isRight" gorm:"column:is_right;comment:是否正确;"`
	BasicModel
}

// TableName QuestionBankJudge 表名
func (Judge) TableName() string {
	return "les_questionBank_judge"
}

type JudgeView struct {
	global.GVA_MODEL
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam     *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
}
