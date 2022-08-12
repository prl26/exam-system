
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

type VideoResourcesApi struct {
}

var videoResourcesService = service.ServiceGroupApp.LessondataServiceGroup.VideoResourcesService


// CreateVideoResources 创建VideoResources
// @Tags VideoResources
// @Summary 创建VideoResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.VideoResources true "创建VideoResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoResources/createVideoResources [post]
func (videoResourcesApi *VideoResourcesApi) CreateVideoResources(c *gin.Context) {
	var videoResources lessondata.VideoResources
	_ = c.ShouldBindJSON(&videoResources)
	if err := videoResourcesService.CreateVideoResources(videoResources); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVideoResources 删除VideoResources
// @Tags VideoResources
// @Summary 删除VideoResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.VideoResources true "删除VideoResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoResources/deleteVideoResources [delete]
func (videoResourcesApi *VideoResourcesApi) DeleteVideoResources(c *gin.Context) {
	var videoResources lessondata.VideoResources
	_ = c.ShouldBindJSON(&videoResources)
	if err := videoResourcesService.DeleteVideoResources(videoResources); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVideoResourcesByIds 批量删除VideoResources
// @Tags VideoResources
// @Summary 批量删除VideoResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /videoResources/deleteVideoResourcesByIds [delete]
func (videoResourcesApi *VideoResourcesApi) DeleteVideoResourcesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := videoResourcesService.DeleteVideoResourcesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVideoResources 更新VideoResources
// @Tags VideoResources
// @Summary 更新VideoResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.VideoResources true "更新VideoResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoResources/updateVideoResources [put]
func (videoResourcesApi *VideoResourcesApi) UpdateVideoResources(c *gin.Context) {
	var videoResources lessondata.VideoResources
	_ = c.ShouldBindJSON(&videoResources)
	if err := videoResourcesService.UpdateVideoResources(videoResources); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVideoResources 用id查询VideoResources
// @Tags VideoResources
// @Summary 用id查询VideoResources-

// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondata.VideoResources true "用id查询VideoResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoResources/findVideoResources [get]
func (videoResourcesApi *VideoResourcesApi) FindVideoResources(c *gin.Context) {
	var videoResources lessondata.VideoResources
	_ = c.ShouldBindQuery(&videoResources)
	if revideoResources, err := videoResourcesService.GetVideoResources(videoResources.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideoResources": revideoResources}, c)
	}
}

// GetVideoResourcesList 分页获取VideoResources列表
// @Tags VideoResources
// @Summary 分页获取VideoResources列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondataReq.VideoResourcesSearch true "分页获取VideoResources列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoResources/getVideoResourcesList [get]
func (videoResourcesApi *VideoResourcesApi) GetVideoResourcesList(c *gin.Context) {
	var pageInfo lessondataReq.VideoResourcesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := videoResourcesService.GetVideoResourcesInfoList(pageInfo); err != nil {
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
