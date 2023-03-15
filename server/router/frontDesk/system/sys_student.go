package system

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
)

type BaseRouter struct {
}

func (c *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	//BaseRouter := Router.Group("common").Use(middleware.OperationRecord())
	CommonRouterWithoutRecord := Router.Group("common")
	var baseApi = api.ApiGroupApp.FrontDesk.SystemApiGroup.BaseApi
	{
		CommonRouterWithoutRecord.POST("studentLogin", baseApi.StudentLogin)
		CommonRouterWithoutRecord.GET("getTeachPlans", baseApi.GetTeachPlans)
		CommonRouterWithoutRecord.GET("getAllTeachPlans", baseApi.GetAllTeachPlans)
		CommonRouterWithoutRecord.POST("changePasswd", baseApi.StudentChangePassword)
		CommonRouterWithoutRecord.GET("getTargetTeachPlans", baseApi.GetTargetTeachPlans)
	}
	//{
	//	BaseRouter.GET("findLessons", baseApi.FindLessons)
	//}
}
