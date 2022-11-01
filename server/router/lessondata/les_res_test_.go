package lessondata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/Backstage"
	"github.com/prl26/exam-system/server/middleware"
)

type ResourcesTestRouter struct {
}

// InitResourcesTestRouter 初始化 ResourcesTest 路由信息
func (s *ResourcesTestRouter) InitResourcesTestRouter(Router *gin.RouterGroup) {
	resourcesTestRouter := Router.Group("resourcesTest").Use(middleware.OperationRecord())
	resourcesTestRouterWithoutRecord := Router.Group("resourcesTest")
	var resourcesTestApi = Backstage.ApiGroupApp.LessondataApiGroup.ResourcesTestApi
	{
		resourcesTestRouter.POST("createResourcesTest", resourcesTestApi.CreateResourcesTest)             // 新建ResourcesTest
		resourcesTestRouter.DELETE("deleteResourcesTest", resourcesTestApi.DeleteResourcesTest)           // 删除ResourcesTest
		resourcesTestRouter.DELETE("deleteResourcesTestByIds", resourcesTestApi.DeleteResourcesTestByIds) // 批量删除ResourcesTest
		resourcesTestRouter.PUT("updateResourcesTest", resourcesTestApi.UpdateResourcesTest)              // 更新ResourcesTest
	}
	{
		resourcesTestRouterWithoutRecord.GET("findResourcesTest", resourcesTestApi.FindResourcesTest)       // 根据ID获取ResourcesTest
		resourcesTestRouterWithoutRecord.GET("getResourcesTestList", resourcesTestApi.GetResourcesTestList) // 获取ResourcesTest列表
	}
}
