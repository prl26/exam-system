package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/prl26/exam-system/server/api"
)

type AutoCodeHistoryRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	autoCodeHistoryRouter := Router.Group("autoCode")
	autoCodeHistoryApi := v1.ApiGroupApp.BackStage.SystemApiGroup.AutoCodeHistoryApi
	{
		autoCodeHistoryRouter.POST("getMeta", autoCodeHistoryApi.First)         // 根据id获取meta信息
		autoCodeHistoryRouter.POST("rollback", autoCodeHistoryApi.RollBack)     // 回滚
		autoCodeHistoryRouter.POST("delSysHistory", autoCodeHistoryApi.Delete)  // 删除回滚记录
		autoCodeHistoryRouter.POST("getSysHistory", autoCodeHistoryApi.GetList) // 获取回滚记录分页
	}
}
