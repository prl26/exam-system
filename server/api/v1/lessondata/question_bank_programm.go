package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lessondata"
	lessondataReq "github.com/flipped-aurora/gin-vue-admin/server/model/lessondata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type QuestionBankProgrammApi struct {
}

var questionBankProgrammService = service.ServiceGroupApp.LessondataServiceGroup.QuestionBankProgrammService

// CreateQuestionBankProgramm 创建QuestionBankProgramm
// @Tags QuestionBankProgramm
// @Summary 创建QuestionBankProgramm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.QuestionBankProgramm true "创建QuestionBankProgramm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankProgramm/createQuestionBankProgramm [post]
func (questionBankProgrammApi *QuestionBankProgrammApi) CreateQuestionBankProgramm(c *gin.Context) {
	var questionBankProgramm lessondata.QuestionBankProgramm
	_ = c.ShouldBindJSON(&questionBankProgramm)
	if err := questionBankProgrammService.CreateQuestionBankProgramm(questionBankProgramm); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionBankProgramm 删除QuestionBankProgramm
// @Tags QuestionBankProgramm
// @Summary 删除QuestionBankProgramm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.QuestionBankProgramm true "删除QuestionBankProgramm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBankProgramm/deleteQuestionBankProgramm [delete]
func (questionBankProgrammApi *QuestionBankProgrammApi) DeleteQuestionBankProgramm(c *gin.Context) {
	var questionBankProgramm lessondata.QuestionBankProgramm
	_ = c.ShouldBindJSON(&questionBankProgramm)
	if err := questionBankProgrammService.DeleteQuestionBankProgramm(questionBankProgramm); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionBankProgrammByIds 批量删除QuestionBankProgramm
// @Tags QuestionBankProgramm
// @Summary 批量删除QuestionBankProgramm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankProgramm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionBankProgramm/deleteQuestionBankProgrammByIds [delete]
func (questionBankProgrammApi *QuestionBankProgrammApi) DeleteQuestionBankProgrammByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := questionBankProgrammService.DeleteQuestionBankProgrammByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionBankProgramm 更新QuestionBankProgramm
// @Tags QuestionBankProgramm
// @Summary 更新QuestionBankProgramm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.QuestionBankProgramm true "更新QuestionBankProgramm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBankProgramm/updateQuestionBankProgramm [put]
func (questionBankProgrammApi *QuestionBankProgrammApi) UpdateQuestionBankProgramm(c *gin.Context) {
	var questionBankProgramm lessondata.QuestionBankProgramm
	_ = c.ShouldBindJSON(&questionBankProgramm)
	if err := questionBankProgrammService.UpdateQuestionBankProgramm(questionBankProgramm); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionBankProgramm 用id查询QuestionBankProgramm
// @Tags QuestionBankProgramm
// @Summary 用id查询QuestionBankProgramm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondata.QuestionBankProgramm true "用id查询QuestionBankProgramm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBankProgramm/findQuestionBankProgramm [get]
func (questionBankProgrammApi *QuestionBankProgrammApi) FindQuestionBankProgramm(c *gin.Context) {
	var questionBankProgramm lessondata.QuestionBankProgramm
	_ = c.ShouldBindQuery(&questionBankProgramm)
	if requestionBankProgramm, err := questionBankProgrammService.GetQuestionBankProgramm(questionBankProgramm.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionBankProgramm": requestionBankProgramm}, c)
	}
}

// GetQuestionBankProgrammList 分页获取QuestionBankProgramm列表
// @Tags QuestionBankProgramm
// @Summary 分页获取QuestionBankProgramm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondataReq.QuestionBankProgrammSearch true "分页获取QuestionBankProgramm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankProgramm/getQuestionBankProgrammList [get]
func (questionBankProgrammApi *QuestionBankProgrammApi) GetQuestionBankProgrammList(c *gin.Context) {
	var pageInfo lessondataReq.QuestionBankProgrammSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := questionBankProgrammService.GetQuestionBankProgrammInfoList(pageInfo); err != nil {
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
