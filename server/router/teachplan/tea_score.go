package teachplan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScoreRouter struct {
}

// InitScoreRouter 初始化 Score 路由信息
func (s *ScoreRouter) InitScoreRouter(Router *gin.RouterGroup) {
	scoreRouter := Router.Group("score").Use(middleware.OperationRecord())
	scoreRouterWithoutRecord := Router.Group("score")
	var scoreApi = v1.ApiGroupApp.TeachplanApiGroup.ScoreApi
	{
		scoreRouter.POST("createScore", scoreApi.CreateScore)   // 新建Score
		scoreRouter.DELETE("deleteScore", scoreApi.DeleteScore) // 删除Score
		scoreRouter.DELETE("deleteScoreByIds", scoreApi.DeleteScoreByIds) // 批量删除Score
		scoreRouter.PUT("updateScore", scoreApi.UpdateScore)    // 更新Score
	}
	{
		scoreRouterWithoutRecord.GET("findScore", scoreApi.FindScore)        // 根据ID获取Score
		scoreRouterWithoutRecord.GET("getScoreList", scoreApi.GetScoreList)  // 获取Score列表
	}
}
