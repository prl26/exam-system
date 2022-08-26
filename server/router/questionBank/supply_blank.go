package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type QuestionBankSupplyBlankRouter struct {
}

// InitQuestionBankSupplyBlankRouter 初始化 QuestionBankSupplyBlank 路由信息
func (s *QuestionBankSupplyBlankRouter) InitQuestionBankSupplyBlankRouter(Router *gin.RouterGroup) {
	questionBank_supply_blankRouter := Router.Group("questionBankSupplyBlank").Use(middleware.OperationRecord())
	questionBank_supply_blankRouterWithoutRecord := Router.Group("questionBankSupplyBlank")
	var questionBank_supply_blankApi = v1.ApiGroupApp.QuestionBankApiGroup.QuestionBankSupplyBlankApi
	{
		questionBank_supply_blankRouter.POST("createQuestionBankSupplyBlank", questionBank_supply_blankApi.CreateQuestionBankSupplyBlank)             // 新建QuestionBankSupplyBlank
		questionBank_supply_blankRouter.DELETE("deleteQuestionBankSupplyBlank", questionBank_supply_blankApi.DeleteQuestionBankSupplyBlank)           // 删除QuestionBankSupplyBlank
		questionBank_supply_blankRouter.DELETE("deleteQuestionBankSupplyBlankByIds", questionBank_supply_blankApi.DeleteQuestionBankSupplyBlankByIds) // 批量删除QuestionBankSupplyBlank
		questionBank_supply_blankRouter.PUT("updateQuestionBankSupplyBlank", questionBank_supply_blankApi.UpdateQuestionBankSupplyBlank)              // 更新QuestionBankSupplyBlank
	}
	{
		questionBank_supply_blankRouterWithoutRecord.GET("findQuestionBankSupplyBlank", questionBank_supply_blankApi.FindQuestionBankSupplyBlank)       // 根据ID获取QuestionBankSupplyBlank
		questionBank_supply_blankRouterWithoutRecord.GET("getQuestionBankSupplyBlankList", questionBank_supply_blankApi.GetQuestionBankSupplyBlankList) // 获取QuestionBankSupplyBlank列表
	}
}
