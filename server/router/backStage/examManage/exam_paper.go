package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type ExamPaperRouter struct {
}

// InitExamPaperRouter 初始化 ExamPaper 路由信息
func (s *ExamPaperRouter) InitExamPaperRouter(Router *gin.RouterGroup) {
	examPaperRouter := Router.Group("examPaper").Use(middleware.OperationRecord())
	examPaperRouterWithoutRecord := Router.Group("examPaper")
	var examPaperApi = api.ApiGroupApp.BackStage.ExamManageApiGroup.ExamPaperApi
	{
		examPaperRouter.POST("createExamPaper", examPaperApi.CreateExamPaperByRand)       // 新建ExamPaper
		examPaperRouter.DELETE("deleteExamPaper", examPaperApi.DeleteExamPaper)           // 删除ExamPaper
		examPaperRouter.DELETE("deleteExamPaperByIds", examPaperApi.DeleteExamPaperByIds) // 批量删除ExamPaper
		examPaperRouter.PUT("updateExamPaper", examPaperApi.UpdateExamPaper)              // 更新ExamPaper
		//examPaperRouter.GET("setStudentsToRedis", examPaperApi.SetStudentsToRedis)
	}
	{
		examPaperRouterWithoutRecord.GET("findExamPaper", examPaperApi.FindExamPaper)          // 根据ID获取ExamPaper
		examPaperRouterWithoutRecord.GET("getExamPaperList", examPaperApi.GetExamPaperList)    // 获取ExamPaper列表
		examPaperRouterWithoutRecord.GET("paperDistribution", examPaperApi.PaperDistribution)  //试卷分发
		examPaperRouterWithoutRecord.POST("exportPaper", examPaperApi.ExportPaper)             //导出成绩表
		examPaperRouterWithoutRecord.POST("exportMultiPaper", examPaperApi.ExportMultiPaper)   //导出成绩表(教学计划)
		examPaperRouterWithoutRecord.POST("exportPaperToHtml", examPaperApi.ExportPaperToHtml) //导出成绩表
		examPaperRouterWithoutRecord.POST("setExamPre", examPaperApi.SetExamPre)               //进入考试准备阶段
	}
}
