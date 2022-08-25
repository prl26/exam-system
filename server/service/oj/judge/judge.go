package judge

import (
	"exam-system/global"
	"exam-system/model/questionBank"
	"fmt"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 19:33

 * @Note:

 **/

type JudgeService struct{}

func (s *JudgeService) Check(questionId *uint, answer bool) (bool, error) {
	question, err := s.FindCanPracticeQuestion(questionId)
	if err != nil {
		return false, err
	}
	return s.check(question, answer), nil
}

func (s *JudgeService) FindCanPracticeQuestion(choiceQuestionId *uint) (*questionBank.Judge, error) {
	var question questionBank.Judge
	result := global.GVA_DB.Where("id=? and can_practice=?", choiceQuestionId, 1).First(&question)
	if result.Error != nil {
		return nil, fmt.Errorf("找不到该题目")
	}
	return &question, nil
}

func (s *JudgeService) check(question *questionBank.Judge, checkAnswer bool) bool {
	if *question.IsRight == 1 {
		return checkAnswer
	} else {
		return !checkAnswer
	}
}
func (c *JudgeService) GetAnswer(question *questionBank.Judge) bool {
	return *question.IsRight == 1
}
