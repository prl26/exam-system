package request

import (
	"fmt"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/questionBank"
	"strconv"
	"strings"
)

type QuestionBankSupplyBlankSearch struct {
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam     *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	request.PageInfo
}
type SupplyBlankCreate struct {
	questionBank.BasicModel
	IsOrder        int                `json:"isOrder" form:"isOrder" gorm:"column:is_order;comment:是否要求有序;"`
	Answers        SupplyBlankAnswers `json:"answers"`
	LessonSupports []*LessonSupport `json:"LessonSupportSupports"`
}

type SupplyBlankAnswers []*SupplyBlankAnswer

//	充血模式
func (this SupplyBlankAnswers) GetAnswersAndProportions() (answer string, proportion string, err error) {

	var thisAnswers []string
	var thisProportion []string
	ans := 0
	for _, answer := range this {
		ans += answer.Proportion
		thisAnswers = append(thisAnswers, answer.Answer)
		thisProportion = append(thisProportion, strconv.Itoa(answer.Proportion))
	}
	if ans != 100 {
		err = fmt.Errorf("占比之和不等于100")
		return
	}
	return strings.Join(thisAnswers, ","), strings.Join(thisProportion, ","), nil
}

type SupplyBlankAnswer struct {
	Answer     string `json:"answer"`
	Proportion int    `json:"proportion"`
}
type SupplyBlankUpdate struct {
	Id uint `json:"id"`
	questionBank.BasicModel
	IsOrder int                `json:"isOrder" form:"isOrder" gorm:"column:is_order;comment:是否要求有序;"`
	Answers SupplyBlankAnswers `json:"answers"`
}
