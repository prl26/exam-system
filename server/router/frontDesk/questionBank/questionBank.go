package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
)

type QuestionBankRouter struct {

}

func (s *QuestionBankRouter) InitQuestionBankRouter(Router *gin.RouterGroup) {
	questionRouter := Router.Group("questionBank")
	var questionApi = v1.ApiGroupApp.FrontDesk.QuestionBankGroup.QuestionBankApi
	{
		questionRouter.GET("findQuestionsByKnowledgeId",questionApi.FindQuestionsByKnowledgeId)
	}


}
