package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type TermApi struct {
}

var termService = service.ServiceGroupApp.BasicdataApiGroup.TermService

// CreateTerm 创建Term
// @Tags Term
// @Summary 创建Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Term true "创建Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /term/createTerm [post]
func (termApi *TermApi) CreateTerm(c *gin.Context) {
	var term basicdata.Term
	_ = c.ShouldBindJSON(&term)
	verify := utils.Rules{
		"Name":      {utils.NotEmpty()},
		"StartTime": {utils.NotEmpty()},
	}
	if err := utils.Verify(term, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := termService.CreateTerm(term); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTerm 删除Term
// @Tags Term
// @Summary 删除Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Term true "删除Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /term/deleteTerm [delete]
func (termApi *TermApi) DeleteTerm(c *gin.Context) {
	var term basicdata.Term
	_ = c.ShouldBindJSON(&term)
	if err := termService.DeleteTerm(term); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTermByIds 批量删除Term
// @Tags Term
// @Summary 批量删除Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /term/deleteTermByIds [delete]
func (termApi *TermApi) DeleteTermByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := termService.DeleteTermByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTerm 更新Term
// @Tags Term
// @Summary 更新Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Term true "更新Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /term/updateTerm [put]
func (termApi *TermApi) UpdateTerm(c *gin.Context) {
	var term basicdata.Term
	_ = c.ShouldBindJSON(&term)
	verify := utils.Rules{
		"Name":       {utils.NotEmpty()},
		"Start_time": {utils.NotEmpty()},
	}
	if err := utils.Verify(term, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := termService.UpdateTerm(term); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTerm 用id查询Term
// @Tags Term
// @Summary 用id查询Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.Term true "用id查询Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /term/findTerm [get]
func (termApi *TermApi) FindTerm(c *gin.Context) {
	var term basicdata.Term
	_ = c.ShouldBindQuery(&term)
	if reterm, err := termService.GetTerm(term.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		termNow, _ := termService.GetTermNow()
		response.OkWithData(gin.H{
			"reterm":  reterm,
			"termNow": termNow,
		}, c)
	}
}

// GetTermList 分页获取Term列表
// @Tags Term
// @Summary 分页获取Term列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.TermSearch true "分页获取Term列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /term/getTermList [get]
func (termApi *TermApi) GetTermList(c *gin.Context) {
	var pageInfo basicdataReq.TermSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := termService.GetTermInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
