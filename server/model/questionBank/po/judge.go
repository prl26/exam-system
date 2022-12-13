package po

import (
	"github.com/prl26/exam-system/server/global"
)

type Judge struct {
	global.GVA_MODEL
	CourseSupport
	JudgeModel
}

type JudgeModel struct {
	IsRight *bool `json:"isRight" form:"isRight" gorm:"column:is_right;comment:是否正确;"`
	BasicModel
}

// TableName QuestionBankJudge 表名
func (Judge) TableName() string {
	return "les_questionBank_judge"
}
