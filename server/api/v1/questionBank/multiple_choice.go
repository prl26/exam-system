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

type QuestionBankMultipleChoiceApi struct {
}

var questionBank_multiple_choiceService = service.ServiceGroupApp.QuestionBankServiceGroup.MultipleChoiceService

// CreateQuestionBankMultipleChoice 创建QuestionBankMultipleChoice
// @Tags MultipleChoice
// @Summary 创建QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.MultipleChoice true "创建QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_multiple_choice/createQuestionBankMultipleChoice [post]
func (questionBank_multiple_choiceApi *QuestionBankMultipleChoiceApi) CreateQuestionBankMultipleChoice(c *gin.Context) {
	var questionBank_multiple_choice questionBank.MultipleChoice
	_ = c.ShouldBindJSON(&questionBank_multiple_choice)
	if err := questionBank_multiple_choiceService.CreateQuestionBankMultipleChoice(questionBank_multiple_choice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionBankMultipleChoice 删除QuestionBankMultipleChoice
// @Tags MultipleChoice
// @Summary 删除QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.MultipleChoice true "删除QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_multiple_choice/deleteQuestionBankMultipleChoice [delete]
func (questionBank_multiple_choiceApi *QuestionBankMultipleChoiceApi) DeleteQuestionBankMultipleChoice(c *gin.Context) {
	var questionBank_multiple_choice questionBank.MultipleChoice
	_ = c.ShouldBindJSON(&questionBank_multiple_choice)
	if err := questionBank_multiple_choiceService.DeleteQuestionBankMultipleChoice(questionBank_multiple_choice); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionBankMultipleChoiceByIds 批量删除QuestionBankMultipleChoice
// @Tags MultipleChoice
// @Summary 批量删除QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionBank_multiple_choice/deleteQuestionBankMultipleChoiceByIds [delete]
func (questionBank_multiple_choiceApi *QuestionBankMultipleChoiceApi) DeleteQuestionBankMultipleChoiceByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := questionBank_multiple_choiceService.DeleteQuestionBankMultipleChoiceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionBankMultipleChoice 更新QuestionBankMultipleChoice
// @Tags MultipleChoice
// @Summary 更新QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.MultipleChoice true "更新QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBank_multiple_choice/updateQuestionBankMultipleChoice [put]
func (questionBank_multiple_choiceApi *QuestionBankMultipleChoiceApi) UpdateQuestionBankMultipleChoice(c *gin.Context) {
	var questionBank_multiple_choice questionBank.MultipleChoice
	_ = c.ShouldBindJSON(&questionBank_multiple_choice)
	if err := questionBank_multiple_choiceService.UpdateQuestionBankMultipleChoice(questionBank_multiple_choice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionBankMultipleChoice 用id查询QuestionBankMultipleChoice
// @Tags MultipleChoice
// @Summary 用id查询QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBank.MultipleChoice true "用id查询QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBank_multiple_choice/findQuestionBankMultipleChoice [get]
func (questionBank_multiple_choiceApi *QuestionBankMultipleChoiceApi) FindQuestionBankMultipleChoice(c *gin.Context) {
	var questionBank_multiple_choice questionBank.MultipleChoice
	_ = c.ShouldBindQuery(&questionBank_multiple_choice)
	if requestionBank_multiple_choice, err := questionBank_multiple_choiceService.GetQuestionBankMultipleChoice(questionBank_multiple_choice.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionBank_multiple_choice": requestionBank_multiple_choice}, c)
	}
}

// GetQuestionBankMultipleChoiceList 分页获取QuestionBankMultipleChoice列表
// @Tags MultipleChoice
// @Summary 分页获取QuestionBankMultipleChoice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBankReq.QuestionBankMultipleChoiceSearch true "分页获取QuestionBankMultipleChoice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_multiple_choice/getQuestionBankMultipleChoiceList [get]
func (questionBank_multiple_choiceApi *QuestionBankMultipleChoiceApi) GetQuestionBankMultipleChoiceList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankMultipleChoiceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBank_multiple_choiceService.GetQuestionBankMultipleChoiceInfoList(pageInfo); err != nil {
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
