package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionBankProgrammRouter struct {
}

// InitQuestionBankProgrammRouter 初始化 Programm 路由信息
func (s *QuestionBankProgrammRouter) InitQuestionBankProgrammRouter(Router *gin.RouterGroup) {
	questionBankProgrammRouter := Router.Group("questionBankProgramm").Use(middleware.OperationRecord())
	questionBankProgrammRouterWithoutRecord := Router.Group("questionBankProgramm")
	var questionBankProgrammApi = v1.ApiGroupApp.QuestionBankApiGroup.QuestionBankProgrammApi
	{
		questionBankProgrammRouter.POST("createQuestionBankProgramm", questionBankProgrammApi.CreateQuestionBankProgramm)             // 新建QuestionBankProgramm
		questionBankProgrammRouter.DELETE("deleteQuestionBankProgramm", questionBankProgrammApi.DeleteQuestionBankProgramm)           // 删除QuestionBankProgramm
		questionBankProgrammRouter.DELETE("deleteQuestionBankProgrammByIds", questionBankProgrammApi.DeleteQuestionBankProgrammByIds) // 批量删除QuestionBankProgramm
		questionBankProgrammRouter.PUT("updateQuestionBankProgramm", questionBankProgrammApi.UpdateQuestionBankProgramm)              // 更新QuestionBankProgramm
	}
	{
		questionBankProgrammRouterWithoutRecord.GET("findQuestionBankProgramm", questionBankProgrammApi.FindQuestionBankProgramm)       // 根据ID获取QuestionBankProgramm
		questionBankProgrammRouterWithoutRecord.GET("getQuestionBankProgrammList", questionBankProgrammApi.GetQuestionBankProgrammList) // 获取QuestionBankProgramm列表
	}
}
