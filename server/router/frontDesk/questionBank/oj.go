package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type OjRouter struct {
}

func (s *OjRouter) InitOjRouter(Router *gin.RouterGroup) {
	judgesRouter := Router.Group("oj/judge")
	var judgeApi = v1.ApiGroupApp.FrontDesk.QuestionBankGroup.OjApi
	{
		judgesRouter.POST("check", judgeApi.CheckJudge)
	}

	multipleChoiceRouter := Router.Group("oj/multipleChoice").Use(middleware.OperationRecord())
	multipleChoiceApi := v1.ApiGroupApp.FrontDesk.QuestionBankGroup.OjApi
	{
		multipleChoiceRouter.POST("check", multipleChoiceApi.CheckMultipleChoice)
	}

	supplyBlankRouter := Router.Group("oj/supplyBlank").Use(middleware.OperationRecord())
	supplyBlankApi := v1.ApiGroupApp.FrontDesk.QuestionBankGroup.OjApi
	{
		supplyBlankRouter.POST("check", supplyBlankApi.CheckSupplyBlank)
	}

	//programmRouter := Router.Group("oj/program")
	//programmApi :=	v1.ApiGroupApp.FrontDesk.QuestionBankGroup.OjApi
	//{
	//	programmRouter.POST("check", programmApi.CheckProgramm)
	//}
}
