package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ResourceApi struct {
}

var resourceService = service.ServiceGroupApp.BasicdataApiGroup.ResourceService


// CreateResource 创建Resource
// @Tags Resource
// @Summary 创建Resource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Resource true "创建Resource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/createResource [post]
func (resourceApi *ResourceApi) CreateResource(c *gin.Context) {
	var resource basicdata.Resource
	_ = c.ShouldBindJSON(&resource)
	if err := resourceService.CreateResource(resource); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteResource 删除Resource
// @Tags Resource
// @Summary 删除Resource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Resource true "删除Resource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resource/deleteResource [delete]
func (resourceApi *ResourceApi) DeleteResource(c *gin.Context) {
	var resource basicdata.Resource
	_ = c.ShouldBindJSON(&resource)
	if err := resourceService.DeleteResource(resource); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteResourceByIds 批量删除Resource
// @Tags Resource
// @Summary 批量删除Resource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Resource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /resource/deleteResourceByIds [delete]
func (resourceApi *ResourceApi) DeleteResourceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := resourceService.DeleteResourceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateResource 更新Resource
// @Tags Resource
// @Summary 更新Resource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Resource true "更新Resource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /resource/updateResource [put]
func (resourceApi *ResourceApi) UpdateResource(c *gin.Context) {
	var resource basicdata.Resource
	_ = c.ShouldBindJSON(&resource)
	if err := resourceService.UpdateResource(resource); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindResource 用id查询Resource
// @Tags Resource
// @Summary 用id查询Resource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.Resource true "用id查询Resource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /resource/findResource [get]
func (resourceApi *ResourceApi) FindResource(c *gin.Context) {
	var resource basicdata.Resource
	_ = c.ShouldBindQuery(&resource)
	if reresource, err := resourceService.GetResource(resource.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reresource": reresource}, c)
	}
}

// GetResourceList 分页获取Resource列表
// @Tags Resource
// @Summary 分页获取Resource列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.ResourceSearch true "分页获取Resource列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/getResourceList [get]
func (resourceApi *ResourceApi) GetResourceList(c *gin.Context) {
	var pageInfo basicdataReq.ResourceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := resourceService.GetResourceInfoList(pageInfo); err != nil {
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
