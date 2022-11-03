package frontDesk

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
)

type CommonRouter struct {
}

func (c *CommonRouter) InitCommonRouter(Router *gin.RouterGroup) {
	CommonRouterWithoutRecord := Router.Group("common")
	var commonApi = api.ApiGroupApp.CommonApiGroup.CommonApi
	{
		CommonRouterWithoutRecord.GET("findLessons", commonApi.FindLessons)
	}
}
