package lessondata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type VideoResourcesRouter struct {
}

// InitVideoResourcesRouter 初始化 VideoResources 路由信息
func (s *VideoResourcesRouter) InitVideoResourcesRouter(Router *gin.RouterGroup) {
	videoResourcesRouter := Router.Group("videoResources").Use(middleware.OperationRecord())
	videoResourcesRouterWithoutRecord := Router.Group("videoResources")
	var videoResourcesApi = v1.ApiGroupApp.LessondataApiGroup.VideoResourcesApi
	{
		videoResourcesRouter.POST("createVideoResources", videoResourcesApi.CreateVideoResources)             // 新建VideoResources
		videoResourcesRouter.DELETE("deleteVideoResources", videoResourcesApi.DeleteVideoResources)           // 删除VideoResources
		videoResourcesRouter.DELETE("deleteVideoResourcesByIds", videoResourcesApi.DeleteVideoResourcesByIds) // 批量删除VideoResources
		videoResourcesRouter.PUT("updateVideoResources", videoResourcesApi.UpdateVideoResources)              // 更新VideoResources
	}
	{
		videoResourcesRouterWithoutRecord.GET("findVideoResources", videoResourcesApi.FindVideoResources)       // 根据ID获取VideoResources
		videoResourcesRouterWithoutRecord.GET("getVideoResourcesList", videoResourcesApi.GetVideoResourcesList) // 获取VideoResources列表
	}
}
