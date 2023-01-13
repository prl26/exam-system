package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 14:56

 * @Note:

 **/

type RouterGroup struct {
	QuestionBankMultipleChoiceRouter
	QuestionBankJudgeRouter
	QuestionBankSupplyBlankRouter
	QuestionBankRouter
	ProgramRouter
	PublicProgramRouter
	OjRouter
	TargetRouter
	SituationRouter
}

type SituationRouter struct {
}

func (g SituationRouter) InitSituation(Router *gin.RouterGroup) {
	var publicProgramApi = v1.ApiGroupApp.BackStage.QuestionBankApiGroup.SituationApi
	programWithoutRecordRouter := Router.Group("situation")
	{
		programWithoutRecordRouter.GET("byTeachClass", publicProgramApi.FindTeachClassSituation)
		programWithoutRecordRouter.GET("byStudent", publicProgramApi.FindStudentSituation)
		programWithoutRecordRouter.GET("findDetail", publicProgramApi.FindDetail)
	}

}
