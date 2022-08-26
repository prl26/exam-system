package questionBank

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/common/response"
	"exam-system/model/questionBank"
	questionBankReq "exam-system/model/questionBank/request"
	"exam-system/service"
	"exam-system/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProgrammLanguageMergeApi struct {
}

var programmLanguageMergeService = service.ServiceGroupApp.QuestionBankServiceGroup.ProgrammLanguageMergeService

// CreateProgrammLanguageMerge 创建ProgrammLanguageMerge
// @Tags ProgrammLanguageMerge
// @Summary 创建ProgrammLanguageMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.ProgrammLanguageMerge true "创建ProgrammLanguageMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /programmLanguageMerge/createProgrammLanguageMerge [post]
func (programmLanguageMergeApi *ProgrammLanguageMergeApi) CreateProgrammLanguageMerge(c *gin.Context) {
	var programmLanguageMerge questionBank.ProgrammLanguageMerge
	_ = c.ShouldBindJSON(&programmLanguageMerge)
	verify := utils.Rules{
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(programmLanguageMerge, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := programmLanguageMergeService.CreateProgrammLanguageMerge(programmLanguageMerge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProgrammLanguageMerge 删除ProgrammLanguageMerge
// @Tags ProgrammLanguageMerge
// @Summary 删除ProgrammLanguageMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.ProgrammLanguageMerge true "删除ProgrammLanguageMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /programmLanguageMerge/deleteProgrammLanguageMerge [delete]
func (programmLanguageMergeApi *ProgrammLanguageMergeApi) DeleteProgrammLanguageMerge(c *gin.Context) {
	var programmLanguageMerge questionBank.ProgrammLanguageMerge
	_ = c.ShouldBindJSON(&programmLanguageMerge)
	if err := programmLanguageMergeService.DeleteProgrammLanguageMerge(programmLanguageMerge); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProgrammLanguageMergeByIds 批量删除ProgrammLanguageMerge
// @Tags ProgrammLanguageMerge
// @Summary 批量删除ProgrammLanguageMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProgrammLanguageMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /programmLanguageMerge/deleteProgrammLanguageMergeByIds [delete]
func (programmLanguageMergeApi *ProgrammLanguageMergeApi) DeleteProgrammLanguageMergeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := programmLanguageMergeService.DeleteProgrammLanguageMergeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProgrammLanguageMerge 更新ProgrammLanguageMerge
// @Tags ProgrammLanguageMerge
// @Summary 更新ProgrammLanguageMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body questionBank.ProgrammLanguageMerge true "更新ProgrammLanguageMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /programmLanguageMerge/updateProgrammLanguageMerge [put]
func (programmLanguageMergeApi *ProgrammLanguageMergeApi) UpdateProgrammLanguageMerge(c *gin.Context) {
	var programmLanguageMerge questionBank.ProgrammLanguageMerge
	_ = c.ShouldBindJSON(&programmLanguageMerge)
	verify := utils.Rules{
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(programmLanguageMerge, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := programmLanguageMergeService.UpdateProgrammLanguageMerge(programmLanguageMerge); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProgrammLanguageMerge 用id查询ProgrammLanguageMerge
// @Tags ProgrammLanguageMerge
// @Summary 用id查询ProgrammLanguageMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query questionBank.ProgrammLanguageMerge true "用id查询ProgrammLanguageMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /programmLanguageMerge/findProgrammLanguageMerge [get]
func (programmLanguageMergeApi *ProgrammLanguageMergeApi) FindProgrammLanguageMerge(c *gin.Context) {
	var programmLanguageMerge questionBank.ProgrammLanguageMerge
	_ = c.ShouldBindQuery(&programmLanguageMerge)
	if reprogrammLanguageMerge, err := programmLanguageMergeService.GetProgrammLanguageMerge(programmLanguageMerge.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprogrammLanguageMerge": reprogrammLanguageMerge}, c)
	}
}

// GetProgrammLanguageMergeList 分页获取ProgrammLanguageMerge列表
// @Tags ProgrammLanguageMerge
// @Summary 分页获取ProgrammLanguageMerge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.ProgrammLanguageMergeSearch true "分页获取ProgrammLanguageMerge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /programmLanguageMerge/getProgrammLanguageMergeList [get]
func (programmLanguageMergeApi *ProgrammLanguageMergeApi) GetProgrammLanguageMergeList(c *gin.Context) {
	var pageInfo questionBankReq.ProgrammLanguageMergeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := programmLanguageMergeService.GetProgrammLanguageMergeInfoList(pageInfo); err != nil {
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
