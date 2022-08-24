package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionBankOptionsRouter struct {
}

// InitQuestionBankOptionsRouter 初始化 Options 路由信息
func (s *QuestionBankOptionsRouter) InitQuestionBankOptionsRouter(Router *gin.RouterGroup) {
	questionBank_optionsRouter := Router.Group("questionBank_options").Use(middleware.OperationRecord())
	questionBank_optionsRouterWithoutRecord := Router.Group("questionBank_options")
	var questionBank_optionsApi = v1.ApiGroupApp.QuestionBankApiGroup.QuestionBankOptionsApi
	{
		questionBank_optionsRouter.POST("createQuestionBankOptions", questionBank_optionsApi.CreateQuestionBankOptions)             // 新建QuestionBankOptions
		questionBank_optionsRouter.DELETE("deleteQuestionBankOptions", questionBank_optionsApi.DeleteQuestionBankOptions)           // 删除QuestionBankOptions
		questionBank_optionsRouter.DELETE("deleteQuestionBankOptionsByIds", questionBank_optionsApi.DeleteQuestionBankOptionsByIds) // 批量删除QuestionBankOptions
		questionBank_optionsRouter.PUT("updateQuestionBankOptions", questionBank_optionsApi.UpdateQuestionBankOptions)              // 更新QuestionBankOptions
	}
	{
		questionBank_optionsRouterWithoutRecord.GET("findQuestionBankOptions", questionBank_optionsApi.FindQuestionBankOptions)       // 根据ID获取QuestionBankOptions
		questionBank_optionsRouterWithoutRecord.GET("getQuestionBankOptionsList", questionBank_optionsApi.GetQuestionBankOptionsList) // 获取QuestionBankOptions列表
	}
}
