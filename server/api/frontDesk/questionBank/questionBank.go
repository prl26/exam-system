package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"strconv"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 16:42

 * @Note:

 **/

type CommonApi struct{}

var questionBankService = service.ServiceGroupApp.QuestionBankServiceGroup.QuestionBankService

//FindQuestionsByChapterId 通过chapterId 获取所有题目
func (q *CommonApi) FindQuestionsByChapterId(c *gin.Context) {
	query := c.Query("id")
	idInt, err := strconv.Atoi(query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	idUint := uint(idInt)
	questions := questionBankService.FindQuestions(idUint)
	response.OkWithData(questions, c)
	return
}

//func (q*CommonApi) FindJudgesByChapterId(c *gin.Context){
//	query := c.Query("id")
//	idInt, err := strconv.Atoi(query)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	idUint := uint(idInt)
//}
