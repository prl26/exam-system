package supplyBlank

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/questionBank/po"
	"strconv"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 19:22

 * @Note:

 **/

type SupplyBlankService struct {
}

func (c *SupplyBlankService) Check(choiceQuestionId uint, answer []string) ([]bool, int, error) {
	question, err := c.FindCanPracticeQuestion(choiceQuestionId)
	if err != nil {
		return nil, 0, err
	}
	return c.check(question, answer)
}

func (c *SupplyBlankService) FindCanPracticeQuestion(choiceQuestionId uint) (*po.SupplyBlank, error) {
	var question po.SupplyBlank
	result := global.GVA_DB.Where("id=? and can_practice=?", choiceQuestionId, 1).First(&question)
	if result.Error != nil {
		return nil, fmt.Errorf("找不到该题目")
	}
	return &question, nil
}

func (c *SupplyBlankService) check(question *po.SupplyBlank, checkAnswers []string) (boolList []bool, proportion int, err error) {
	n := len(checkAnswers)
	if n != *question.Num {
		return nil, 0, fmt.Errorf("应该要填入%d个空", n)
	}
	boolList = make([]bool, n)
	answers := strings.Split(question.Answer, ",")
	proportions := strings.Split(question.Proportion, ",")
	if *question.IsOrder == 0 {
		table := make(map[string]int)
		var answerIndex [][]string
		for i, a := range answers {
			split := strings.Split(a, "|")
			for _, v := range split {
				table[v] = i
			}
			answerIndex = append(answerIndex, split)
		}
		for i, answer := range checkAnswers {
			if v, ok := table[answer]; ok {
				index := answerIndex[v]
				for _, s := range index {
					delete(table, s)
				}
				boolList[i] = true
				num, _ := strconv.Atoi(proportions[i])
				proportion += num
			} else {
				boolList[i] = false
			}
		}
	} else {
		for i, answer := range checkAnswers {
			split := strings.Split(answers[i], "|")
			flag := false
			for _, s := range split {
				if s == answer {
					flag = true
					break
				}
			}
			if flag {
				boolList[i] = true
				num, _ := strconv.Atoi(proportions[i])
				proportion += num
			}
		}
	}
	return boolList, proportion, nil
}

func (c *SupplyBlankService) GetAnswer(question *po.SupplyBlank) []string {
	return strings.Split(question.Answer, ",")
}
