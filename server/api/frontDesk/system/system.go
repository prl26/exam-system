package system

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type SystemApi struct{}

var termService = service.ServiceGroupApp.BasicdataApiGroup.TermService

func (s *SystemApi) GetTerms(c *gin.Context) {
	var pageInfo basicdataReq.FrontTermSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := FrontSystemService.GetTermInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		termNow, _ := termService.GetTermNow()
		response.OkWithDetailed(response.PageResultAndTerm{
			TermNow: termNow,
			List:    list,
			Total:   total,
		}, "获取成功", c)
	}
}
func (lessonApi *SystemApi) GetLessons(c *gin.Context) {
	var pageInfo basicdataReq.FrontLessonSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := FrontSystemService.GetLessonInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.FrontResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
