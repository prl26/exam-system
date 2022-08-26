package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type QuestionBankOptionsApi struct {
}

var questionBank_optionsService = service.ServiceGroupApp.QuestionBankServiceGroup.OptionsService

// CreateQuestionBankOptions 创建QuestionBankOptions
// @Tags Options
// @Summary 创建QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.Options true "创建QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankOptions/createQuestionBankOptions [post]
func (questionBank_optionsApi *QuestionBankOptionsApi) CreateQuestionBankOptions(c *gin.Context) {
	var questionBank_options questionBank.Options
	_ = c.ShouldBindJSON(&questionBank_options)
	verify := utils.Rules{
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(questionBank_options, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionBank_optionsService.CreateQuestionBankOptions(questionBank_options); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionBankOptions 删除QuestionBankOptions
// @Tags Options
// @Summary 删除QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.Options true "删除QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBankOptions/deleteQuestionBankOptions [delete]
func (questionBank_optionsApi *QuestionBankOptionsApi) DeleteQuestionBankOptions(c *gin.Context) {
	var questionBank_options questionBank.Options
	_ = c.ShouldBindJSON(&questionBank_options)
	if err := questionBank_optionsService.DeleteQuestionBankOptions(questionBank_options); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionBankOptionsByIds 批量删除QuestionBankOptions
// @Tags Options
// @Summary 批量删除QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionBankOptions/deleteQuestionBankOptionsByIds [delete]
func (questionBank_optionsApi *QuestionBankOptionsApi) DeleteQuestionBankOptionsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := questionBank_optionsService.DeleteQuestionBankOptionsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionBankOptions 更新QuestionBankOptions
// @Tags Options
// @Summary 更新QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.Options true "更新QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBankOptions/updateQuestionBankOptions [put]
func (questionBank_optionsApi *QuestionBankOptionsApi) UpdateQuestionBankOptions(c *gin.Context) {
	var questionBank_options questionBank.Options
	_ = c.ShouldBindJSON(&questionBank_options)
	verify := utils.Rules{
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(questionBank_options, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionBank_optionsService.UpdateQuestionBankOptions(questionBank_options); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionBankOptions 用id查询QuestionBankOptions
// @Tags Options
// @Summary 用id查询QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBank.Options true "用id查询QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBankOptions/findQuestionBankOptions [get]
func (questionBank_optionsApi *QuestionBankOptionsApi) FindQuestionBankOptions(c *gin.Context) {
	var questionBank_options questionBank.Options
	_ = c.ShouldBindQuery(&questionBank_options)
	if requestionBank_options, err := questionBank_optionsService.GetQuestionBankOptions(questionBank_options.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionBank_options": requestionBank_options}, c)
	}
}

// GetQuestionBankOptionsList 分页获取QuestionBankOptions列表
// @Tags Options
// @Summary 分页获取QuestionBankOptions列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBankReq.QuestionBankOptionsSearch true "分页获取QuestionBankOptions列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankOptions/getQuestionBankOptionsList [get]
func (questionBank_optionsApi *QuestionBankOptionsApi) GetQuestionBankOptionsList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankOptionsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBank_optionsService.GetQuestionBankOptionsInfoList(pageInfo); err != nil {
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
