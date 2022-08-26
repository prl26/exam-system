package lessondata

import (
	"exam-system/api/v1"
	"exam-system/middleware"
	"github.com/gin-gonic/gin"
)

type ResourcePracticeRouter struct {
}

// InitResourcePracticeRouter 初始化 ResourcePractice 路由信息
func (s *ResourcePracticeRouter) InitResourcePracticeRouter(Router *gin.RouterGroup) {
	resourcesPracticeRouter := Router.Group("resourcesPractice").Use(middleware.OperationRecord())
	resourcesPracticeRouterWithoutRecord := Router.Group("resourcesPractice")
	var resourcesPracticeApi = v1.ApiGroupApp.LessondataApiGroup.ResourcePracticeApi
	{
		resourcesPracticeRouter.POST("createResourcePractice", resourcesPracticeApi.CreateResourcePractice)             // 新建ResourcePractice
		resourcesPracticeRouter.DELETE("deleteResourcePractice", resourcesPracticeApi.DeleteResourcePractice)           // 删除ResourcePractice
		resourcesPracticeRouter.DELETE("deleteResourcePracticeByIds", resourcesPracticeApi.DeleteResourcePracticeByIds) // 批量删除ResourcePractice
		resourcesPracticeRouter.PUT("updateResourcePractice", resourcesPracticeApi.UpdateResourcePractice)              // 更新ResourcePractice
	}
	{
		resourcesPracticeRouterWithoutRecord.GET("findResourcePractice", resourcesPracticeApi.FindResourcePractice)       // 根据ID获取ResourcePractice
		resourcesPracticeRouterWithoutRecord.GET("getResourcePracticeList", resourcesPracticeApi.GetResourcePracticeList) // 获取ResourcePractice列表
	}
}
