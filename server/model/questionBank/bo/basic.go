package bo

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
)

type CourseSupportPtr struct {
	Chapter   *basicdata.Chapter
	Knowledge *basicdata.Knowledge
}

type PracticeModel struct {
	global.GVA_MODEL
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	Describe    string `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
}
