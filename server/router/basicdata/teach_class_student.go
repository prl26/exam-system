package basicdata

import (
	"exam-system/api/v1"
	"exam-system/middleware"
	"github.com/gin-gonic/gin"
)

type TeachClassStudentRouter struct {
}

// InitTeachClassStudentRouter 初始化 TeachClassStudent 路由信息
func (s *TeachClassStudentRouter) InitTeachClassStudentRouter(Router *gin.RouterGroup) {
	teachClassStudentRouter := Router.Group("teachClassStudent").Use(middleware.OperationRecord())
	teachClassStudentRouterWithoutRecord := Router.Group("teachClassStudent")
	//var teachClassStudentApi = v1.ApiGroupApp.BasicdataApiGroup.TeachClassStudentApi
	var multiTableApi = v1.ApiGroupApp.BasicdataApiGroup.MultiTableApi
	{

		teachClassStudentRouter.POST("initTeachClassStudent", multiTableApi.InitTeachClassStudent)     // 教学班中添加学生
		teachClassStudentRouter.POST("deleteTeachClassStudent", multiTableApi.DeleteTeachClassStudent) // 教学班中移除学生
		teachClassStudentRouter.POST("addStudentByClass", multiTableApi.AddStudentByClass)             // 教学班中添加一个班的学生
		teachClassStudentRouter.POST("deleteStudentByClass", multiTableApi.DeleteStudentByClass)       // 教学班中添加一个班的学生
	}
	{
		teachClassStudentRouterWithoutRecord.GET("getTeachClassStudentList", multiTableApi.GetTeachClassStudentList) // 获取TeachClassStudent列表
	}
}
