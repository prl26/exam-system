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
		examstudentPaperRouter.POST("reportStudentScore", examstudentPaperApi.ReportStudentScore)                     //单独上报学生成绩
		examstudentPaperRouter.POST("paperCheating", examstudentPaperApi.PaperCheating)                               //答案修正
		examstudentPaperRouter.GET("paperReview", examstudentPaperApi.PaperReview)                                    //分页查看学生成绩-试卷批阅
		examstudentPaperRouter.GET("PaperMultiReview", examstudentPaperApi.PaperMultiReview)                          //教学计划下,成绩查看
		examstudentPaperRouter.GET("statusMonitor", examstudentPaperApi.StatusMonitor)                                //状态检测
		examstudentPaperRouter.POST("execAgain", examstudentPaperApi.ExecAgain)                                       //单学生重新批阅
		examstudentPaperRouter.POST("allExecAgain", examstudentPaperApi.AllExecAgain)                                 //考试计划下所有学生试卷重批阅
		examstudentPaperRouter.GET("getCommitRecord", examstudentPaperApi.GetCommitRecord)                            //获取提交日志记录
		examstudentPaperRouter.POST("recoverByRecord", examstudentPaperApi.RecoverByRecord)                           //恢复学生答卷
		examstudentPaperRouter.POST("deleteStudentAnswer", examstudentPaperApi.DeleteStudentAnswer)                   //删除学生答卷
		examstudentPaperRouter.POST("forceCommitStudent", examstudentPaperApi.ForceCommitStudent)                     //强制提交
		examstudentPaperRouter.GET("getDistribution", examstudentPaperApi.GetDistribution)                            // 查看分发情况

	}
	{
		examstudentPaperRouterWithoutRecord.GET("findExamStudentPaper", examstudentPaperApi.FindExamStudentPaper)       // 根据ID获取ExamStudentPaper
		examstudentPaperRouterWithoutRecord.GET("getExamStudentPaperList", examstudentPaperApi.GetExamStudentPaperList) // 获取ExamStudentPaper列表
	}
}
