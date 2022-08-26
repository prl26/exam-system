package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type CollegeRouter struct {
}

// InitCollegeRouter 初始化 College 路由信息
func (s *CollegeRouter) InitCollegeRouter(Router *gin.RouterGroup) {
	collegeRouter := Router.Group("college").Use(middleware.OperationRecord())
	collegeRouterWithoutRecord := Router.Group("college")
	var collegeApi = v1.ApiGroupApp.BasicdataApiGroup.CollegeApi
	{
		collegeRouter.POST("createCollege", collegeApi.CreateCollege)             // 新建College
		collegeRouter.DELETE("deleteCollege", collegeApi.DeleteCollege)           // 删除College
		collegeRouter.DELETE("deleteCollegeByIds", collegeApi.DeleteCollegeByIds) // 批量删除College
		collegeRouter.PUT("updateCollege", collegeApi.UpdateCollege)              // 更新College
	}
	{
		collegeRouterWithoutRecord.GET("findCollege", collegeApi.FindCollege)       // 根据ID获取College
		collegeRouterWithoutRecord.GET("getCollegeList", collegeApi.GetCollegeList) // 获取College列表
	}
}
