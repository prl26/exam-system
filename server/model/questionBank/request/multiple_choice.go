package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank"
)

type MultipleChoiceFindList struct {
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam     *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	request.PageInfo
}

type MultipleChoiceCreate struct {
	questionBank.MultipleChoice
	LessonSupports []*LessonSupport `json:"LessonSupportSupports"`
}

type MultipleChoiceUpdate struct {
	questionBank.MultipleChoice
}
