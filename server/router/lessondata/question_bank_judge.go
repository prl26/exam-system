package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionBankJudgeRouter struct {
}

// InitQuestionBankJudgeRouter 初始化 QuestionBankJudge 路由信息
func (s *QuestionBankJudgeRouter) InitQuestionBankJudgeRouter(Router *gin.RouterGroup) {
	questionBank_judgeRouter := Router.Group("questionBank_judge").Use(middleware.OperationRecord())
	questionBank_judgeRouterWithoutRecord := Router.Group("questionBank_judge")
	var questionBank_judgeApi = v1.ApiGroupApp.LessondataApiGroup.QuestionBankJudgeApi
	{
		questionBank_judgeRouter.POST("createQuestionBankJudge", questionBank_judgeApi.CreateQuestionBankJudge)   // 新建QuestionBankJudge
		questionBank_judgeRouter.DELETE("deleteQuestionBankJudge", questionBank_judgeApi.DeleteQuestionBankJudge) // 删除QuestionBankJudge
		questionBank_judgeRouter.DELETE("deleteQuestionBankJudgeByIds", questionBank_judgeApi.DeleteQuestionBankJudgeByIds) // 批量删除QuestionBankJudge
		questionBank_judgeRouter.PUT("updateQuestionBankJudge", questionBank_judgeApi.UpdateQuestionBankJudge)    // 更新QuestionBankJudge
	}
	{
		questionBank_judgeRouterWithoutRecord.GET("findQuestionBankJudge", questionBank_judgeApi.FindQuestionBankJudge)        // 根据ID获取QuestionBankJudge
		questionBank_judgeRouterWithoutRecord.GET("getQuestionBankJudgeList", questionBank_judgeApi.GetQuestionBankJudgeList)  // 获取QuestionBankJudge列表
	}
}
