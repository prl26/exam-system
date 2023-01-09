package judge

import (
	"github.com/prl26/exam-system/server/global"
	exception "github.com/prl26/exam-system/server/model/common/error"
	"github.com/prl26/exam-system/server/model/questionBank/po"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 19:33

 * @Note:

 **/

type JudgeService struct{}

func (s *JudgeService) Check(questionId uint, answer bool) (bool, error) {
	question, err := s.FindCanPracticeQuestion(questionId)
	if err != nil {
		return false, exception.NotFoundError
	}
	return s.check(question, answer), nil
}
func (s *JudgeService) ExamCheck(questionId uint, answer bool) (bool, error) {
	question, err := s.FindCanExamQuestion(questionId)
	if err != nil {
		return false, exception.NotFoundError
	}
	return s.check(question, answer), nil
}
func (s *JudgeService) FindCanPracticeQuestion(choiceQuestionId uint) (*po.Judge, error) {
	var question po.Judge
	result := global.GVA_DB.Where("id=? and can_practice=?", choiceQuestionId, 1).First(&question)
	if result.Error != nil {
		return nil, exception.NotFoundError
	}
	return &question, nil
}
func (s *JudgeService) FindCanExamQuestion(choiceQuestionId uint) (*po.Judge, error) {
	var question po.Judge
	result := global.GVA_DB.Where("id=? and can_exam=?", choiceQuestionId, 1).First(&question)
	if result.Error != nil {
		return nil, exception.NotFoundError
	}
	return &question, nil
}
func (s *JudgeService) check(question *po.Judge, checkAnswer bool) bool {
	if *question.IsRight {
		return checkAnswer
	} else {
		return !checkAnswer
	}
}
func (c *JudgeService) GetAnswer(question *po.Judge) bool {
	return *question.IsRight
}
