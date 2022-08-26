package lessondata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type ArticleResourcesRouter struct {
}

// InitArticleResourcesRouter 初始化 ArticleResources 路由信息
func (s *ArticleResourcesRouter) InitArticleResourcesRouter(Router *gin.RouterGroup) {
	articleResourcesRouter := Router.Group("articleResources").Use(middleware.OperationRecord())
	articleResourcesRouterWithoutRecord := Router.Group("articleResources")
	var articleResourcesApi = v1.ApiGroupApp.LessondataApiGroup.ArticleResourcesApi
	{
		articleResourcesRouter.POST("createArticleResources", articleResourcesApi.CreateArticleResources)             // 新建ArticleResources
		articleResourcesRouter.DELETE("deleteArticleResources", articleResourcesApi.DeleteArticleResources)           // 删除ArticleResources
		articleResourcesRouter.DELETE("deleteArticleResourcesByIds", articleResourcesApi.DeleteArticleResourcesByIds) // 批量删除ArticleResources
		articleResourcesRouter.PUT("updateArticleResources", articleResourcesApi.UpdateArticleResources)              // 更新ArticleResources
	}
	{
		articleResourcesRouterWithoutRecord.GET("findArticleResources", articleResourcesApi.FindArticleResources)       // 根据ID获取ArticleResources
		articleResourcesRouterWithoutRecord.GET("getArticleResourcesList", articleResourcesApi.GetArticleResourcesList) // 获取ArticleResources列表
	}
}
