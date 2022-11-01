package teachplan

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type ScoreApi struct {
}

var scoreService = service.ServiceGroupApp.TeachplanServiceGroup.ScoreService

// CreateScore 创建Score
// @Tags Score
// @Summary 创建Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.Score true "创建Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /score/createScore [post]
func (scoreApi *ScoreApi) CreateScore(c *gin.Context) {
	var score teachplan.Score
	_ = c.ShouldBindJSON(&score)
	if err := scoreService.CreateScore(score); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteScore 删除Score
// @Tags Score
// @Summary 删除Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.Score true "删除Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /score/deleteScore [delete]
func (scoreApi *ScoreApi) DeleteScore(c *gin.Context) {
	var score teachplan.Score
	_ = c.ShouldBindJSON(&score)
	if err := scoreService.DeleteScore(score); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteScoreByIds 批量删除Score
// @Tags Score
// @Summary 批量删除Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /score/deleteScoreByIds [delete]
func (scoreApi *ScoreApi) DeleteScoreByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := scoreService.DeleteScoreByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateScore 更新Score
// @Tags Score
// @Summary 更新Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.Score true "更新Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /score/updateScore [put]
func (scoreApi *ScoreApi) UpdateScore(c *gin.Context) {
	var score teachplan.Score
	_ = c.ShouldBindJSON(&score)
	if err := scoreService.UpdateScore(score); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindScore 用id查询Score
// @Tags Score
// @Summary 用id查询Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplan.Score true "用id查询Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /score/findScore [get]
func (scoreApi *ScoreApi) FindScore(c *gin.Context) {
	var score teachplan.Score
	_ = c.ShouldBindQuery(&score)
	if rescore, err := scoreService.GetScore(score.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rescore": rescore}, c)
	}
}

// GetScoreList 分页获取Score列表
// @Tags Score
// @Summary 分页获取Score列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplanReq.ScoreSearch true "分页获取Score列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /score/getScoreList [get]
func (scoreApi *ScoreApi) GetScoreList(c *gin.Context) {
	var pageInfo teachplanReq.ScoreSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := scoreService.GetScoreInfoList(pageInfo); err != nil {
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
