package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/Backstage"
	"github.com/prl26/exam-system/server/middleware"
)

type QuestionBankJudgeRouter struct {
}

// InitQuestionBankJudgeRouter 初始化 QuestionBankJudge 路由信息
func (s *QuestionBankJudgeRouter) InitQuestionBankJudgeRouter(Router *gin.RouterGroup) {
	judgeRouter := Router.Group("judge").Use(middleware.OperationRecord())
	judgeRouterWithoutRecord := Router.Group("judge")
	var judgeApi = Backstage.ApiGroupApp.QuestionBankApiGroup.JudgeApi
	{
		judgeRouter.POST("create", judgeApi.Create)   // 新建QuestionBankJudge
		judgeRouter.DELETE("delete", judgeApi.Delete) // 删除QuestionBankJudge
		judgeRouter.PUT("update", judgeApi.Update)    // 更新QuestionBankJudge
	}
	{
		judgeRouterWithoutRecord.GET("findDetail", judgeApi.FindDetail)  // 根据ID获取QuestionBankJudge
		judgeRouterWithoutRecord.GET("findList", judgeApi.FindJudgeList) // 获取QuestionBankJudge列表
	}
}
