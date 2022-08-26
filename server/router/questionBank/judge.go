package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type QuestionBankJudgeRouter struct {
}

// InitQuestionBankJudgeRouter 初始化 QuestionBankJudge 路由信息
func (s *QuestionBankJudgeRouter) InitQuestionBankJudgeRouter(Router *gin.RouterGroup) {
	questionBank_judgeRouter := Router.Group("questionBankJudge").Use(middleware.OperationRecord())
	questionBank_judgeRouterWithoutRecord := Router.Group("questionBankJudge")
	var questionBank_judgeApi = v1.ApiGroupApp.QuestionBankApiGroup.QuestionBankJudgeApi
	{
		questionBank_judgeRouter.POST("createQuestionBankJudge", questionBank_judgeApi.CreateQuestionBankJudge)             // 新建QuestionBankJudge
		questionBank_judgeRouter.DELETE("deleteQuestionBankJudge", questionBank_judgeApi.DeleteQuestionBankJudge)           // 删除QuestionBankJudge
		questionBank_judgeRouter.DELETE("deleteQuestionBankJudgeByIds", questionBank_judgeApi.DeleteQuestionBankJudgeByIds) // 批量删除QuestionBankJudge
		questionBank_judgeRouter.PUT("updateQuestionBankJudge", questionBank_judgeApi.UpdateQuestionBankJudge)              // 更新QuestionBankJudge
	}
	{
		questionBank_judgeRouterWithoutRecord.GET("findQuestionBankJudge", questionBank_judgeApi.FindQuestionBankJudge)       // 根据ID获取QuestionBankJudge
		questionBank_judgeRouterWithoutRecord.GET("getQuestionBankJudgeList", questionBank_judgeApi.GetQuestionBankJudgeList) // 获取QuestionBankJudge列表
	}
}
