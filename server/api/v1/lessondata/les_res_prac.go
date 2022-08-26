package lessondata

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/common/response"
	"exam-system/model/lessondata"
	lessondataReq "exam-system/model/lessondata/request"
	"exam-system/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResourcePracticeApi struct {
}

var resourcesPracticeService = service.ServiceGroupApp.LessondataServiceGroup.ResourcePracticeService

// CreateResourcePractice 创建ResourcePractice
// @Tags ResourcePractice
// @Summary 创建ResourcePractice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ResourcePractice true "创建ResourcePractice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resourcesPractice/createResourcePractice [post]
func (resourcesPracticeApi *ResourcePracticeApi) CreateResourcePractice(c *gin.Context) {
	var resourcesPractice lessondata.ResourcePractice
	_ = c.ShouldBindJSON(&resourcesPractice)
	if err := resourcesPracticeService.CreateResourcePractice(resourcesPractice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteResourcePractice 删除ResourcePractice
// @Tags ResourcePractice
// @Summary 删除ResourcePractice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ResourcePractice true "删除ResourcePractice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resourcesPractice/deleteResourcePractice [delete]
func (resourcesPracticeApi *ResourcePracticeApi) DeleteResourcePractice(c *gin.Context) {
	var resourcesPractice lessondata.ResourcePractice
	_ = c.ShouldBindJSON(&resourcesPractice)
	if err := resourcesPracticeService.DeleteResourcePractice(resourcesPractice); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteResourcePracticeByIds 批量删除ResourcePractice
// @Tags ResourcePractice
// @Summary 批量删除ResourcePractice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ResourcePractice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /resourcesPractice/deleteResourcePracticeByIds [delete]
func (resourcesPracticeApi *ResourcePracticeApi) DeleteResourcePracticeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := resourcesPracticeService.DeleteResourcePracticeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateResourcePractice 更新ResourcePractice
// @Tags ResourcePractice
// @Summary 更新ResourcePractice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ResourcePractice true "更新ResourcePractice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /resourcesPractice/updateResourcePractice [put]
func (resourcesPracticeApi *ResourcePracticeApi) UpdateResourcePractice(c *gin.Context) {
	var resourcesPractice lessondata.ResourcePractice
	_ = c.ShouldBindJSON(&resourcesPractice)
	if err := resourcesPracticeService.UpdateResourcePractice(resourcesPractice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindResourcePractice 用id查询ResourcePractice
// @Tags ResourcePractice
// @Summary 用id查询ResourcePractice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondata.ResourcePractice true "用id查询ResourcePractice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /resourcesPractice/findResourcePractice [get]
func (resourcesPracticeApi *ResourcePracticeApi) FindResourcePractice(c *gin.Context) {
	var resourcesPractice lessondata.ResourcePractice
	_ = c.ShouldBindQuery(&resourcesPractice)
	if reresourcesPractice, err := resourcesPracticeService.GetResourcePractice(resourcesPractice.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reresourcesPractice": reresourcesPractice}, c)
	}
}

// GetResourcePracticeList 分页获取ResourcePractice列表
// @Tags ResourcePractice
// @Summary 分页获取ResourcePractice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondataReq.ResourcePracticeSearch true "分页获取ResourcePractice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resourcesPractice/getResourcePracticeList [get]
func (resourcesPracticeApi *ResourcePracticeApi) GetResourcePracticeList(c *gin.Context) {
	var pageInfo lessondataReq.ResourcePracticeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := resourcesPracticeService.GetResourcePracticeInfoList(pageInfo); err != nil {
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
