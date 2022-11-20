package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type PaperQuestionMergeApi struct {
}

var paperQuestionMergeService = service.ServiceGroupApp.ExammanageServiceGroup.PaperQuestionMergeService

// CreatePaperQuestionMerge 创建PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 创建PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.PaperQuestionMerge true "创建PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paperQuestionMerge/createPaperQuestionMerge [post]
func (paperQuestionMergeApi *PaperQuestionMergeApi) CreatePaperQuestionMerge(c *gin.Context) {
	var paperQuestionMerge examManage.PaperQuestionMerge
	_ = c.ShouldBindJSON(&paperQuestionMerge)
	if err := paperQuestionMergeService.CreatePaperQuestionMerge(paperQuestionMerge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePaperQuestionMerge 删除PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 删除PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.PaperQuestionMerge true "删除PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /paperQuestionMerge/deletePaperQuestionMerge [delete]
func (paperQuestionMergeApi *PaperQuestionMergeApi) DeletePaperQuestionMerge(c *gin.Context) {
	var paperQuestionMerge examManage.PaperQuestionMerge
	_ = c.ShouldBindJSON(&paperQuestionMerge)
	if err := paperQuestionMergeService.DeletePaperQuestionMerge(paperQuestionMerge); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePaperQuestionMergeByIds 批量删除PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 批量删除PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /paperQuestionMerge/deletePaperQuestionMergeByIds [delete]
func (paperQuestionMergeApi *PaperQuestionMergeApi) DeletePaperQuestionMergeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := paperQuestionMergeService.DeletePaperQuestionMergeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePaperQuestionMerge 更新PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 更新PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.PaperQuestionMerge true "更新PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /paperQuestionMerge/updatePaperQuestionMerge [put]
func (paperQuestionMergeApi *PaperQuestionMergeApi) UpdatePaperQuestionMerge(c *gin.Context) {
	var paperQuestionMerge examManage.PaperQuestionMerge
	_ = c.ShouldBindJSON(&paperQuestionMerge)
	if err := paperQuestionMergeService.UpdatePaperQuestionMerge(paperQuestionMerge); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPaperQuestionMerge 用id查询PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 用id查询PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query frontExamManage.PaperQuestionMerge true "用id查询PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /paperQuestionMerge/findPaperQuestionMerge [get]
func (paperQuestionMergeApi *PaperQuestionMergeApi) FindPaperQuestionMerge(c *gin.Context) {
	var paperQuestionMerge examManage.PaperQuestionMerge
	_ = c.ShouldBindQuery(&paperQuestionMerge)
	if repaperQuestionMerge, err := paperQuestionMergeService.GetPaperQuestionMerge(paperQuestionMerge.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repaperQuestionMerge": repaperQuestionMerge}, c)
	}
}

// GetPaperQuestionMergeList 分页获取PaperQuestionMerge列表
// @Tags PaperQuestionMerge
// @Summary 分页获取PaperQuestionMerge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManageReq.PaperQuestionMergeSearch true "分页获取PaperQuestionMerge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paperQuestionMerge/getPaperQuestionMergeList [get]
func (paperQuestionMergeApi *PaperQuestionMergeApi) GetPaperQuestionMergeList(c *gin.Context) {
	var pageInfo examManageReq.PaperQuestionMergeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := paperQuestionMergeService.GetPaperQuestionMergeInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
