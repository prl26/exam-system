package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api"
	"github.com/prl26/exam-system/server/middleware"
)

type PaperTemplateRouter struct {
}

// InitPaperTemplateRouter 初始化 PaperTemplate 路由信息
func (s *PaperTemplateRouter) InitPaperTemplateRouter(Router *gin.RouterGroup) {
	PapertemplateRouter := Router.Group("Papertemplate").Use(middleware.OperationRecord())
	PapertemplateRouterWithoutRecord := Router.Group("Papertemplate")
	var PapertemplateApi = api.ApiGroupApp.BackStage.ExamManageApiGroup.PaperTemplateApi
	{
		PapertemplateRouter.POST("createPaperTemplate", PapertemplateApi.CreatePaperTemplate)             // 新建PaperTemplate
		PapertemplateRouter.GET("deletePaperTemplate", PapertemplateApi.DeletePaperTemplate)              // 删除PaperTemplate
		PapertemplateRouter.DELETE("deletePaperTemplateByIds", PapertemplateApi.DeletePaperTemplateByIds) // 批量删除PaperTemplate
		PapertemplateRouter.PUT("updatePaperTemplate", PapertemplateApi.UpdatePaperTemplate)              // 更新PaperTemplate
		PapertemplateRouter.GET("beforeTemplate", PapertemplateApi.BeforeTemplate)
	}
	{
		PapertemplateRouterWithoutRecord.GET("findPaperTemplate", PapertemplateApi.FindPaperTemplate)       // 根据ID获取PaperTemplate
		PapertemplateRouterWithoutRecord.GET("getPaperTemplateList", PapertemplateApi.GetPaperTemplateList) // 获取PaperTemplate列表
	}
}
