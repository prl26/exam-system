package multipleChoice

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/questionBank/po"
	"strconv"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 18:48

 * @Note:

 **/

type MultipleChoiceService struct{}

func (c *MultipleChoiceService) Check(choiceQuestionId uint, answer []string) (bool, error) {
	question, err := c.FindCanPracticeQuestion(choiceQuestionId)
	if err != nil {
		return false, err
	}
	return c.check(question, answer), nil
}

func (c *MultipleChoiceService) FindCanPracticeQuestion(choiceQuestionId uint) (*po.MultipleChoice, error) {
	var question po.MultipleChoice
	result := global.GVA_DB.Where("id=? and can_practice=?", choiceQuestionId, 1).First(&question)
	if result.Error != nil {
		return nil, fmt.Errorf("找不到该题目")
	}
	return &question, nil
}

func (c *MultipleChoiceService) check(question *po.MultipleChoice, answer []string) bool {
	//n := len(answer)
	//if n != question.MostOptions {
	//	return false
	//}
	// 前端需要做好的
	//sort.Slice(answer, func(i, j int) bool {
	//	return answer[i]<answer[j]
	//})
	checkAnswer := strings.Join(answer, ",")
	return checkAnswer == question.Answer
}

func (c *MultipleChoiceService) GetAnswer(question *po.MultipleChoice) []int {
	answers := strings.Split(question.Answer, ",")
	result := make([]int, len(answers))
	for i := 0; i < len(answers); i++ {
		result[i], _ = strconv.Atoi(answers[i])
	}
	return result
}
