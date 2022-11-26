package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type StudentRouter struct {
}

// InitStudentRouter 初始化 Student 路由信息
func (s *StudentRouter) InitStudentRouter(Router *gin.RouterGroup) {
	studentRouter := Router.Group("student").Use(middleware.OperationRecord())
	studentRouterWithoutRecord := Router.Group("student")

	var studentApi = api.ApiGroupApp.BackStage.BasicDataApiGroup.StudentApi
	{
		studentRouter.POST("excel", studentApi.AddStudentsByExcel)                  //表格添加Student
		studentRouter.POST("createStudent", studentApi.CreateStudent)               // 新建Student
		studentRouter.DELETE("deleteStudent", studentApi.DeleteStudent)             // 删除Student
		studentRouter.DELETE("deleteStudentByIds", studentApi.DeleteStudentByIds)   // 批量删除Student
		studentRouter.PUT("updateStudent", studentApi.UpdateStudent)                // 更新Student
		studentRouter.POST("resetStudentPassword", studentApi.ResetStudentPassword) // 重置学生密码
	}
	{
		studentRouter.StaticFile("excel", "./static/StudentInfo.xlsx")              //获取表格导入学生的示例文档(生产中可不放置后端)
		studentRouterWithoutRecord.GET("findStudent", studentApi.FindStudent)       // 根据ID获取Student
		studentRouterWithoutRecord.GET("getStudentList", studentApi.GetStudentList) // 获取Student列表
	}
}
