package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeachClassStudentRouter struct {
}

// InitTeachClassStudentRouter 初始化 TeachClassStudent 路由信息
func (s *TeachClassStudentRouter) InitTeachClassStudentRouter(Router *gin.RouterGroup) {
	teachClassStudentRouter := Router.Group("teachClassStudent").Use(middleware.OperationRecord())
	teachClassStudentRouterWithoutRecord := Router.Group("teachClassStudent")
	var teachClassStudentApi = v1.ApiGroupApp.BasicdataApiGroup.TeachClassStudentApi
	{
		teachClassStudentRouter.POST("createTeachClassStudent", teachClassStudentApi.CreateTeachClassStudent)   // 新建TeachClassStudent
		teachClassStudentRouter.DELETE("deleteTeachClassStudent", teachClassStudentApi.DeleteTeachClassStudent) // 删除TeachClassStudent
		teachClassStudentRouter.DELETE("deleteTeachClassStudentByIds", teachClassStudentApi.DeleteTeachClassStudentByIds) // 批量删除TeachClassStudent
		teachClassStudentRouter.PUT("updateTeachClassStudent", teachClassStudentApi.UpdateTeachClassStudent)    // 更新TeachClassStudent
	}
	{
		teachClassStudentRouterWithoutRecord.GET("findTeachClassStudent", teachClassStudentApi.FindTeachClassStudent)        // 根据ID获取TeachClassStudent
		teachClassStudentRouterWithoutRecord.GET("getTeachClassStudentList", teachClassStudentApi.GetTeachClassStudentList)  // 获取TeachClassStudent列表
	}
}
