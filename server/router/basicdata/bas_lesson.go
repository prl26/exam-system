package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type LessonRouter struct {
}

// InitLessonRouter 初始化 Lesson 路由信息
func (s *LessonRouter) InitLessonRouter(Router *gin.RouterGroup) {
	lessonRouter := Router.Group("lesson").Use(middleware.OperationRecord())
	lessonRouterWithoutRecord := Router.Group("lesson")
	var lessonApi = api.ApiGroupApp.BasicdataApiGroup.LessonApi
	{
		lessonRouter.POST("createLesson", lessonApi.CreateLesson)             // 新建Lesson
		lessonRouter.DELETE("deleteLesson", lessonApi.DeleteLesson)           // 删除Lesson
		lessonRouter.DELETE("deleteLessonByIds", lessonApi.DeleteLessonByIds) // 批量删除Lesson
		lessonRouter.PUT("updateLesson", lessonApi.UpdateLesson)              // 更新Lesson
	}
	{
		lessonRouterWithoutRecord.GET("findLesson", lessonApi.FindLesson)       // 根据ID获取Lesson
		lessonRouterWithoutRecord.GET("getLessonList", lessonApi.GetLessonList) // 获取Lesson列表
	}
}
