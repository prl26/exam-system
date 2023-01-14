package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type SituationApi struct {
}

var situationService = service.ServiceGroupApp.QuestionBankServiceGroup.SituationService

func (a SituationApi) FindTeachClassSituation(c *gin.Context) {
	var req questionBankReq.TeachClassSituation
	_ = c.ShouldBind(&req)
	verify := utils.Rules{
		"Page":         {utils.NotEmpty()},
		"PageSize":     {utils.NotEmpty()},
		"LessonId":     {utils.NotEmpty()},
		"TeachClassId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if list, total, err := situationService.FindTeachClassSituation(req.PageInfo, req.LessonId, req.TeachClassId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, err)
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

func (a SituationApi) FindStudentSituation(c *gin.Context) {
	var req questionBankReq.StudentSituation
	_ = c.ShouldBind(&req)
	verify := utils.Rules{
		"Page":      {utils.NotEmpty()},
		"PageSize":  {utils.NotEmpty()},
		"LessonId":  {utils.NotEmpty()},
		"StudentId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	situationService.UpdateSituation(req.LessonId, req.StudentId)
	if list, total, err := situationService.FindStudentSituation(req.PageInfo, req.LessonId, req.StudentId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, err)
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

func (a SituationApi) FindDetail(c *gin.Context) {
	var req questionBankReq.SituationDetail
	_ = c.ShouldBind(&req)
	verify := utils.Rules{
		"Page":     {utils.NotEmpty()},
		"PageSize": {utils.NotEmpty()},
		"RecordId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if list, total, err := situationService.FindDetail(req.PageInfo, req.RecordId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, err)
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}
