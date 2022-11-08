package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/service"
	"strconv"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 16:42

 * @Note:

 **/

type QuestionBankApi struct{}

var questionBankService = service.ServiceGroupApp.QuestionBankServiceGroup.QuestionBankService

//FindQuestionsByChapterId 通过chapterId 获取所有题目
//func (q *QuestionBankApi) FindQuestionsByChapterId(c *gin.Context) {
//	query := c.Query("id")
//	idInt, err := strconv.Atoi(query)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	idUint := uint(idInt)
//	questions := questionBankService.FindQuestions(idUint)
//	response.OkWithData(questions, c)
//	return
//}

//func (q*QuestionBankApi) FindJudgesByChapterId(c *gin.Context){
//	query := c.Query("id")
//	idInt, err := strconv.Atoi(query)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	idUint := uint(idInt)
//}

func (q *QuestionBankApi) FindQuestionsByChapterId(c *gin.Context) {
	chapterId, err := strconv.Atoi(c.Query("chapterId"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	questionT, err := strconv.Atoi(c.Query("questionType"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	switch questionT {
	case questionType.JUDGE:
		judges := questionBankService.FindJudges(uint(chapterId))
		response.OkWithData(judges, c)
		break
	case questionType.PROGRAM:
		programms := questionBankService.FindProgramms(uint(chapterId))
		response.OkWithData(programms, c)
		break
	case questionType.MULTIPLE_CHOICE:
		multipleChoics := questionBankService.FindMultipleChoices(uint(chapterId))
		response.OkWithData(multipleChoics, c)
		break
	case questionType.SUPPLY_BLANK:
		supplyBlanks := questionBankService.FindSupplyBlank(uint(chapterId))
		response.OkWithData(supplyBlanks, c)
		break
	default:
		response.FailWithMessage("题型输入错误", c)
	}
}
