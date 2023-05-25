package questionBank

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
)

type QuestionBankRouter struct {
}

func (s *QuestionBankRouter) InitQuestionBankRouter(Router *gin.RouterGroup) {
	questionRouter := Router.Group("questionBank")
	var questionApi = v1.ApiGroupApp.FrontDesk.QuestionBankGroup.QuestionBankApi
	{
		questionRouter.GET("findQuestionsByChapterId", questionApi.FindQuestionsByChapterId)
		questionRouter.POST("beginPractice", questionApi.BeginPractice)
		questionRouter.POST("findHistory", questionApi.FindHistoryAnswer)
		questionRouter.POST("findAnswer", questionApi.FindAnswer)
	}
	target := Router.Group("target")
	targetApi := v1.ApiGroupApp.FrontDesk.QuestionBankGroup.TargetApi
	{
		target.POST("beginPractice", targetApi.BeginPractice)
		target.GET("findPractice", targetApi.FindTargetByKnowledgeId)
		target.GET("findPracticeDetail", targetApi.FindTargetDetail)
		target.POST("practiceGenerateInstance", targetApi.PracticeGenerateInstance)
		target.POST("practiceScore", targetApi.PracticeScore)
		target.GET("rankingList", targetApi.RankingList)
		target.GET("myRank", targetApi.MyRank)
	}
}
