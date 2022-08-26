package lessondata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/lessondata"
	lessondataReq "github.com/prl26/exam-system/server/model/lessondata/request"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type ResourcesTestApi struct {
}

var resourcesTestService = service.ServiceGroupApp.LessondataServiceGroup.ResourcesTestService

// CreateResourcesTest 创建ResourcesTest
// @Tags ResourcesTest
// @Summary 创建ResourcesTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ResourcesTest true "创建ResourcesTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resourcesTest/createResourcesTest [post]
func (resourcesTestApi *ResourcesTestApi) CreateResourcesTest(c *gin.Context) {
	var resourcesTest lessondata.ResourcesTest
	_ = c.ShouldBindJSON(&resourcesTest)
	if err := resourcesTestService.CreateResourcesTest(resourcesTest); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteResourcesTest 删除ResourcesTest
// @Tags ResourcesTest
// @Summary 删除ResourcesTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ResourcesTest true "删除ResourcesTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resourcesTest/deleteResourcesTest [delete]
func (resourcesTestApi *ResourcesTestApi) DeleteResourcesTest(c *gin.Context) {
	var resourcesTest lessondata.ResourcesTest
	_ = c.ShouldBindJSON(&resourcesTest)
	if err := resourcesTestService.DeleteResourcesTest(resourcesTest); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteResourcesTestByIds 批量删除ResourcesTest
// @Tags ResourcesTest
// @Summary 批量删除ResourcesTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ResourcesTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /resourcesTest/deleteResourcesTestByIds [delete]
func (resourcesTestApi *ResourcesTestApi) DeleteResourcesTestByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := resourcesTestService.DeleteResourcesTestByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateResourcesTest 更新ResourcesTest
// @Tags ResourcesTest
// @Summary 更新ResourcesTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ResourcesTest true "更新ResourcesTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /resourcesTest/updateResourcesTest [put]
func (resourcesTestApi *ResourcesTestApi) UpdateResourcesTest(c *gin.Context) {
	var resourcesTest lessondata.ResourcesTest
	_ = c.ShouldBindJSON(&resourcesTest)
	if err := resourcesTestService.UpdateResourcesTest(resourcesTest); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindResourcesTest 用id查询ResourcesTest
// @Tags ResourcesTest
// @Summary 用id查询ResourcesTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondata.ResourcesTest true "用id查询ResourcesTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /resourcesTest/findResourcesTest [get]
func (resourcesTestApi *ResourcesTestApi) FindResourcesTest(c *gin.Context) {
	var resourcesTest lessondata.ResourcesTest
	_ = c.ShouldBindQuery(&resourcesTest)
	if reresourcesTest, err := resourcesTestService.GetResourcesTest(resourcesTest.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reresourcesTest": reresourcesTest}, c)
	}
}

// GetResourcesTestList 分页获取ResourcesTest列表
// @Tags ResourcesTest
// @Summary 分页获取ResourcesTest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondataReq.ResourcesTestSearch true "分页获取ResourcesTest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resourcesTest/getResourcesTestList [get]
func (resourcesTestApi *ResourcesTestApi) GetResourcesTestList(c *gin.Context) {
	var pageInfo lessondataReq.ResourcesTestSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := resourcesTestService.GetResourcesTestInfoList(pageInfo); err != nil {
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
