package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type DraftPaperRouter struct {
}

// InitDraftPaperRouter 初始化 ExamPaper 路由信息
func (s *DraftPaperRouter) InitDraftPaperRouter(Router *gin.RouterGroup) {
	DraftPaperRouter := Router.Group("paperDraft").Use(middleware.OperationRecord())
	var draftPaperApi = api.ApiGroupApp.BackStage.ExamManageApiGroup.DraftPaperApi
	{
		DraftPaperRouter.POST("createExamPaperDraft", draftPaperApi.CreatePaperDraft)       // 新建ExamPaper
		DraftPaperRouter.DELETE("deleteExamPaperDraft", draftPaperApi.DeleteExamPaperDraft) // 删除ExamPaper
		DraftPaperRouter.PUT("updateExamPaperDraft", draftPaperApi.UpdateExamPaperDraft)    // 更新ExamPaper
		DraftPaperRouter.GET("findExamPaperDraft", draftPaperApi.FindExamPaperDraft)        // 查找ExamPaper
		DraftPaperRouter.POST("convertDraftToPaper", draftPaperApi.ConvertDraftToPaper)     // 新建ExamPaper
		DraftPaperRouter.GET("getExamPaperDraftList", draftPaperApi.GetExamPaperDraftList)  // 新建ExamPaper
	}
}
