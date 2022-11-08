package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type ProfessionalApi struct {
}

var professionalService = service.ServiceGroupApp.BasicdataApiGroup.ProfessionalService

// CreateProfessional 创建Professional
// @Tags Professional
// @Summary 创建Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Professional true "创建Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /professional/createProfessional [post]
func (professionalApi *ProfessionalApi) CreateProfessional(c *gin.Context) {
	var professional basicdata.Professional
	_ = c.ShouldBindJSON(&professional)
	verify := utils.Rules{
		"Name":      {utils.NotEmpty()},
		"CollegeId": {utils.NotEmpty()},
	}
	if err := utils.Verify(professional, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := professionalService.CreateProfessional(professional); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProfessional 删除Professional
// @Tags Professional
// @Summary 删除Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Professional true "删除Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /professional/deleteProfessional [delete]
func (professionalApi *ProfessionalApi) DeleteProfessional(c *gin.Context) {
	var professional basicdata.Professional
	_ = c.ShouldBindJSON(&professional)
	if err := professionalService.DeleteProfessional(professional); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProfessionalByIds 批量删除Professional
// @Tags Professional
// @Summary 批量删除Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /professional/deleteProfessionalByIds [delete]
func (professionalApi *ProfessionalApi) DeleteProfessionalByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := professionalService.DeleteProfessionalByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProfessional 更新Professional
// @Tags Professional
// @Summary 更新Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Professional true "更新Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /professional/updateProfessional [put]
func (professionalApi *ProfessionalApi) UpdateProfessional(c *gin.Context) {
	var professional basicdata.Professional
	_ = c.ShouldBindJSON(&professional)
	verify := utils.Rules{
		"Name":      {utils.NotEmpty()},
		"CollegeId": {utils.NotEmpty()},
	}
	if err := utils.Verify(professional, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := professionalService.UpdateProfessional(professional); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProfessional 用id查询Professional
// @Tags Professional
// @Summary 用id查询Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.Professional true "用id查询Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /professional/findProfessional [get]
func (professionalApi *ProfessionalApi) FindProfessional(c *gin.Context) {
	var professional basicdata.Professional
	_ = c.ShouldBindQuery(&professional)
	if reprofessional, err := professionalService.GetProfessional(professional.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprofessional": reprofessional}, c)
	}
}

// GetProfessionalList 分页获取Professional列表
// @Tags Professional
// @Summary 分页获取Professional列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.ProfessionalSearch true "分页获取Professional列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /professional/getProfessionalList [get]
func (professionalApi *ProfessionalApi) GetProfessionalList(c *gin.Context) {
	var pageInfo basicdataReq.ProfessionalSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := professionalService.GetProfessionalInfoList(pageInfo); err != nil {
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
