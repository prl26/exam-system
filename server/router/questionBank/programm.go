package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 20:23

 * @Note:

 **/
type ProgramRouter struct{}

func (*ProgramRouter) InitProgram(Router *gin.RouterGroup) {
	baseRouter := Router.Group("program").Use(middleware.OperationRecord())
	var programApi = v1.ApiGroupApp.QuestionBankApiGroup.ProgramApi
	{
		baseRouter.GET("findDetail", programApi.FindDetail)
		baseRouter.PUT("editDetail", programApi.EditProgramDetail)
		baseRouter.DELETE("delete", programApi.DeleteProgramm)
		baseRouter.GET("findCase", programApi.FindProgrammCases)
		baseRouter.POST("addCase", programApi.AddProgrammCase)
		baseRouter.PUT("editCase", programApi.EditProgrammCases)
		baseRouter.DELETE("deleteCase", programApi.DeleteProgrammCases)
		baseRouter.POST("addLanguageSupport", programApi.AddLanguageSupport)
		baseRouter.PUT("editLanguageSupport", programApi.EditLanguageSupport)
		baseRouter.DELETE("deleteLanguageSupport", programApi.DeleteLanguageSupport)
	}
}
