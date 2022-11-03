package frontDesk

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
)

type FrontExamRouter struct {
}

func (c *CommonRouter) InitFrontExamRouter(Router *gin.RouterGroup) {
	FrontExamRouterWithoutRecord := Router.Group("frontExam")
	var frontExamApi = api.ApiGroupApp.FrontExamGroup.ExamApi
	{
		FrontExamRouterWithoutRecord.GET("findExamPlans", frontExamApi.FindExamPlans)
		FrontExamRouterWithoutRecord.GET("getExamPapers", frontExamApi.GetExamPaper)
		FrontExamRouterWithoutRecord.POST("commitExamPaper", frontExamApi.CommitExamPaper)
	}
}
