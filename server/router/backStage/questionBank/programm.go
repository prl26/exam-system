package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 20:23

 * @Note:

 **/
type ProgramRouter struct{}

func (*ProgramRouter) InitProgram(Router *gin.RouterGroup) {
	var programApi = v1.ApiGroupApp.BackStage.QuestionBankApiGroup.ProgramApi
	programRouter := Router.Group("program").Use(middleware.OperationRecord())
	{
		programRouter.POST("create", programApi.Create)
		programRouter.PUT("editCase", programApi.EditProgramCases)
		programRouter.PUT("editDetail", programApi.EditProgramDetail)
		programRouter.DELETE("deleteCase", programApi.DeleteProgramCases)
		programRouter.DELETE("delete", programApi.DeleteProgramm)
		programRouter.POST("addCase", programApi.AddProgramCase)
		programRouter.POST("addLanguageSupport", programApi.AddLanguageSupport)
		programRouter.PUT("editLanguageSupport", programApi.EditLanguageSupport)
		programRouter.DELETE("deleteLanguageSupport", programApi.DeleteLanguageSupport)
	}
	programWithoutRecordRouter := Router.Group("program")
	{
		programWithoutRecordRouter.GET("findDetail", programApi.FindDetail)
		programWithoutRecordRouter.GET("findList", programApi.FindList)
		programWithoutRecordRouter.GET("findCase", programApi.FindProgramCases)
	}
}
