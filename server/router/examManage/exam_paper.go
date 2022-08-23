package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExamPaperRouter struct {
}

// InitExamPaperRouter 初始化 ExamPaper 路由信息
func (s *ExamPaperRouter) InitExamPaperRouter(Router *gin.RouterGroup) {
	examPaperRouter := Router.Group("examPaper").Use(middleware.OperationRecord())
	examPaperRouterWithoutRecord := Router.Group("examPaper")
	var examPaperApi = v1.ApiGroupApp.ExammanageApiGroup.ExamPaperApi
	{
		examPaperRouter.POST("createExamPaper", examPaperApi.CreateExamPaper)   // 新建ExamPaper
		examPaperRouter.DELETE("deleteExamPaper", examPaperApi.DeleteExamPaper) // 删除ExamPaper
		examPaperRouter.DELETE("deleteExamPaperByIds", examPaperApi.DeleteExamPaperByIds) // 批量删除ExamPaper
		examPaperRouter.PUT("updateExamPaper", examPaperApi.UpdateExamPaper)    // 更新ExamPaper
	}
	{
		examPaperRouterWithoutRecord.GET("findExamPaper", examPaperApi.FindExamPaper)        // 根据ID获取ExamPaper
		examPaperRouterWithoutRecord.GET("getExamPaperList", examPaperApi.GetExamPaperList)  // 获取ExamPaper列表
	}
}