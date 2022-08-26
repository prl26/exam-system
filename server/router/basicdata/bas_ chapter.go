package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type ChapterRouter struct {
}

// InitChapterRouter 初始化 Chapter 路由信息
func (s *ChapterRouter) InitChapterRouter(Router *gin.RouterGroup) {
	chapterRouter := Router.Group("chapter").Use(middleware.OperationRecord())
	chapterRouterWithoutRecord := Router.Group("chapter")
	var chapterApi = v1.ApiGroupApp.BasicdataApiGroup.ChapterApi
	{
		chapterRouter.POST("createChapter", chapterApi.CreateChapter)             // 新建Chapter
		chapterRouter.DELETE("deleteChapter", chapterApi.DeleteChapter)           // 删除Chapter
		chapterRouter.DELETE("deleteChapterByIds", chapterApi.DeleteChapterByIds) // 批量删除Chapter
		chapterRouter.PUT("updateChapter", chapterApi.UpdateChapter)              // 更新Chapter
	}
	{
		chapterRouterWithoutRecord.GET("findChapter", chapterApi.FindChapter)       // 根据ID获取Chapter
		chapterRouterWithoutRecord.GET("getChapterList", chapterApi.GetChapterList) // 获取Chapter列表
	}
}
