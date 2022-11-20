package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type PaperTemplateItemRouter struct {
}

// InitPaperTemplateItemRouter 初始化 PaperTemplateItem 路由信息
func (s *PaperTemplateItemRouter) InitPaperTemplateItemRouter(Router *gin.RouterGroup) {
	paperTemplateItemRouter := Router.Group("paperTemplateItem").Use(middleware.OperationRecord())
	paperTemplateItemRouterWithoutRecord := Router.Group("paperTemplateItem")
	var paperTemplateItemApi = api.ApiGroupApp.BackStage.ExamManageApiGroup.PaperTemplateItemApi
	{
		paperTemplateItemRouter.POST("createPaperTemplateItem", paperTemplateItemApi.CreatePaperTemplateItem)             // 新建PaperTemplateItem
		paperTemplateItemRouter.DELETE("deletePaperTemplateItem", paperTemplateItemApi.DeletePaperTemplateItem)           // 删除PaperTemplateItem
		paperTemplateItemRouter.DELETE("deletePaperTemplateItemByIds", paperTemplateItemApi.DeletePaperTemplateItemByIds) // 批量删除PaperTemplateItem
		paperTemplateItemRouter.PUT("updatePaperTemplateItem", paperTemplateItemApi.UpdatePaperTemplateItem)              // 更新PaperTemplateItem
	}
	{
		paperTemplateItemRouterWithoutRecord.GET("findPaperTemplateItem", paperTemplateItemApi.FindPaperTemplateItem)       // 根据ID获取PaperTemplateItem
		paperTemplateItemRouterWithoutRecord.GET("getPaperTemplateItemList", paperTemplateItemApi.GetPaperTemplateItemList) // 获取PaperTemplateItem列表
	}
}
