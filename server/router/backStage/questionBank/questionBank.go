package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/backStage/questionBank"
)

/*
*

  - @Author: AloneAtWar

  - @Date:   2022/8/26 18:56

  - @Note:

    *
*/
type QuestionBankRouter struct{}

func (s *QuestionBankRouter) InitQuestionBankRouter(Router *gin.RouterGroup) {
	//commonRouter := Router.Group("").Use(middleware.OperationRecord())
	//commonWithoutRecordRouter := Router.Group("common")
	//var backgroundCommonApi =  v1.ApiGroupApp.BackStage.QuestionBankApiGroup
	// 后台
	{
		Router.POST("uploadFile", questionBank.UploadFile)
		//commonWithoutRecordRouter.GET("findQuestionSupport", questionBankApi.FindQuestionSupport)
		//
		//commonRouter.POST("addCourseSupport", questionBankApi.AddCourseSupport)
		//commonRouter.DELETE("deleteCourseSupport", questionBankApi.DeleteCourseSupport)
	}
	// 前台
	{
		//commonWithoutRecordRouter.GET("findQuestions", frontDeskCommonApi.FindQuestionsByKnowledgeId)

	}
}
