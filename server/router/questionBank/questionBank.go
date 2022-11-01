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
	var backgroundCommonApi = v1.ApiGroupApp.QuestionBankApiGroup.Background.CommonApi
	var frontDeskCommonApi = v1.ApiGroupApp.QuestionBankApiGroup.FrontDesk.CommonApi
	// 后台
	{
		commonWithoutRecordRouter.GET("findQuestionSupport", backgroundCommonApi.FindQuestionSupport)

		commonRouter.POST("addCourseSupport", backgroundCommonApi.AddCourseSupport)
		commonRouter.DELETE("deleteCourseSupport", backgroundCommonApi.DeleteCourseSupport)
	}
	// 前台
	{
		commonWithoutRecordRouter.GET("findQuestions", frontDeskCommonApi.FindQuestionsByChapterId)

	}
}
