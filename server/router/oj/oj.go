package oj

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 20:26

 * @Note:

 **/

type OjRouter struct {
}

func (s *OjRouter) InitFrontDeskOjRouter(Router *gin.RouterGroup) {
	judgesRouter := Router.Group("oj/judge")
	var judgeApi = v1.ApiGroupApp.OjApiGroup.JudgeApi
	{
		judgesRouter.POST("check", judgeApi.CheckJudge)
	}

	//multipleChoiceRouter := Router.Group("oj/multipleChoice").Use(middleware.OperationRecord())
	//multipleChoiceApi := v1.ApiGroupApp.OjApiGroup.MultipleChoiceApi
	//{
	//	multipleChoiceRouter.POST("check", multipleChoiceApi.CheckMultipleChoice)
	//}

	supplyBlankRouter := Router.Group("oj/supplyBlank").Use(middleware.OperationRecord())
	supplyBlankApi := v1.ApiGroupApp.OjApiGroup.SupplyBlankApi
	{
		supplyBlankRouter.POST("check", supplyBlankApi.CheckSupplyBlank)
	}

	programmRouter := Router.Group("oj/program")
	programmApi := v1.ApiGroupApp.OjApiGroup.ProgrammApi
	{
		programmRouter.POST("check", programmApi.CheckProgramm)
	}
}

func (s *OjRouter) InitBackgroundOjRouter(Router *gin.RouterGroup) {
	programmRouter := Router.Group("oj/program")
	programmApi := v1.ApiGroupApp.OjApiGroup.ProgrammApi
	{
		programmRouter.POST("compile", programmApi.Compile) //编译
		programmRouter.POST("execute", programmApi.Execute) //运行
	}

}
