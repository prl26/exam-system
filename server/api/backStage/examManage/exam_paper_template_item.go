package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type PaperTemplateItemApi struct {
}

var paperTemplateItemService = service.ServiceGroupApp.ExammanageServiceGroup.PaperTemplateItemService

// CreatePaperTemplateItem 创建PaperTemplateItem
// @Tags PaperTemplateItem
// @Summary 创建PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.PaperTemplateItem true "创建PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paperTemplateItem/createPaperTemplateItem [post]
func (paperTemplateItemApi *PaperTemplateItemApi) CreatePaperTemplateItem(c *gin.Context) {
	var paperTemplateItem examManage.PaperTemplateItem
	_ = c.ShouldBindJSON(&paperTemplateItem)
	if err := paperTemplateItemService.CreatePaperTemplateItem(paperTemplateItem); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePaperTemplateItem 删除PaperTemplateItem
// @Tags PaperTemplateItem
// @Summary 删除PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.PaperTemplateItem true "删除PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /paperTemplateItem/deletePaperTemplateItem [delete]
func (paperTemplateItemApi *PaperTemplateItemApi) DeletePaperTemplateItem(c *gin.Context) {
	var paperTemplateItem examManage.PaperTemplateItem
	_ = c.ShouldBindJSON(&paperTemplateItem)
	if err := paperTemplateItemService.DeletePaperTemplateItem(paperTemplateItem); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePaperTemplateItemByIds 批量删除PaperTemplateItem
// @Tags PaperTemplateItem
// @Summary 批量删除PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /paperTemplateItem/deletePaperTemplateItemByIds [delete]
func (paperTemplateItemApi *PaperTemplateItemApi) DeletePaperTemplateItemByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := paperTemplateItemService.DeletePaperTemplateItemByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePaperTemplateItem 更新PaperTemplateItem
// @Tags PaperTemplateItem
// @Summary 更新PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.PaperTemplateItem true "更新PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /paperTemplateItem/updatePaperTemplateItem [put]
func (paperTemplateItemApi *PaperTemplateItemApi) UpdatePaperTemplateItem(c *gin.Context) {
	var paperTemplateItem []examManage.PaperTemplateItem
	_ = c.ShouldBindJSON(&paperTemplateItem)
	if err := paperTemplateItemService.UpdatePaperTemplateItem(paperTemplateItem); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPaperTemplateItem 用id查询PaperTemplateItem
// @Tags PaperTemplateItem
// @Summary 用id查询PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query frontExamManage.PaperTemplateItem true "用id查询PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /paperTemplateItem/findPaperTemplateItem [get]
func (paperTemplateItemApi *PaperTemplateItemApi) FindPaperTemplateItem(c *gin.Context) {
	var paperTemplateItem examManage.PaperTemplateItem
	_ = c.ShouldBindQuery(&paperTemplateItem)
	if repaperTemplateItem, err := paperTemplateItemService.GetPaperTemplateItem(paperTemplateItem.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repaperTemplateItem": repaperTemplateItem}, c)
	}
}

// GetPaperTemplateItemList 分页获取PaperTemplateItem列表
// @Tags PaperTemplateItem
// @Summary 分页获取PaperTemplateItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManageReq.PaperTemplateItemSearch true "分页获取PaperTemplateItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paperTemplateItem/getPaperTemplateItemList [get]
func (paperTemplateItemApi *PaperTemplateItemApi) GetPaperTemplateItemList(c *gin.Context) {
	var pageInfo examManageReq.PaperTemplateItemSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := paperTemplateItemService.GetPaperTemplateItemInfoList(pageInfo); err != nil {
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
