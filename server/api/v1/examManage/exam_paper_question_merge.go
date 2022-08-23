package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    examManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/examManage/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PaperQuestionMergeApi struct {
}

var 试卷题目表Service = service.ServiceGroupApp.ExammanageServiceGroup.PaperQuestionMergeService


// CreatePaperQuestionMerge 创建PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 创建PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.PaperQuestionMerge true "创建PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /试卷题目表/createPaperQuestionMerge [post]
func (试卷题目表Api *PaperQuestionMergeApi) CreatePaperQuestionMerge(c *gin.Context) {
	var 试卷题目表 examManage.PaperQuestionMerge
	_ = c.ShouldBindJSON(&试卷题目表)
	if err := 试卷题目表Service.CreatePaperQuestionMerge(试卷题目表); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePaperQuestionMerge 删除PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 删除PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.PaperQuestionMerge true "删除PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /试卷题目表/deletePaperQuestionMerge [delete]
func (试卷题目表Api *PaperQuestionMergeApi) DeletePaperQuestionMerge(c *gin.Context) {
	var 试卷题目表 examManage.PaperQuestionMerge
	_ = c.ShouldBindJSON(&试卷题目表)
	if err := 试卷题目表Service.DeletePaperQuestionMerge(试卷题目表); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePaperQuestionMergeByIds 批量删除PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 批量删除PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /试卷题目表/deletePaperQuestionMergeByIds [delete]
func (试卷题目表Api *PaperQuestionMergeApi) DeletePaperQuestionMergeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := 试卷题目表Service.DeletePaperQuestionMergeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePaperQuestionMerge 更新PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 更新PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.PaperQuestionMerge true "更新PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /试卷题目表/updatePaperQuestionMerge [put]
func (试卷题目表Api *PaperQuestionMergeApi) UpdatePaperQuestionMerge(c *gin.Context) {
	var 试卷题目表 examManage.PaperQuestionMerge
	_ = c.ShouldBindJSON(&试卷题目表)
	if err := 试卷题目表Service.UpdatePaperQuestionMerge(试卷题目表); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPaperQuestionMerge 用id查询PaperQuestionMerge
// @Tags PaperQuestionMerge
// @Summary 用id查询PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManage.PaperQuestionMerge true "用id查询PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /试卷题目表/findPaperQuestionMerge [get]
func (试卷题目表Api *PaperQuestionMergeApi) FindPaperQuestionMerge(c *gin.Context) {
	var 试卷题目表 examManage.PaperQuestionMerge
	_ = c.ShouldBindQuery(&试卷题目表)
	if re试卷题目表, err := 试卷题目表Service.GetPaperQuestionMerge(试卷题目表.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"re试卷题目表": re试卷题目表}, c)
	}
}

// GetPaperQuestionMergeList 分页获取PaperQuestionMerge列表
// @Tags PaperQuestionMerge
// @Summary 分页获取PaperQuestionMerge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManageReq.PaperQuestionMergeSearch true "分页获取PaperQuestionMerge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /试卷题目表/getPaperQuestionMergeList [get]
func (试卷题目表Api *PaperQuestionMergeApi) GetPaperQuestionMergeList(c *gin.Context) {
	var pageInfo examManageReq.PaperQuestionMergeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := 试卷题目表Service.GetPaperQuestionMergeInfoList(pageInfo); err != nil {
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
