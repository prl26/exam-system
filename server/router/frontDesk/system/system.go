package system

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
)

type SystemRouter struct {
}

func (c *BaseRouter) InitSystemRouter(Router *gin.RouterGroup) {
	CommonRouterWithoutRecord := Router.Group("system")
	var systemApi = api.ApiGroupApp.FrontDesk.SystemApiGroup.SystemApi
	{
		CommonRouterWithoutRecord.GET("getTerms", systemApi.GetTerms)
		CommonRouterWithoutRecord.GET("getLessons", systemApi.GetLessons)
		CommonRouterWithoutRecord.POST("uploadImage", systemApi.UploadFile)
	}

}
