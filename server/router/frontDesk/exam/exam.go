package exam

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type ExamRouter struct {
}

func (c *ExamRouter) InitFrontExamRouter(Router *gin.RouterGroup) {
	FrontExamRouterWithoutRecord := Router.Group("frontExam").Use(middleware.OperationRecord())
	var frontExamApi = v1.ApiGroupApp.FrontDesk.ExamApiGroup.ExamApi
	{
		FrontExamRouterWithoutRecord.GET("findExamPlans", frontExamApi.FindExamPlans)
		FrontExamRouterWithoutRecord.GET("getExamPapers", frontExamApi.GetExamPaper)
		FrontExamRouterWithoutRecord.POST("commitExamPaper", frontExamApi.CommitExamPaper)
		FrontExamRouterWithoutRecord.GET("getExamScore", frontExamApi.GetExamScore)
		FrontExamRouterWithoutRecord.POST("commitProgram", frontExamApi.CommitProgram)
	}
}
