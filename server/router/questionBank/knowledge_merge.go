package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionBankKnowledgeMergeRouter struct {
}

// InitQuestionBankKnowledgeMergeRouter 初始化 QuestionBankKnowledgeMerge 路由信息
func (s *QuestionBankKnowledgeMergeRouter) InitQuestionBankKnowledgeMergeRouter(Router *gin.RouterGroup) {
	questionBankKnowledgeMergeRouter := Router.Group("questionBankKnowledgeMerge").Use(middleware.OperationRecord())
	questionBankKnowledgeMergeRouterWithoutRecord := Router.Group("questionBankKnowledgeMerge")
	var questionBankKnowledgeMergeApi = v1.ApiGroupApp.QuestionBankApiGroup.QuestionBankKnowledgeMergeApi
	{
		questionBankKnowledgeMergeRouter.POST("createQuestionBankKnowledgeMerge", questionBankKnowledgeMergeApi.CreateQuestionBankKnowledgeMerge)             // 新建QuestionBankKnowledgeMerge
		questionBankKnowledgeMergeRouter.DELETE("deleteQuestionBankKnowledgeMerge", questionBankKnowledgeMergeApi.DeleteQuestionBankKnowledgeMerge)           // 删除QuestionBankKnowledgeMerge
		questionBankKnowledgeMergeRouter.DELETE("deleteQuestionBankKnowledgeMergeByIds", questionBankKnowledgeMergeApi.DeleteQuestionBankKnowledgeMergeByIds) // 批量删除QuestionBankKnowledgeMerge
		questionBankKnowledgeMergeRouter.PUT("updateQuestionBankKnowledgeMerge", questionBankKnowledgeMergeApi.UpdateQuestionBankKnowledgeMerge)              // 更新QuestionBankKnowledgeMerge
	}
	{
		questionBankKnowledgeMergeRouterWithoutRecord.GET("findQuestionBankKnowledgeMerge", questionBankKnowledgeMergeApi.FindQuestionBankKnowledgeMerge)       // 根据ID获取QuestionBankKnowledgeMerge
		questionBankKnowledgeMergeRouterWithoutRecord.GET("getQuestionBankKnowledgeMergeList", questionBankKnowledgeMergeApi.GetQuestionBankKnowledgeMergeList) // 获取QuestionBankKnowledgeMerge列表
	}
}
