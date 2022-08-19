package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TermRouter struct {
}

// InitTermRouter 初始化 Term 路由信息
func (s *TermRouter) InitTermRouter(Router *gin.RouterGroup) {
	termRouter := Router.Group("term").Use(middleware.OperationRecord())
	termRouterWithoutRecord := Router.Group("term")
	var termApi = v1.ApiGroupApp.BasicdataApiGroup.TermApi
	{
		termRouter.POST("createTerm", termApi.CreateTerm)   // 新建Term
		termRouter.DELETE("deleteTerm", termApi.DeleteTerm) // 删除Term
		termRouter.DELETE("deleteTermByIds", termApi.DeleteTermByIds) // 批量删除Term
		termRouter.PUT("updateTerm", termApi.UpdateTerm)    // 更新Term
	}
	{
		termRouterWithoutRecord.GET("findTerm", termApi.FindTerm)        // 根据ID获取Term
		termRouterWithoutRecord.GET("getTermList", termApi.GetTermList)  // 获取Term列表
	}
}
