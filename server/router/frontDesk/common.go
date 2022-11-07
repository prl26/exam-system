package frontDesk

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
)

type CommonRouter struct {
}

func (c *CommonRouter) InitCommonRouter(Router *gin.RouterGroup) {
	//CommonRouter := Router.Group("common").Use(middleware.OperationRecord())
	CommonRouterWithoutRecord := Router.Group("common")
	var commonApi = api.ApiGroupApp.CommonApiGroup.CommonApi
	{
		CommonRouterWithoutRecord.POST("studentLogin", commonApi.StudentLogin)
	}
	//{
	//	CommonRouter.GET("findLessons", commonApi.FindLessons)
	//}
}
