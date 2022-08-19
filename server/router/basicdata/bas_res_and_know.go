package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ResandknowRouter struct {
}

// InitResandknowRouter 初始化 Resandknow 路由信息
func (s *ResandknowRouter) InitResandknowRouter(Router *gin.RouterGroup) {
	resandknowRouter := Router.Group("resandknow").Use(middleware.OperationRecord())
	resandknowRouterWithoutRecord := Router.Group("resandknow")
	var resandknowApi = v1.ApiGroupApp.BasicdataApiGroup.ResandknowApi
	{
		resandknowRouter.POST("createResandknow", resandknowApi.CreateResandknow)   // 新建Resandknow
		resandknowRouter.DELETE("deleteResandknow", resandknowApi.DeleteResandknow) // 删除Resandknow
		resandknowRouter.DELETE("deleteResandknowByIds", resandknowApi.DeleteResandknowByIds) // 批量删除Resandknow
		resandknowRouter.PUT("updateResandknow", resandknowApi.UpdateResandknow)    // 更新Resandknow
	}
	{
		resandknowRouterWithoutRecord.GET("findResandknow", resandknowApi.FindResandknow)        // 根据ID获取Resandknow
		resandknowRouterWithoutRecord.GET("getResandknowList", resandknowApi.GetResandknowList)  // 获取Resandknow列表
	}
}
