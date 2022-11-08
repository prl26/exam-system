package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 18:56

 * @Note:

 **/
type QuestionBankRouter struct{}

func (s *QuestionBankRouter) InitQuestionBankRouter(Router *gin.RouterGroup) {
	commonRouter := Router.Group("common").Use(middleware.OperationRecord())
	commonWithoutRecordRouter := Router.Group("common")
	var backgroundCommonApi =  v1.ApiGroupApp.BackStage.QuestionBankApiGroup
	questionBankApi := backgroundCommonApi.CommonApi
	// 后台
	{
		commonWithoutRecordRouter.GET("findQuestionSupport", questionBankApi.FindQuestionSupport)

		commonRouter.POST("addCourseSupport", questionBankApi.AddCourseSupport)
		commonRouter.DELETE("deleteCourseSupport", questionBankApi.DeleteCourseSupport)
	}
	//// 前台
	//{
	//	commonWithoutRecordRouter.GET("findQuestions", frontDeskCommonApi.FindQuestionsByKnowledgeId)
	//
	//}
}
