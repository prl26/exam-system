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

type QuestionBankProgrammCaseApi struct {
}

var questionBankProgrammCaseService = service.ServiceGroupApp.QuestionBankServiceGroup.ProgrammCaseService

// CreateQuestionBankProgrammCase 创建QuestionBankProgrammCase
// @Tags QuestionBankProgrammCase
// @Summary 创建QuestionBankProgrammCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.ProgrammCase true "创建QuestionBankProgrammCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankProgrammCase/createQuestionBankProgrammCase [post]
func (questionBankProgrammCaseApi *QuestionBankProgrammCaseApi) CreateQuestionBankProgrammCase(c *gin.Context) {
	var questionBankProgrammCase questionBank.ProgrammCase
	_ = c.ShouldBindJSON(&questionBankProgrammCase)
	verify := utils.Rules{
		"ProgrammId": {utils.NotEmpty()},
		"Score":      {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
		"Output":     {utils.NotEmpty()},
	}
	if err := utils.Verify(questionBankProgrammCase, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionBankProgrammCaseService.CreateQuestionBankProgrammCase(questionBankProgrammCase); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionBankProgrammCase 删除QuestionBankProgrammCase
// @Tags QuestionBankProgrammCase
// @Summary 删除QuestionBankProgrammCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.ProgrammCase true "删除QuestionBankProgrammCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBankProgrammCase/deleteQuestionBankProgrammCase [delete]
func (questionBankProgrammCaseApi *QuestionBankProgrammCaseApi) DeleteQuestionBankProgrammCase(c *gin.Context) {
	var questionBankProgrammCase questionBank.ProgrammCase
	_ = c.ShouldBindJSON(&questionBankProgrammCase)
	if err := questionBankProgrammCaseService.DeleteQuestionBankProgrammCase(questionBankProgrammCase); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionBankProgrammCaseByIds 批量删除QuestionBankProgrammCase
// @Tags QuestionBankProgrammCase
// @Summary 批量删除QuestionBankProgrammCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankProgrammCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionBankProgrammCase/deleteQuestionBankProgrammCaseByIds [delete]
func (questionBankProgrammCaseApi *QuestionBankProgrammCaseApi) DeleteQuestionBankProgrammCaseByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := questionBankProgrammCaseService.DeleteQuestionBankProgrammCaseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionBankProgrammCase 更新QuestionBankProgrammCase
// @Tags QuestionBankProgrammCase
// @Summary 更新QuestionBankProgrammCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.ProgrammCase true "更新QuestionBankProgrammCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBankProgrammCase/updateQuestionBankProgrammCase [put]
func (questionBankProgrammCaseApi *QuestionBankProgrammCaseApi) UpdateQuestionBankProgrammCase(c *gin.Context) {
	var questionBankProgrammCase questionBank.ProgrammCase
	_ = c.ShouldBindJSON(&questionBankProgrammCase)
	verify := utils.Rules{
		"ProgrammId": {utils.NotEmpty()},
		"Score":      {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
		"OutType":    {utils.NotEmpty()},
		"Output":     {utils.NotEmpty()},
	}
	if err := utils.Verify(questionBankProgrammCase, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionBankProgrammCaseService.UpdateQuestionBankProgrammCase(questionBankProgrammCase); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionBankProgrammCase 用id查询QuestionBankProgrammCase
// @Tags QuestionBankProgrammCase
// @Summary 用id查询QuestionBankProgrammCase
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBank.ProgrammCase true "用id查询QuestionBankProgrammCase"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBankProgrammCase/findQuestionBankProgrammCase [get]
func (questionBankProgrammCaseApi *QuestionBankProgrammCaseApi) FindQuestionBankProgrammCase(c *gin.Context) {
	var questionBankProgrammCase questionBank.ProgrammCase
	_ = c.ShouldBindQuery(&questionBankProgrammCase)
	if requestionBankProgrammCase, err := questionBankProgrammCaseService.GetQuestionBankProgrammCase(questionBankProgrammCase.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionBankProgrammCase": requestionBankProgrammCase}, c)
	}
}

// GetQuestionBankProgrammCaseList 分页获取QuestionBankProgrammCase列表
// @Tags QuestionBankProgrammCase
// @Summary 分页获取QuestionBankProgrammCase列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.QuestionBankProgrammCaseSearch true "分页获取QuestionBankProgrammCase列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankProgrammCase/getQuestionBankProgrammCaseList [get]
func (questionBankProgrammCaseApi *QuestionBankProgrammCaseApi) GetQuestionBankProgrammCaseList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankProgrammCaseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBankProgrammCaseService.GetQuestionBankProgrammCaseInfoList(pageInfo); err != nil {
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
