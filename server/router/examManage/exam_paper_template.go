package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/v1"
	"github.com/prl26/exam-system/server/middleware"
)

type ExamPaperTemplateRouter struct {
}

// InitExamPaperTemplateRouter 初始化 ExamPaperTemplate 路由信息
func (s *ExamPaperTemplateRouter) InitExamPaperTemplateRouter(Router *gin.RouterGroup) {
	examPaperTemplateRouter := Router.Group("examPaperTemplate").Use(middleware.OperationRecord())
	examPaperTemplateRouterWithoutRecord := Router.Group("examPaperTemplate")
	var examPaperTemplateApi = v1.ApiGroupApp.ExammanageApiGroup.PaperTemplateApi
	{
		examPaperTemplateRouter.POST("createExamPaperTemplate", examPaperTemplateApi.CreatePaperTemplate)             // 新建ExamPaperTemplate
		examPaperTemplateRouter.DELETE("deleteExamPaperTemplate", examPaperTemplateApi.DeletePaperTemplate)           // 删除ExamPaperTemplate
		examPaperTemplateRouter.DELETE("deleteExamPaperTemplateByIds", examPaperTemplateApi.DeletePaperTemplateByIds) // 批量删除ExamPaperTemplate
		examPaperTemplateRouter.PUT("updateExamPaperTemplate", examPaperTemplateApi.UpdatePaperTemplate)              // 更新ExamPaperTemplate
	}
	{
		examPaperTemplateRouterWithoutRecord.GET("findExamPaperTemplate", examPaperTemplateApi.FindPaperTemplate)       // 根据ID获取ExamPaperTemplate
		examPaperTemplateRouterWithoutRecord.GET("getExamPaperTemplateList", examPaperTemplateApi.GetPaperTemplateList) // 获取ExamPaperTemplate列表
	}
}
