package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type TargetRouter struct {
}

// // InitQuestionBankSupplyBlankRouter 初始化 QuestionBankSupplyBlank 路由信息
func (s *TargetRouter) InitTargetRouter(Router *gin.RouterGroup) {
	targetRouter := Router.Group("target").Use(middleware.OperationRecord())
	targetWithoutRecord := Router.Group("target")
	var api = api.ApiGroupApp.BackStage.QuestionBankApiGroup.TargetApi
	{
		targetRouter.POST("create", api.Create)   // 新建QuestionBankSupplyBlank
		targetRouter.DELETE("delete", api.Delete) // 删除QuestionBankSupplyBlank
		targetRouter.PUT("update", api.Update)    // 更新QuestionBankSupplyBlank
		targetRouter.StaticFile("excel", "./static/TargetImport.xlsx")
	}
	{
		targetWithoutRecord.GET("findDetail", api.FindDetail) // 根据ID获取QuestionBankSupplyBlank
		targetWithoutRecord.GET("findList", api.FindList)     // 获取QuestionBankSupplyBlank列表
		targetWithoutRecord.POST("importExcel", api.Import)   // 获取QuestionBankSupplyBlank列表

	}
}
