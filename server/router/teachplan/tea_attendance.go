package teachplan

import (
	"exam-system/api/v1"
	"exam-system/middleware"
	"github.com/gin-gonic/gin"
)

type TeachAttendanceRouter struct {
}

// InitTeachAttendanceRouter 初始化 TeachAttendance 路由信息
func (s *TeachAttendanceRouter) InitTeachAttendanceRouter(Router *gin.RouterGroup) {
	teachAttendanceRouter := Router.Group("teachAttendance").Use(middleware.OperationRecord())
	teachAttendanceRouterWithoutRecord := Router.Group("teachAttendance")
	var teachAttendanceApi = v1.ApiGroupApp.TeachplanApiGroup.TeachAttendanceApi
	{
		teachAttendanceRouter.POST("createTeachAttendance", teachAttendanceApi.CreateTeachAttendance)             // 新建TeachAttendance
		teachAttendanceRouter.DELETE("deleteTeachAttendance", teachAttendanceApi.DeleteTeachAttendance)           // 删除TeachAttendance
		teachAttendanceRouter.DELETE("deleteTeachAttendanceByIds", teachAttendanceApi.DeleteTeachAttendanceByIds) // 批量删除TeachAttendance
		teachAttendanceRouter.PUT("updateTeachAttendance", teachAttendanceApi.UpdateTeachAttendance)              // 更新TeachAttendance
	}
	{
		teachAttendanceRouterWithoutRecord.GET("findTeachAttendance", teachAttendanceApi.FindTeachAttendance)       // 根据ID获取TeachAttendance
		teachAttendanceRouterWithoutRecord.GET("getTeachAttendanceList", teachAttendanceApi.GetTeachAttendanceList) // 获取TeachAttendance列表
	}
}
