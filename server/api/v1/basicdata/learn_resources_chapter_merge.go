package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type LearnResourcesChapterMergeApi struct {
}

var learnResourcesChapterMergeService = service.ServiceGroupApp.BasicdataApiGroup.LearnResourcesChapterMergeService

// CreateLearnResourcesChapterMerge 创建LearnResourcesChapterMerge
// @Tags LearnResourcesChapterMerge
// @Summary 创建LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.LearnResourcesChapterMerge true "创建LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /learnResourcesChapterMerge/createLearnResourcesChapterMerge [post]
func (learnResourcesChapterMergeApi *LearnResourcesChapterMergeApi) CreateLearnResourcesChapterMerge(c *gin.Context) {
	var learnResourcesChapterMerge basicdata.LearnResourcesChapterMerge
	_ = c.ShouldBindJSON(&learnResourcesChapterMerge)
	if err := learnResourcesChapterMergeService.CreateLearnResourcesChapterMerge(learnResourcesChapterMerge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLearnResourcesChapterMerge 删除LearnResourcesChapterMerge
// @Tags LearnResourcesChapterMerge
// @Summary 删除LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.LearnResourcesChapterMerge true "删除LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /learnResourcesChapterMerge/deleteLearnResourcesChapterMerge [delete]
func (learnResourcesChapterMergeApi *LearnResourcesChapterMergeApi) DeleteLearnResourcesChapterMerge(c *gin.Context) {
	var learnResourcesChapterMerge basicdata.LearnResourcesChapterMerge
	_ = c.ShouldBindJSON(&learnResourcesChapterMerge)
	if err := learnResourcesChapterMergeService.DeleteLearnResourcesChapterMerge(learnResourcesChapterMerge); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLearnResourcesChapterMergeByIds 批量删除LearnResourcesChapterMerge
// @Tags LearnResourcesChapterMerge
// @Summary 批量删除LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /learnResourcesChapterMerge/deleteLearnResourcesChapterMergeByIds [delete]
func (learnResourcesChapterMergeApi *LearnResourcesChapterMergeApi) DeleteLearnResourcesChapterMergeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := learnResourcesChapterMergeService.DeleteLearnResourcesChapterMergeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLearnResourcesChapterMerge 更新LearnResourcesChapterMerge
// @Tags LearnResourcesChapterMerge
// @Summary 更新LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.LearnResourcesChapterMerge true "更新LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /learnResourcesChapterMerge/updateLearnResourcesChapterMerge [put]
func (learnResourcesChapterMergeApi *LearnResourcesChapterMergeApi) UpdateLearnResourcesChapterMerge(c *gin.Context) {
	var learnResourcesChapterMerge basicdata.LearnResourcesChapterMerge
	_ = c.ShouldBindJSON(&learnResourcesChapterMerge)
	if err := learnResourcesChapterMergeService.UpdateLearnResourcesChapterMerge(learnResourcesChapterMerge); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLearnResourcesChapterMerge 用id查询LearnResourcesChapterMerge
// @Tags LearnResourcesChapterMerge
// @Summary 用id查询LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.LearnResourcesChapterMerge true "用id查询LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /learnResourcesChapterMerge/findLearnResourcesChapterMerge [get]
func (learnResourcesChapterMergeApi *LearnResourcesChapterMergeApi) FindLearnResourcesChapterMerge(c *gin.Context) {
	var learnResourcesChapterMerge basicdata.LearnResourcesChapterMerge
	_ = c.ShouldBindQuery(&learnResourcesChapterMerge)
	if relearnResourcesChapterMerge, err := learnResourcesChapterMergeService.GetLearnResourcesChapterMerge(learnResourcesChapterMerge.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relearnResourcesChapterMerge": relearnResourcesChapterMerge}, c)
	}
}

// GetLearnResourcesChapterMergeList 分页获取LearnResourcesChapterMerge列表
// @Tags LearnResourcesChapterMerge
// @Summary 分页获取LearnResourcesChapterMerge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.LearnResourcesChapterMergeSearch true "分页获取LearnResourcesChapterMerge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /learnResourcesChapterMerge/getLearnResourcesChapterMergeList [get]
func (learnResourcesChapterMergeApi *LearnResourcesChapterMergeApi) GetLearnResourcesChapterMergeList(c *gin.Context) {
	var pageInfo basicdataReq.LearnResourcesChapterMergeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := learnResourcesChapterMergeService.GetLearnResourcesChapterMergeInfoList(pageInfo); err != nil {
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
