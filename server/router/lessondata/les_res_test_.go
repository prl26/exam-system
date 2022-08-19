package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ResourcesTestRouter struct {
}

// InitResourcesTestRouter 初始化 ResourcesTest 路由信息
func (s *ResourcesTestRouter) InitResourcesTestRouter(Router *gin.RouterGroup) {
	resourcesTestRouter := Router.Group("resourcesTest").Use(middleware.OperationRecord())
	resourcesTestRouterWithoutRecord := Router.Group("resourcesTest")
	var resourcesTestApi = v1.ApiGroupApp.LessondataApiGroup.ResourcesTestApi
	{
		resourcesTestRouter.POST("createResourcesTest", resourcesTestApi.CreateResourcesTest)   // 新建ResourcesTest
		resourcesTestRouter.DELETE("deleteResourcesTest", resourcesTestApi.DeleteResourcesTest) // 删除ResourcesTest
		resourcesTestRouter.DELETE("deleteResourcesTestByIds", resourcesTestApi.DeleteResourcesTestByIds) // 批量删除ResourcesTest
		resourcesTestRouter.PUT("updateResourcesTest", resourcesTestApi.UpdateResourcesTest)    // 更新ResourcesTest
	}
	{
		resourcesTestRouterWithoutRecord.GET("findResourcesTest", resourcesTestApi.FindResourcesTest)        // 根据ID获取ResourcesTest
		resourcesTestRouterWithoutRecord.GET("getResourcesTestList", resourcesTestApi.GetResourcesTestList)  // 获取ResourcesTest列表
	}
}
