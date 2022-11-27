package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type PublicProgramRouter struct{}

func (*PublicProgramRouter) InitPublicProgram(Router *gin.RouterGroup) {
	var publicProgramApi = v1.ApiGroupApp.BackStage.QuestionBankApiGroup.PublicProgramApi
	programRouter := Router.Group("publicProgram").Use(middleware.OperationRecord())
	{
		programRouter.POST("create", publicProgramApi.Create)
		programRouter.PUT("update", publicProgramApi.Update)
	}
	programWithoutRecordRouter := Router.Group("publicProgram")
	{
		programWithoutRecordRouter.GET("findList", publicProgramApi.FindList)
		programWithoutRecordRouter.GET("findDetail", publicProgramApi.FindDetail)
	}

	//{
	//	programRouter.POST("create", programApi.Create)
	//	programRouter.PUT("editCase", programApi.EditProgramCases)
	//	programRouter.PUT("editDetail", programApi.EditProgramDetail)
	//	programRouter.DELETE("deleteCase", programApi.DeleteProgramCases)
	//	programRouter.DELETE("delete", programApi.DeleteProgramm)
	//	programRouter.POST("addCase", programApi.AddProgramCase)
	//	programRouter.POST("addLanguageSupport", programApi.AddLanguageSupport)
	//	programRouter.PUT("editLanguageSupport", programApi.EditLanguageSupport)
	//	programRouter.DELETE("deleteLanguageSupport", programApi.DeleteLanguageSupport)
	//}
	//programWithoutRecordRouter := Router.Group("program")
	//{
	//	programWithoutRecordRouter.GET("findDetail", programApi.FindDetail)
	//	programWithoutRecordRouter.GET("findList", programApi.FindList)
	//	programWithoutRecordRouter.GET("findCase", programApi.FindProgramCases)
	//}
}
