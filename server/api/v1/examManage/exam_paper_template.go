package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	examManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/examManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExamPaperTemplateApi struct {
}

var examPaperTemplateService = service.ServiceGroupApp.ExammanageServiceGroup.ExamPaperTemplateService

// CreateExamPaperTemplate 创建ExamPaperTemplate
// @Tags ExamPaperTemplate
// @Summary 创建ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.ExamPaperTemplate true "创建ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaperTemplate/createExamPaperTemplate [post]
func (examPaperTemplateApi *ExamPaperTemplateApi) CreateExamPaperTemplate(c *gin.Context) {
	var examPaperTemplate examManage.ExamPaperTemplate
	_ = c.ShouldBindJSON(&examPaperTemplate)
	if err := examPaperTemplateService.CreateExamPaperTemplate(examPaperTemplate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteExamPaperTemplate 删除ExamPaperTemplate
// @Tags ExamPaperTemplate
// @Summary 删除ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.ExamPaperTemplate true "删除ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPaperTemplate/deleteExamPaperTemplate [delete]
func (examPaperTemplateApi *ExamPaperTemplateApi) DeleteExamPaperTemplate(c *gin.Context) {
	var examPaperTemplate examManage.ExamPaperTemplate
	_ = c.ShouldBindJSON(&examPaperTemplate)
	if err := examPaperTemplateService.DeleteExamPaperTemplate(examPaperTemplate); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExamPaperTemplateByIds 批量删除ExamPaperTemplate
// @Tags ExamPaperTemplate
// @Summary 批量删除ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /examPaperTemplate/deleteExamPaperTemplateByIds [delete]
func (examPaperTemplateApi *ExamPaperTemplateApi) DeleteExamPaperTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := examPaperTemplateService.DeleteExamPaperTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExamPaperTemplate 更新ExamPaperTemplate
// @Tags ExamPaperTemplate
// @Summary 更新ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.ExamPaperTemplate true "更新ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examPaperTemplate/updateExamPaperTemplate [put]
func (examPaperTemplateApi *ExamPaperTemplateApi) UpdateExamPaperTemplate(c *gin.Context) {
	var examPaperTemplate examManage.ExamPaperTemplate
	_ = c.ShouldBindJSON(&examPaperTemplate)
	if err := examPaperTemplateService.UpdateExamPaperTemplate(examPaperTemplate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExamPaperTemplate 用id查询ExamPaperTemplate
// @Tags ExamPaperTemplate
// @Summary 用id查询ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManage.ExamPaperTemplate true "用id查询ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examPaperTemplate/findExamPaperTemplate [get]
func (examPaperTemplateApi *ExamPaperTemplateApi) FindExamPaperTemplate(c *gin.Context) {
	var examPaperTemplate examManage.ExamPaperTemplate
	_ = c.ShouldBindQuery(&examPaperTemplate)
	if reexamPaperTemplate, err := examPaperTemplateService.GetExamPaperTemplate(examPaperTemplate.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexamPaperTemplate": reexamPaperTemplate}, c)
	}
}

// GetExamPaperTemplateList 分页获取ExamPaperTemplate列表
// @Tags ExamPaperTemplate
// @Summary 分页获取ExamPaperTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManageReq.ExamPaperTemplateSearch true "分页获取ExamPaperTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaperTemplate/getExamPaperTemplateList [get]
func (examPaperTemplateApi *ExamPaperTemplateApi) GetExamPaperTemplateList(c *gin.Context) {
	var pageInfo examManageReq.ExamPaperTemplateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := examPaperTemplateService.GetExamPaperTemplateInfoList(pageInfo); err != nil {
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
