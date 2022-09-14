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

var questionBankService = service.ServiceGroupApp.QuestionBankServiceGroup.QuestionBankService

//
//  QuestionBankApi
//  @Description: 给前台使用的
//

type QuestionBankApi struct{}

// FindQuestionsByChapterId 根据章节ID获取所有题库内的练习题
// @Tags QuestionBank
// @Summary 根据章节ID获取所有题库内的练习题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank/findQuestionsByChapterId [get]
func (q *QuestionBankApi) FindQuestionsByChapterId(c *gin.Context) {
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
