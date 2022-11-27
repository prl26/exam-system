package bo

import (
	"github.com/prl26/exam-system/server/global"
	questionBankError "github.com/prl26/exam-system/server/model/questionBank/error"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	"strconv"
	"strings"
)

type SupplyBlankSearchCriteria struct {
	questionBankPo.SimpleModel
	questionBankPo.CourseSupport
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
		err = questionBankError.ScoreError
		return
	}
	return strings.Join(thisAnswers, ","), strings.Join(thisProportion, ","), nil
}

func (this *SupplyBlankAnswers) Deserialization(answer string, proportion string) error {
	answers := strings.Split(answer, ",")
	proportions := strings.Split(proportion, ",")
	*this = make([]*SupplyBlankAnswer, len(answers))
	for i, s := range answers {
		atoi, err := strconv.Atoi(proportions[i])
		if err != nil {
			return err
		}
		(*this)[i] = &SupplyBlankAnswer{
			Answer:     s,
			Proportion: atoi,
		}
	}
	return nil
}

type SupplyBlankAnswer struct {
	Answer     string `json:"answer"`
	Proportion int    `json:"proportion"`
}

type SupplyBlankDetail struct {
	global.GVA_MODEL
	questionBankPo.CourseSupport
	questionBankPo.JudgeModel
	CourseSupportPtr
	Answers SupplyBlankAnswers `json:"answers"`
}

type SupplyBlankPracticeCriteria struct {
	questionBankPo.CourseSupport
}
