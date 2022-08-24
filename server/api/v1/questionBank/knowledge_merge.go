package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	questionBankReq "github.com/flipped-aurora/gin-vue-admin/server/model/questionBank/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type QuestionBankKnowledgeMergeApi struct {
}

var questionBankKnowledgeMergeService = service.ServiceGroupApp.QuestionBankServiceGroup.KnowledgeMergeService

// CreateQuestionBankKnowledgeMerge 创建QuestionBankKnowledgeMerge
// @Tags QuestionBankKnowledgeMerge
// @Summary 创建QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.KnowledgeMerge true "创建QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankKnowledgeMerge/createQuestionBankKnowledgeMerge [post]
func (questionBankKnowledgeMergeApi *QuestionBankKnowledgeMergeApi) CreateQuestionBankKnowledgeMerge(c *gin.Context) {
	var questionBankKnowledgeMerge questionBank.KnowledgeMerge
	_ = c.ShouldBindJSON(&questionBankKnowledgeMerge)
	if err := questionBankKnowledgeMergeService.CreateQuestionBankKnowledgeMerge(questionBankKnowledgeMerge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionBankKnowledgeMerge 删除QuestionBankKnowledgeMerge
// @Tags QuestionBankKnowledgeMerge
// @Summary 删除QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.KnowledgeMerge true "删除QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBankKnowledgeMerge/deleteQuestionBankKnowledgeMerge [delete]
func (questionBankKnowledgeMergeApi *QuestionBankKnowledgeMergeApi) DeleteQuestionBankKnowledgeMerge(c *gin.Context) {
	var questionBankKnowledgeMerge questionBank.KnowledgeMerge
	_ = c.ShouldBindJSON(&questionBankKnowledgeMerge)
	if err := questionBankKnowledgeMergeService.DeleteQuestionBankKnowledgeMerge(questionBankKnowledgeMerge); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionBankKnowledgeMergeByIds 批量删除QuestionBankKnowledgeMerge
// @Tags QuestionBankKnowledgeMerge
// @Summary 批量删除QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionBankKnowledgeMerge/deleteQuestionBankKnowledgeMergeByIds [delete]
func (questionBankKnowledgeMergeApi *QuestionBankKnowledgeMergeApi) DeleteQuestionBankKnowledgeMergeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := questionBankKnowledgeMergeService.DeleteQuestionBankKnowledgeMergeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionBankKnowledgeMerge 更新QuestionBankKnowledgeMerge
// @Tags QuestionBankKnowledgeMerge
// @Summary 更新QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.KnowledgeMerge true "更新QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBankKnowledgeMerge/updateQuestionBankKnowledgeMerge [put]
func (questionBankKnowledgeMergeApi *QuestionBankKnowledgeMergeApi) UpdateQuestionBankKnowledgeMerge(c *gin.Context) {
	var questionBankKnowledgeMerge questionBank.KnowledgeMerge
	_ = c.ShouldBindJSON(&questionBankKnowledgeMerge)
	if err := questionBankKnowledgeMergeService.UpdateQuestionBankKnowledgeMerge(questionBankKnowledgeMerge); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionBankKnowledgeMerge 用id查询QuestionBankKnowledgeMerge
// @Tags QuestionBankKnowledgeMerge
// @Summary 用id查询QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBank.KnowledgeMerge true "用id查询QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBankKnowledgeMerge/findQuestionBankKnowledgeMerge [get]
func (questionBankKnowledgeMergeApi *QuestionBankKnowledgeMergeApi) FindQuestionBankKnowledgeMerge(c *gin.Context) {
	var questionBankKnowledgeMerge questionBank.KnowledgeMerge
	_ = c.ShouldBindQuery(&questionBankKnowledgeMerge)
	if requestionBankKnowledgeMerge, err := questionBankKnowledgeMergeService.GetQuestionBankKnowledgeMerge(questionBankKnowledgeMerge.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionBankKnowledgeMerge": requestionBankKnowledgeMerge}, c)
	}
}

// GetQuestionBankKnowledgeMergeList 分页获取QuestionBankKnowledgeMerge列表
// @Tags QuestionBankKnowledgeMerge
// @Summary 分页获取QuestionBankKnowledgeMerge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.QuestionBankKnowledgeMergeSearch true "分页获取QuestionBankKnowledgeMerge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankKnowledgeMerge/getQuestionBankKnowledgeMergeList [get]
func (questionBankKnowledgeMergeApi *QuestionBankKnowledgeMergeApi) GetQuestionBankKnowledgeMergeList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankKnowledgeMergeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBankKnowledgeMergeService.GetQuestionBankKnowledgeMergeInfoList(pageInfo); err != nil {
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
