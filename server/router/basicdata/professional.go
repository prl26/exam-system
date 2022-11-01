package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/api/Backstage"
	"github.com/prl26/exam-system/server/middleware"
)

type ProfessionalRouter struct {
}

// InitProfessionalRouter 初始化 Professional 路由信息
func (s *ProfessionalRouter) InitProfessionalRouter(Router *gin.RouterGroup) {
	professionalRouter := Router.Group("professional").Use(middleware.OperationRecord())
	professionalRouterWithoutRecord := Router.Group("professional")
	var professionalApi = Backstage.ApiGroupApp.BasicdataApiGroup.ProfessionalApi
	{
		professionalRouter.POST("createProfessional", professionalApi.CreateProfessional)             // 新建Professional
		professionalRouter.DELETE("deleteProfessional", professionalApi.DeleteProfessional)           // 删除Professional
		professionalRouter.DELETE("deleteProfessionalByIds", professionalApi.DeleteProfessionalByIds) // 批量删除Professional
		professionalRouter.PUT("updateProfessional", professionalApi.UpdateProfessional)              // 更新Professional
	}
	{
		professionalRouterWithoutRecord.GET("findProfessional", professionalApi.FindProfessional)       // 根据ID获取Professional
		professionalRouterWithoutRecord.GET("getProfessionalList", professionalApi.GetProfessionalList) // 获取Professional列表
	}
}
