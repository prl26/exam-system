package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionBankMultipleChoiceRouter struct {
}

// InitQuestionBankMultipleChoiceRouter 初始化 MultipleChoice 路由信息
func (s *QuestionBankMultipleChoiceRouter) InitQuestionBankMultipleChoiceRouter(Router *gin.RouterGroup) {
	questionBank_multiple_choiceRouter := Router.Group("questionBank_multiple_choice").Use(middleware.OperationRecord())
	questionBank_multiple_choiceRouterWithoutRecord := Router.Group("questionBank_multiple_choice")
	var questionBank_multiple_choiceApi = v1.ApiGroupApp.QuestionBankApiGroup.QuestionBankMultipleChoiceApi
	{
		questionBank_multiple_choiceRouter.POST("createQuestionBankMultipleChoice", questionBank_multiple_choiceApi.CreateQuestionBankMultipleChoice)             // 新建QuestionBankMultipleChoice
		questionBank_multiple_choiceRouter.DELETE("deleteQuestionBankMultipleChoice", questionBank_multiple_choiceApi.DeleteQuestionBankMultipleChoice)           // 删除QuestionBankMultipleChoice
		questionBank_multiple_choiceRouter.DELETE("deleteQuestionBankMultipleChoiceByIds", questionBank_multiple_choiceApi.DeleteQuestionBankMultipleChoiceByIds) // 批量删除QuestionBankMultipleChoice
		questionBank_multiple_choiceRouter.PUT("updateQuestionBankMultipleChoice", questionBank_multiple_choiceApi.UpdateQuestionBankMultipleChoice)              // 更新QuestionBankMultipleChoice
	}
	{
		questionBank_multiple_choiceRouterWithoutRecord.GET("findQuestionBankMultipleChoice", questionBank_multiple_choiceApi.FindQuestionBankMultipleChoice)       // 根据ID获取QuestionBankMultipleChoice
		questionBank_multiple_choiceRouterWithoutRecord.GET("getQuestionBankMultipleChoiceList", questionBank_multiple_choiceApi.GetQuestionBankMultipleChoiceList) // 获取QuestionBankMultipleChoice列表
	}
}
