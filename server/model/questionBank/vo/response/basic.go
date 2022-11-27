package response

import "github.com/prl26/exam-system/server/global"

type PracticeModel struct {
	global.GVA_MODEL
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	Describe    string `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
}
