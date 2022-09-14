package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 18:56

 * @Note:

 **/
type QuestionBankRouter struct{}

func (s *QuestionBankProgrammCaseRouter) InitQuestionBankRouter(Router *gin.RouterGroup) {
	questionBankRouter := Router.Group("questionBank").Use(middleware.OperationRecord())
	_ = Router.Group("questionBank")
	var questionBankApi = v1.ApiGroupApp.QuestionBankApiGroup.QuestionBankApi
	{
		questionBankRouter.GET("findQuestionsByChapterId", questionBankApi.FindQuestionsByChapterId)
	}
	{

	}
}
