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
		examstudentPaperRouter.POST("recoverPower", examstudentPaperApi.RecoverPower)                                 //恢复考试资格
		examstudentPaperRouter.POST("reportScore", examstudentPaperApi.ReportScore)                                   //上报成绩
		examstudentPaperRouter.POST("paperCheating", examstudentPaperApi.PaperCheating)                               //答案修正
		examstudentPaperRouter.GET("paperReview", examstudentPaperApi.PaperReview)                                    //分页查看学生成绩-试卷批阅
		examstudentPaperRouter.POST("statusMonitor", examstudentPaperApi.StatusMonitor)                               //状态检测
	}
	{
		examstudentPaperRouterWithoutRecord.GET("findExamStudentPaper", examstudentPaperApi.FindExamStudentPaper)       // 根据ID获取ExamStudentPaper
		examstudentPaperRouterWithoutRecord.GET("getExamStudentPaperList", examstudentPaperApi.GetExamStudentPaperList) // 获取ExamStudentPaper列表
	}
}
