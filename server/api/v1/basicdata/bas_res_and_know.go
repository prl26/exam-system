package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResandknowApi struct {
}

var resandknowService = service.ServiceGroupApp.BasicdataApiGroup.ResandknowService


// CreateResandknow 创建Resandknow
// @Tags Resandknow
// @Summary 创建Resandknow
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Resandknow true "创建Resandknow"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resandknow/createResandknow [post]
func (resandknowApi *ResandknowApi) CreateResandknow(c *gin.Context) {
	var resandknow basicdata.Resandknow
	_ = c.ShouldBindJSON(&resandknow)
	if err := resandknowService.CreateResandknow(resandknow); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteResandknow 删除Resandknow
// @Tags Resandknow
// @Summary 删除Resandknow
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Resandknow true "删除Resandknow"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resandknow/deleteResandknow [delete]
func (resandknowApi *ResandknowApi) DeleteResandknow(c *gin.Context) {
	var resandknow basicdata.Resandknow
	_ = c.ShouldBindJSON(&resandknow)
	if err := resandknowService.DeleteResandknow(resandknow); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteResandknowByIds 批量删除Resandknow
// @Tags Resandknow
// @Summary 批量删除Resandknow
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Resandknow"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /resandknow/deleteResandknowByIds [delete]
func (resandknowApi *ResandknowApi) DeleteResandknowByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := resandknowService.DeleteResandknowByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateResandknow 更新Resandknow
// @Tags Resandknow
// @Summary 更新Resandknow
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Resandknow true "更新Resandknow"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /resandknow/updateResandknow [put]
func (resandknowApi *ResandknowApi) UpdateResandknow(c *gin.Context) {
	var resandknow basicdata.Resandknow
	_ = c.ShouldBindJSON(&resandknow)
	if err := resandknowService.UpdateResandknow(resandknow); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindResandknow 用id查询Resandknow
// @Tags Resandknow
// @Summary 用id查询Resandknow
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.Resandknow true "用id查询Resandknow"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /resandknow/findResandknow [get]
func (resandknowApi *ResandknowApi) FindResandknow(c *gin.Context) {
	var resandknow basicdata.Resandknow
	_ = c.ShouldBindQuery(&resandknow)
	if reresandknow, err := resandknowService.GetResandknow(resandknow.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reresandknow": reresandknow}, c)
	}
}

// GetResandknowList 分页获取Resandknow列表
// @Tags Resandknow
// @Summary 分页获取Resandknow列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.ResandknowSearch true "分页获取Resandknow列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resandknow/getResandknowList [get]
func (resandknowApi *ResandknowApi) GetResandknowList(c *gin.Context) {
	var pageInfo basicdataReq.ResandknowSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := resandknowService.GetResandknowInfoList(pageInfo); err != nil {
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
