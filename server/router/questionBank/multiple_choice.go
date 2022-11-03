package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type QuestionBankMultipleChoiceRouter struct {
}

// InitQuestionBankMultipleChoiceRouter 初始化 MultipleChoice 路由信息
func (s *QuestionBankMultipleChoiceRouter) InitQuestionBankMultipleChoiceRouter(Router *gin.RouterGroup) {
	multipleChoiceRouter := Router.Group("multipleChoice").Use(middleware.OperationRecord())
	multipleChoiceRouterWithoutRecord := Router.Group("multipleChoice")
	var api = api.ApiGroupApp.QuestionBankApiGroup.MultipleChoiceApi
	{
		multipleChoiceRouter.POST("create", api.Create)   // 新建QuestionBankMultipleChoice
		multipleChoiceRouter.DELETE("delete", api.Delete) // 批量删除QuestionBankMultipleChoice
		multipleChoiceRouter.PUT("update", api.Update)    // 更新QuestionBankMultipleChoice
	}
	{
		multipleChoiceRouterWithoutRecord.GET("findDetail", api.FindDetail) // 根据ID获取QuestionBankMultipleChoice
		multipleChoiceRouterWithoutRecord.GET("findList", api.FindList)     // 获取QuestionBankMultipleChoice列表
	}
}
