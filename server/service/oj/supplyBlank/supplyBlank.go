package supplyBlank

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/questionBank"
	"sort"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 19:22

 * @Note:

 **/

type SupplyBlankService struct {
}

func (c *SupplyBlankService) Check(choiceQuestionId *uint, answer []string) (bool, error) {
	question, err := c.FindCanPracticeQuestion(choiceQuestionId)
	if err != nil {
		return false, err
	}
	return c.check(question, answer), nil
}

func (c *SupplyBlankService) FindCanPracticeQuestion(choiceQuestionId *uint) (*questionBank.SupplyBlank, error) {
	var question questionBank.SupplyBlank
	result := global.GVA_DB.Where("id=? and can_practice=?", choiceQuestionId, 1).First(&question)
	if result.Error != nil {
		return nil, fmt.Errorf("找不到该题目")
	}
	return &question, nil
}

func (c *SupplyBlankService) check(question *questionBank.SupplyBlank, checkAnswers []string) bool {
	n := len(checkAnswers)
	if n != *question.Num {
		return false
	}
	answers := strings.Split(question.Answer, ",")
	if *question.IsOrder == 0 {
		sort.Slice(checkAnswers, func(i, j int) bool {
			return checkAnswers[i] > checkAnswers[j]
		})
		sort.Slice(answers, func(i, j int) bool {
			return answers[i] > answers[j]
		})
	}
	for i := 0; i < n; i++ {
		if checkAnswers[i] != answers[i] {
			return false
		}
	}
	return true
}

func (c *SupplyBlankService) GetAnswer(question *questionBank.SupplyBlank) []string {
	return strings.Split(question.Answer, ",")
}
