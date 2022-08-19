package teachplan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/teachplan"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    teachplanReq "github.com/flipped-aurora/gin-vue-admin/server/model/teachplan/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ExamPlanApi struct {
}

var examPlanService = service.ServiceGroupApp.TeachplanServiceGroup.ExamPlanService


// CreateExamPlan 创建ExamPlan
// @Tags ExamPlan
// @Summary 创建ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.ExamPlan true "创建ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPlan/createExamPlan [post]
func (examPlanApi *ExamPlanApi) CreateExamPlan(c *gin.Context) {
	var examPlan teachplan.ExamPlan
	_ = c.ShouldBindJSON(&examPlan)
	if err := examPlanService.CreateExamPlan(examPlan); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteExamPlan 删除ExamPlan
// @Tags ExamPlan
// @Summary 删除ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.ExamPlan true "删除ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPlan/deleteExamPlan [delete]
func (examPlanApi *ExamPlanApi) DeleteExamPlan(c *gin.Context) {
	var examPlan teachplan.ExamPlan
	_ = c.ShouldBindJSON(&examPlan)
	if err := examPlanService.DeleteExamPlan(examPlan); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExamPlanByIds 批量删除ExamPlan
// @Tags ExamPlan
// @Summary 批量删除ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /examPlan/deleteExamPlanByIds [delete]
func (examPlanApi *ExamPlanApi) DeleteExamPlanByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := examPlanService.DeleteExamPlanByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExamPlan 更新ExamPlan
// @Tags ExamPlan
// @Summary 更新ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.ExamPlan true "更新ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examPlan/updateExamPlan [put]
func (examPlanApi *ExamPlanApi) UpdateExamPlan(c *gin.Context) {
	var examPlan teachplan.ExamPlan
	_ = c.ShouldBindJSON(&examPlan)
	if err := examPlanService.UpdateExamPlan(examPlan); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExamPlan 用id查询ExamPlan
// @Tags ExamPlan
// @Summary 用id查询ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplan.ExamPlan true "用id查询ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examPlan/findExamPlan [get]
func (examPlanApi *ExamPlanApi) FindExamPlan(c *gin.Context) {
	var examPlan teachplan.ExamPlan
	_ = c.ShouldBindQuery(&examPlan)
	if reexamPlan, err := examPlanService.GetExamPlan(examPlan.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexamPlan": reexamPlan}, c)
	}
}

// GetExamPlanList 分页获取ExamPlan列表
// @Tags ExamPlan
// @Summary 分页获取ExamPlan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplanReq.ExamPlanSearch true "分页获取ExamPlan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPlan/getExamPlanList [get]
func (examPlanApi *ExamPlanApi) GetExamPlanList(c *gin.Context) {
	var pageInfo teachplanReq.ExamPlanSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := examPlanService.GetExamPlanInfoList(pageInfo); err != nil {
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
