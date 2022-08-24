package questionBank

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 14:56

 * @Note:

 **/

type RouterGroup struct {
	QuestionBankKnowledgeMergeRouter
	QuestionBankMultipleChoiceRouter
	QuestionBankJudgeRouter
	QuestionBankSupplyBlankRouter
	QuestionBankOptionsRouter

	QuestionBankProgrammRouter
	QuestionBankProgrammCaseRouter
}
