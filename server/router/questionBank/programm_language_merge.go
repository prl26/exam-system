package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProgrammLanguageMergeRouter struct {
}

// InitProgrammLanguageMergeRouter 初始化 ProgrammLanguageMerge 路由信息
func (s *ProgrammLanguageMergeRouter) InitProgrammLanguageMergeRouter(Router *gin.RouterGroup) {
	programmLanguageMergeRouter := Router.Group("programmLanguageMerge").Use(middleware.OperationRecord())
	programmLanguageMergeRouterWithoutRecord := Router.Group("programmLanguageMerge")
	var programmLanguageMergeApi = v1.ApiGroupApp.QuestionBankApiGroup.ProgrammLanguageMergeApi
	{
		programmLanguageMergeRouter.POST("createProgrammLanguageMerge", programmLanguageMergeApi.CreateProgrammLanguageMerge)             // 新建ProgrammLanguageMerge
		programmLanguageMergeRouter.DELETE("deleteProgrammLanguageMerge", programmLanguageMergeApi.DeleteProgrammLanguageMerge)           // 删除ProgrammLanguageMerge
		programmLanguageMergeRouter.DELETE("deleteProgrammLanguageMergeByIds", programmLanguageMergeApi.DeleteProgrammLanguageMergeByIds) // 批量删除ProgrammLanguageMerge
		programmLanguageMergeRouter.PUT("updateProgrammLanguageMerge", programmLanguageMergeApi.UpdateProgrammLanguageMerge)              // 更新ProgrammLanguageMerge
	}
	{
		programmLanguageMergeRouterWithoutRecord.GET("findProgrammLanguageMerge", programmLanguageMergeApi.FindProgrammLanguageMerge)       // 根据ID获取ProgrammLanguageMerge
		programmLanguageMergeRouterWithoutRecord.GET("getProgrammLanguageMergeList", programmLanguageMergeApi.GetProgrammLanguageMergeList) // 获取ProgrammLanguageMerge列表
	}
}
