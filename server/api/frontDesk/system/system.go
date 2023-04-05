package system

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/system"
	systemRes "github.com/prl26/exam-system/server/model/system/response"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
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

//上传考试截图
func (s *SystemApi) UploadFile(c *gin.Context) {
	var plan teachplanReq.ExamPlan
	_ = c.ShouldBindQuery(&plan)
	sd := utils.GetStudentId(c)
	var file system.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = FrontSystemService.UploadFile(header, noSave, plan.PlanId, sd) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(systemRes.ExaFileResponse{File: file}, "上传成功", c)
}
