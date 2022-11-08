package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type ExamStudentPaperRouter struct {
}

// InitExamStudentPaperRouter 初始化 ExamStudentPaper 路由信息
func (s *ExamStudentPaperRouter) InitExamStudentPaperRouter(Router *gin.RouterGroup) {
	examstudentPaperRouter := Router.Group("examstudentPaper").Use(middleware.OperationRecord())
	examstudentPaperRouterWithoutRecord := Router.Group("examstudentPaper")
	var examstudentPaperApi = api.ApiGroupApp.BackStage.ExamManageApiGroup.ExamStudentPaperApi
	{
		examstudentPaperRouter.POST("createExamStudentPaper", examstudentPaperApi.CreateExamStudentPaper)             // 新建ExamStudentPaper
		examstudentPaperRouter.DELETE("deleteExamStudentPaper", examstudentPaperApi.DeleteExamStudentPaper)           // 删除ExamStudentPaper
		examstudentPaperRouter.DELETE("deleteExamStudentPaperByIds", examstudentPaperApi.DeleteExamStudentPaperByIds) // 批量删除ExamStudentPaper
		examstudentPaperRouter.PUT("updateExamStudentPaper", examstudentPaperApi.UpdateExamStudentPaper)              // 更新ExamStudentPaper
	}
	{
		examstudentPaperRouterWithoutRecord.GET("findExamStudentPaper", examstudentPaperApi.FindExamStudentPaper)       // 根据ID获取ExamStudentPaper
		examstudentPaperRouterWithoutRecord.GET("getExamStudentPaperList", examstudentPaperApi.GetExamStudentPaperList) // 获取ExamStudentPaper列表
	}
}
