package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type StudentRouter struct {
}

// InitStudentRouter 初始化 Student 路由信息
func (s *StudentRouter) InitStudentRouter(Router *gin.RouterGroup) {
	studentRouter := Router.Group("student").Use(middleware.OperationRecord())
	studentRouterWithoutRecord := Router.Group("student")
	var studentApi = v1.ApiGroupApp.BasicdataApiGroup.StudentApi
	{
		studentRouter.POST("createStudent", studentApi.CreateStudent)             // 新建Student
		studentRouter.DELETE("deleteStudent", studentApi.DeleteStudent)           // 删除Student
		studentRouter.DELETE("deleteStudentByIds", studentApi.DeleteStudentByIds) // 批量删除Student
		studentRouter.PUT("updateStudent", studentApi.UpdateStudent)              // 更新Student
	}
	{
		studentRouterWithoutRecord.GET("findStudent", studentApi.FindStudent)       // 根据ID获取Student
		studentRouterWithoutRecord.GET("getStudentList", studentApi.GetStudentList) // 获取Student列表
	}
}
