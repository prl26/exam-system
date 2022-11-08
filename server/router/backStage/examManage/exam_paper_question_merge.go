package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type PaperQuestionMergeRouter struct {
}

// InitPaperQuestionMergeRouter 初始化 PaperQuestionMerge 路由信息
func (s *PaperQuestionMergeRouter) InitPaperQuestionMergeRouter(Router *gin.RouterGroup) {
	paperQuestionMergeRouter := Router.Group("paperQuestionMerge").Use(middleware.OperationRecord())
	paperQuestionMergeRouterWithoutRecord := Router.Group("paperQuestionMerge")
	var paperQuestionMergeApi = api.ApiGroupApp.BackStage.ExamManageApiGroup.PaperQuestionMergeApi
	{
		paperQuestionMergeRouter.POST("createPaperQuestionMerge", paperQuestionMergeApi.CreatePaperQuestionMerge)             // 新建PaperQuestionMerge
		paperQuestionMergeRouter.DELETE("deletePaperQuestionMerge", paperQuestionMergeApi.DeletePaperQuestionMerge)           // 删除PaperQuestionMerge
		paperQuestionMergeRouter.DELETE("deletePaperQuestionMergeByIds", paperQuestionMergeApi.DeletePaperQuestionMergeByIds) // 批量删除PaperQuestionMerge
		paperQuestionMergeRouter.PUT("updatePaperQuestionMerge", paperQuestionMergeApi.UpdatePaperQuestionMerge)              // 更新PaperQuestionMerge
	}
	{
		paperQuestionMergeRouterWithoutRecord.GET("findPaperQuestionMerge", paperQuestionMergeApi.FindPaperQuestionMerge)       // 根据ID获取PaperQuestionMerge
		paperQuestionMergeRouterWithoutRecord.GET("getPaperQuestionMergeList", paperQuestionMergeApi.GetPaperQuestionMergeList) // 获取PaperQuestionMerge列表
	}
}
