package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KnowledgeRouter struct {
}

// InitKnowledgeRouter 初始化 Knowledge 路由信息
func (s *KnowledgeRouter) InitKnowledgeRouter(Router *gin.RouterGroup) {
	knowledgeRouter := Router.Group("knowledge").Use(middleware.OperationRecord())
	knowledgeRouterWithoutRecord := Router.Group("knowledge")
	var knowledgeApi = v1.ApiGroupApp.BasicdataApiGroup.KnowledgeApi
	{
		knowledgeRouter.POST("createKnowledge", knowledgeApi.CreateKnowledge)   // 新建Knowledge
		knowledgeRouter.DELETE("deleteKnowledge", knowledgeApi.DeleteKnowledge) // 删除Knowledge
		knowledgeRouter.DELETE("deleteKnowledgeByIds", knowledgeApi.DeleteKnowledgeByIds) // 批量删除Knowledge
		knowledgeRouter.PUT("updateKnowledge", knowledgeApi.UpdateKnowledge)    // 更新Knowledge
	}
	{
		knowledgeRouterWithoutRecord.GET("findKnowledge", knowledgeApi.FindKnowledge)        // 根据ID获取Knowledge
		knowledgeRouterWithoutRecord.GET("getKnowledgeList", knowledgeApi.GetKnowledgeList)  // 获取Knowledge列表
	}
}
