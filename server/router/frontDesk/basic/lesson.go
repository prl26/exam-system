package basic

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
)

type LessonRouter struct {
}

func (c *LessonRouter) InitLessonRouter(Router *gin.RouterGroup) {
	lessonRouter := Router.Group("lesson")
	var lessonApi = v1.ApiGroupApp.FrontDesk.BasicApiGroup.LessonApi
	{
		lessonRouter.GET("findChapter", lessonApi.FindLessonDetail)
		lessonRouter.GET("findKnowledge", lessonApi.FindKnowledge)
	}
}
