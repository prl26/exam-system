package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExamPaperTemplateRouter struct {
}

// InitExamPaperTemplateRouter 初始化 ExamPaperTemplate 路由信息
func (s *ExamPaperTemplateRouter) InitExamPaperTemplateRouter(Router *gin.RouterGroup) {
	examPaperTemplateRouter := Router.Group("examPaperTemplate").Use(middleware.OperationRecord())
	examPaperTemplateRouterWithoutRecord := Router.Group("examPaperTemplate")
	var examPaperTemplateApi = v1.ApiGroupApp.ExammanageApiGroup.ExamPaperTemplateApi
	{
		examPaperTemplateRouter.POST("createExamPaperTemplate", examPaperTemplateApi.CreateExamPaperTemplate)             // 新建ExamPaperTemplate
		examPaperTemplateRouter.DELETE("deleteExamPaperTemplate", examPaperTemplateApi.DeleteExamPaperTemplate)           // 删除ExamPaperTemplate
		examPaperTemplateRouter.DELETE("deleteExamPaperTemplateByIds", examPaperTemplateApi.DeleteExamPaperTemplateByIds) // 批量删除ExamPaperTemplate
		examPaperTemplateRouter.PUT("updateExamPaperTemplate", examPaperTemplateApi.UpdateExamPaperTemplate)              // 更新ExamPaperTemplate
	}
	{
		examPaperTemplateRouterWithoutRecord.GET("findExamPaperTemplate", examPaperTemplateApi.FindExamPaperTemplate)       // 根据ID获取ExamPaperTemplate
		examPaperTemplateRouterWithoutRecord.GET("getExamPaperTemplateList", examPaperTemplateApi.GetExamPaperTemplateList) // 获取ExamPaperTemplate列表
	}
}
