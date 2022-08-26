package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type TeachClassRouter struct {
}

// InitTeachClassRouter 初始化 TeachClass 路由信息
func (s *TeachClassRouter) InitTeachClassRouter(Router *gin.RouterGroup) {
	teachClassRouter := Router.Group("teachClass").Use(middleware.OperationRecord())
	teachClassRouterWithoutRecord := Router.Group("teachClass")
	var teachClassApi = v1.ApiGroupApp.BasicdataApiGroup.TeachClassApi
	{
		teachClassRouter.POST("createTeachClass", teachClassApi.CreateTeachClass)             // 新建TeachClass
		teachClassRouter.DELETE("deleteTeachClass", teachClassApi.DeleteTeachClass)           // 删除TeachClass
		teachClassRouter.DELETE("deleteTeachClassByIds", teachClassApi.DeleteTeachClassByIds) // 批量删除TeachClass
		teachClassRouter.PUT("updateTeachClass", teachClassApi.UpdateTeachClass)              // 更新TeachClass
	}
	{
		teachClassRouterWithoutRecord.GET("findTeachClass", teachClassApi.FindTeachClass)       // 根据ID获取TeachClass
		teachClassRouterWithoutRecord.GET("getTeachClassList", teachClassApi.GetTeachClassList) // 获取TeachClass列表
	}
}
