package teachplan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeachAttendanceRecordRouter struct {
}

// InitTeachAttendanceRecordRouter 初始化 TeachAttendanceRecord 路由信息
func (s *TeachAttendanceRecordRouter) InitTeachAttendanceRecordRouter(Router *gin.RouterGroup) {
	teachAttendanceRecordRouter := Router.Group("teachAttendanceRecord").Use(middleware.OperationRecord())
	teachAttendanceRecordRouterWithoutRecord := Router.Group("teachAttendanceRecord")
	var teachAttendanceRecordApi = v1.ApiGroupApp.TeachplanApiGroup.TeachAttendanceRecordApi
	{
		teachAttendanceRecordRouter.POST("createTeachAttendanceRecord", teachAttendanceRecordApi.CreateTeachAttendanceRecord)   // 新建TeachAttendanceRecord
		teachAttendanceRecordRouter.DELETE("deleteTeachAttendanceRecord", teachAttendanceRecordApi.DeleteTeachAttendanceRecord) // 删除TeachAttendanceRecord
		teachAttendanceRecordRouter.DELETE("deleteTeachAttendanceRecordByIds", teachAttendanceRecordApi.DeleteTeachAttendanceRecordByIds) // 批量删除TeachAttendanceRecord
		teachAttendanceRecordRouter.PUT("updateTeachAttendanceRecord", teachAttendanceRecordApi.UpdateTeachAttendanceRecord)    // 更新TeachAttendanceRecord
	}
	{
		teachAttendanceRecordRouterWithoutRecord.GET("findTeachAttendanceRecord", teachAttendanceRecordApi.FindTeachAttendanceRecord)        // 根据ID获取TeachAttendanceRecord
		teachAttendanceRecordRouterWithoutRecord.GET("getTeachAttendanceRecordList", teachAttendanceRecordApi.GetTeachAttendanceRecordList)  // 获取TeachAttendanceRecord列表
	}
}
