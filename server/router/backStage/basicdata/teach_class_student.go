package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type TeachClassStudentRouter struct {
}

// InitTeachClassStudentRouter 初始化 TeachClassStudent 路由信息
func (s *TeachClassStudentRouter) InitTeachClassStudentRouter(Router *gin.RouterGroup) {
	teachClassStudentRouter := Router.Group("teachClassStudent").Use(middleware.OperationRecord())
	teachClassStudentRouterWithoutRecord := Router.Group("teachClassStudent")
	//var teachClassStudentApi = Backstage.ApiGroupApp.BasicdataApiGroup.TeachClassStudentApi
	var multiTableApi = api.ApiGroupApp.BackStage.BasicDataApiGroup.MultiTableApi
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
