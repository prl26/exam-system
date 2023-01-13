package exam

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type TargetExamRouter struct {
}

func (c *ExamRouter) InitTargetExamRouter(Router *gin.RouterGroup) {
	FrontExamRouterWithoutRecord := Router.Group("targetExam").Use(middleware.OperationRecord())
	var targetExamApi = v1.ApiGroupApp.FrontDesk.ExamApiGroup.TargetExamApi
	{
		FrontExamRouterWithoutRecord.GET("getTargetExamPapers", targetExamApi.GetTargetExamPaper)
		FrontExamRouterWithoutRecord.POST("commitTargetExamPaper", targetExamApi.CommitTargetExamPaper)
		FrontExamRouterWithoutRecord.POST("examGenerateInstance", targetExamApi.ExamGenerateInstance)
		FrontExamRouterWithoutRecord.GET("getTargetExamScore", targetExamApi.GetTargetExamScore)
	}
}
