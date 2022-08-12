package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProfessionalRouter struct {
}

// InitProfessionalRouter 初始化 Professional 路由信息
func (s *ProfessionalRouter) InitProfessionalRouter(Router *gin.RouterGroup) {
	professionalRouter := Router.Group("professional").Use(middleware.OperationRecord())
	professionalRouterWithoutRecord := Router.Group("professional")
	var professionalApi = v1.ApiGroupApp.BasicdataApiGroup.ProfessionalApi
	{
		professionalRouter.POST("createProfessional", professionalApi.CreateProfessional)   // 新建Professional
		professionalRouter.DELETE("deleteProfessional", professionalApi.DeleteProfessional) // 删除Professional
		professionalRouter.DELETE("deleteProfessionalByIds", professionalApi.DeleteProfessionalByIds) // 批量删除Professional
		professionalRouter.PUT("updateProfessional", professionalApi.UpdateProfessional)    // 更新Professional
	}
	{
		professionalRouterWithoutRecord.GET("findProfessional", professionalApi.FindProfessional)        // 根据ID获取Professional
		professionalRouterWithoutRecord.GET("getProfessionalList", professionalApi.GetProfessionalList)  // 获取Professional列表
	}
}
