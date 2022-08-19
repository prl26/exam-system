package teachplan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExamPlanRouter struct {
}

// InitExamPlanRouter 初始化 ExamPlan 路由信息
func (s *ExamPlanRouter) InitExamPlanRouter(Router *gin.RouterGroup) {
	examPlanRouter := Router.Group("examPlan").Use(middleware.OperationRecord())
	examPlanRouterWithoutRecord := Router.Group("examPlan")
	var examPlanApi = v1.ApiGroupApp.TeachplanApiGroup.ExamPlanApi
	{
		examPlanRouter.POST("createExamPlan", examPlanApi.CreateExamPlan)   // 新建ExamPlan
		examPlanRouter.DELETE("deleteExamPlan", examPlanApi.DeleteExamPlan) // 删除ExamPlan
		examPlanRouter.DELETE("deleteExamPlanByIds", examPlanApi.DeleteExamPlanByIds) // 批量删除ExamPlan
		examPlanRouter.PUT("updateExamPlan", examPlanApi.UpdateExamPlan)    // 更新ExamPlan
	}
	{
		examPlanRouterWithoutRecord.GET("findExamPlan", examPlanApi.FindExamPlan)        // 根据ID获取ExamPlan
		examPlanRouterWithoutRecord.GET("getExamPlanList", examPlanApi.GetExamPlanList)  // 获取ExamPlan列表
	}
}
