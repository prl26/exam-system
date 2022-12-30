package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type TargetRouter struct {
}

//// InitQuestionBankSupplyBlankRouter 初始化 QuestionBankSupplyBlank 路由信息
func (s *TargetRouter) InitTargetRouter(Router *gin.RouterGroup) {
	supplyBlankRouter := Router.Group("target").Use(middleware.OperationRecord())
	supplyBlankRouterWithoutRecord := Router.Group("target")
	var api = api.ApiGroupApp.BackStage.QuestionBankApiGroup.TargetApi
	{
		supplyBlankRouter.POST("create", api.Create)   // 新建QuestionBankSupplyBlank
		supplyBlankRouter.DELETE("delete", api.Delete) // 删除QuestionBankSupplyBlank
		supplyBlankRouter.PUT("update", api.Update)    // 更新QuestionBankSupplyBlank
	}
	{
		supplyBlankRouterWithoutRecord.GET("findDetail", api.FindDetail) // 根据ID获取QuestionBankSupplyBlank
		supplyBlankRouterWithoutRecord.GET("findList", api.FindList)     // 获取QuestionBankSupplyBlank列表
	}
}
