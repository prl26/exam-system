package teachplan

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type ExamPlanRouter struct {
}

// InitExamPlanRouter 初始化 ExamPlan 路由信息
func (s *ExamPlanRouter) InitExamPlanRouter(Router *gin.RouterGroup) {
	examPlanRouter := Router.Group("examPlan").Use(middleware.OperationRecord())
	examPlanRouterWithoutRecord := Router.Group("examPlan")
	var examPlanApi = api.ApiGroupApp.BackStage.TeachPlanApiGroup.ExamPlanApi
	{
		examPlanRouter.POST("createExamPlan", examPlanApi.CreateExamPlan)             // 新建ExamPlan
		examPlanRouter.DELETE("deleteExamPlan", examPlanApi.DeleteExamPlan)           // 删除ExamPlan
		examPlanRouter.DELETE("deleteExamPlanByIds", examPlanApi.DeleteExamPlanByIds) // 批量删除ExamPlan
		examPlanRouter.PUT("updateExamPlan", examPlanApi.UpdateExamPlan)              // 更新ExamPlan
	}
	{
		examPlanRouterWithoutRecord.GET("findExamPlan", examPlanApi.FindExamPlanById)   // 根据ID获取ExamPlan
		examPlanRouterWithoutRecord.GET("getExamPlanList", examPlanApi.GetExamPlanList) // 获取ExamPlan列表
	}
}
