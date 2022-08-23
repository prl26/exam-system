package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PaperTemplateRouter struct {
}

// InitPaperTemplateRouter 初始化 PaperTemplate 路由信息
func (s *PaperTemplateRouter) InitPaperTemplateRouter(Router *gin.RouterGroup) {
	PapertemplateRouter := Router.Group("Papertemplate").Use(middleware.OperationRecord())
	PapertemplateRouterWithoutRecord := Router.Group("Papertemplate")
	var PapertemplateApi = v1.ApiGroupApp.ExammanageApiGroup.PaperTemplateApi
	{
		PapertemplateRouter.POST("createPaperTemplate", PapertemplateApi.CreatePaperTemplate)   // 新建PaperTemplate
		PapertemplateRouter.DELETE("deletePaperTemplate", PapertemplateApi.DeletePaperTemplate) // 删除PaperTemplate
		PapertemplateRouter.DELETE("deletePaperTemplateByIds", PapertemplateApi.DeletePaperTemplateByIds) // 批量删除PaperTemplate
		PapertemplateRouter.PUT("updatePaperTemplate", PapertemplateApi.UpdatePaperTemplate)    // 更新PaperTemplate
	}
	{
		PapertemplateRouterWithoutRecord.GET("findPaperTemplate", PapertemplateApi.FindPaperTemplate)        // 根据ID获取PaperTemplate
		PapertemplateRouterWithoutRecord.GET("getPaperTemplateList", PapertemplateApi.GetPaperTemplateList)  // 获取PaperTemplate列表
	}
}
