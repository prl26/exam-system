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

type QuestionBankJudgeApi struct {
}

var questionBank_judgeService = service.ServiceGroupApp.QuestionBankServiceGroup.JudgeService

// CreateQuestionBankJudge 创建QuestionBankJudge
// @Tags QuestionBankJudge
// @Summary 创建QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.Judge true "创建QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankJudge/createQuestionBankJudge [post]
func (questionBank_judgeApi *QuestionBankJudgeApi) CreateQuestionBankJudge(c *gin.Context) {
	var questionBank_judge questionBank.Judge
	_ = c.ShouldBindJSON(&questionBank_judge)
	verify := utils.Rules{
		"Describe": {utils.NotEmpty()},
		"Is_right": {utils.NotEmpty()},
	}
	if err := utils.Verify(questionBank_judge, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionBank_judgeService.CreateQuestionBankJudge(questionBank_judge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionBankJudge 删除QuestionBankJudge
// @Tags QuestionBankJudge
// @Summary 删除QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.Judge true "删除QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBankJudge/deleteQuestionBankJudge [delete]
func (questionBank_judgeApi *QuestionBankJudgeApi) DeleteQuestionBankJudge(c *gin.Context) {
	var questionBank_judge questionBank.Judge
	_ = c.ShouldBindJSON(&questionBank_judge)
	if err := questionBank_judgeService.DeleteQuestionBankJudge(questionBank_judge); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionBankJudgeByIds 批量删除QuestionBankJudge
// @Tags QuestionBankJudge
// @Summary 批量删除QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionBankJudge/deleteQuestionBankJudgeByIds [delete]
func (questionBank_judgeApi *QuestionBankJudgeApi) DeleteQuestionBankJudgeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := questionBank_judgeService.DeleteQuestionBankJudgeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionBankJudge 更新QuestionBankJudge
// @Tags QuestionBankJudge
// @Summary 更新QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.Judge true "更新QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBankJudge/updateQuestionBankJudge [put]
func (questionBank_judgeApi *QuestionBankJudgeApi) UpdateQuestionBankJudge(c *gin.Context) {
	var questionBank_judge questionBank.Judge
	_ = c.ShouldBindJSON(&questionBank_judge)
	verify := utils.Rules{
		"Describe": {utils.NotEmpty()},
		"Is_right": {utils.NotEmpty()},
	}
	if err := utils.Verify(questionBank_judge, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := questionBank_judgeService.UpdateQuestionBankJudge(questionBank_judge); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionBankJudge 用id查询QuestionBankJudge
// @Tags QuestionBankJudge
// @Summary 用id查询QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBank.Judge true "用id查询QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBankJudge/findQuestionBankJudge [get]
func (questionBank_judgeApi *QuestionBankJudgeApi) FindQuestionBankJudge(c *gin.Context) {
	var questionBank_judge questionBank.Judge
	_ = c.ShouldBindQuery(&questionBank_judge)
	if requestionBank_judge, err := questionBank_judgeService.GetQuestionBankJudge(questionBank_judge.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionBank_judge": requestionBank_judge}, c)
	}
}

// GetQuestionBankJudgeList 分页获取QuestionBankJudge列表
// @Tags QuestionBankJudge
// @Summary 分页获取QuestionBankJudge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBankReq.QuestionBankJudgeSearch true "分页获取QuestionBankJudge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankJudge/getQuestionBankJudgeList [get]
func (questionBank_judgeApi *QuestionBankJudgeApi) GetQuestionBankJudgeList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankJudgeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBank_judgeService.GetQuestionBankJudgeInfoList(pageInfo); err != nil {
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
