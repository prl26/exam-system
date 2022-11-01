package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/Backstage"
	"github.com/prl26/exam-system/server/middleware"
)

type ResourceRouter struct {
}

// InitResourceRouter 初始化 Resource 路由信息
func (s *ResourceRouter) InitResourceRouter(Router *gin.RouterGroup) {
	resourceRouter := Router.Group("resource").Use(middleware.OperationRecord())
	resourceRouterWithoutRecord := Router.Group("resource")
	var resourceApi = Backstage.ApiGroupApp.BasicdataApiGroup.ResourceApi
	{
		resourceRouter.POST("createResource", resourceApi.CreateResource)             // 新建Resource
		resourceRouter.DELETE("deleteResource", resourceApi.DeleteResource)           // 删除Resource
		resourceRouter.DELETE("deleteResourceByIds", resourceApi.DeleteResourceByIds) // 批量删除Resource
		resourceRouter.PUT("updateResource", resourceApi.UpdateResource)              // 更新Resource
	}
	{
		resourceRouterWithoutRecord.GET("findResource", resourceApi.FindResource)       // 根据ID获取Resource
		resourceRouterWithoutRecord.GET("getResourceList", resourceApi.GetResourceList) // 获取Resource列表
	}
}
