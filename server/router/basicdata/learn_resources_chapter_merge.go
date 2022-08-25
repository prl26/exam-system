package basicdata

import (
	"exam-system/api/v1"
	"exam-system/middleware"
	"github.com/gin-gonic/gin"
)

type LearnResourcesChapterMergeRouter struct {
}

// InitLearnResourcesChapterMergeRouter 初始化 LearnResourcesChapterMerge 路由信息
func (s *LearnResourcesChapterMergeRouter) InitLearnResourcesChapterMergeRouter(Router *gin.RouterGroup) {
	learnResourcesChapterMergeRouter := Router.Group("learnResourcesChapterMerge").Use(middleware.OperationRecord())
	learnResourcesChapterMergeRouterWithoutRecord := Router.Group("learnResourcesChapterMerge")
	var learnResourcesChapterMergeApi = v1.ApiGroupApp.BasicdataApiGroup.LearnResourcesChapterMergeApi
	{
		learnResourcesChapterMergeRouter.POST("createLearnResourcesChapterMerge", learnResourcesChapterMergeApi.CreateLearnResourcesChapterMerge)             // 新建LearnResourcesChapterMerge
		learnResourcesChapterMergeRouter.DELETE("deleteLearnResourcesChapterMerge", learnResourcesChapterMergeApi.DeleteLearnResourcesChapterMerge)           // 删除LearnResourcesChapterMerge
		learnResourcesChapterMergeRouter.DELETE("deleteLearnResourcesChapterMergeByIds", learnResourcesChapterMergeApi.DeleteLearnResourcesChapterMergeByIds) // 批量删除LearnResourcesChapterMerge
		learnResourcesChapterMergeRouter.PUT("updateLearnResourcesChapterMerge", learnResourcesChapterMergeApi.UpdateLearnResourcesChapterMerge)              // 更新LearnResourcesChapterMerge
	}
	{
		learnResourcesChapterMergeRouterWithoutRecord.GET("findLearnResourcesChapterMerge", learnResourcesChapterMergeApi.FindLearnResourcesChapterMerge)       // 根据ID获取LearnResourcesChapterMerge
		learnResourcesChapterMergeRouterWithoutRecord.GET("getLearnResourcesChapterMergeList", learnResourcesChapterMergeApi.GetLearnResourcesChapterMergeList) // 获取LearnResourcesChapterMerge列表
	}
}
