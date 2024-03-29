package teachplan

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type ExamPlanApi struct {
}

var examPlanService = service.ServiceGroupApp.TeachplanServiceGroup.ExamPlanService
var lessonService = service.ServiceGroupApp.BasicdataApiGroup.LessonService

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
	var examPlan teachplanReq.ExamPlanRq
	_ = c.ShouldBindJSON(&examPlan)
	fmt.Println(examPlan)
	verify := utils.Rules{
		"Name":         {utils.NotEmpty()},
		"TeachClassId": {utils.NotEmpty()},
		"StartTime":    {utils.NotEmpty()},
		"Time":         {utils.NotEmpty()},
		"EndTime":      {utils.NotEmpty()},
		"LessonId":     {utils.NotEmpty()},
		"Type":         {utils.NotEmpty()},
		"PassScore":    {utils.NotEmpty()},
		"TermId":       {utils.NotEmpty()},
		"Weight":       {utils.NotEmpty()},
	}
	if err := utils.Verify(examPlan, verify); err != nil {
		response.CheckHandle(c, err)
		return
	}
	userId := utils.GetUserID(c)
	if examPlan.Type == 1 && examPlan.Weight != 100 {
		response.FailWithMessage("期末考试比重应为100%", c)
	} else {
		if err := examPlanService.CreateExamPlan(examPlan, userId); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	}
}

//设置前置计划
func (examPlanApi *ExamPlanApi) ChoosePrePlan(c *gin.Context) {
	var IDS request.PrePlanReq
	_ = c.ShouldBindJSON(&IDS)
	if err := examPlanService.UpdatePrePlan(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
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
	var examPlan teachplanReq.ExamPlanRq1
	_ = c.ShouldBindJSON(&examPlan)
	fmt.Println(examPlan.IsLimitTime)
	if err := examPlanService.UpdateExamPlan(examPlan); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExamPlanById 用id查询ExamPlan
// @Tags ExamPlan
// @Summary 用id查询ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplan.ExamPlan true "用id查询ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examPlan/findExamPlan [get]
func (examPlanApi *ExamPlanApi) FindExamPlanById(c *gin.Context) {
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
	userId := utils.GetUserID(c)
	authorityId := utils.GetUserAuthorityID(c)
	if list, total, err := examPlanService.GetExamPlanInfoList(pageInfo, userId, authorityId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		result := examPlanService.GetExamPlanDetail(list)
		response.OkWithDetailed(response.PageResult{
			List:     result,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
func (examPlanApi *ExamPlanApi) ChangeAudit(c *gin.Context) {
	var planId teachplanReq.ExamPlanAudit
	_ = c.ShouldBindQuery(&planId)
	if err := examPlanService.ChangeAudit(planId.PlanId, planId.Value); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
func (examPlanApi *ExamPlanApi) ChangeStatus(c *gin.Context) {
	var planId teachplanReq.ExamPlan
	_ = c.ShouldBindQuery(&planId)
	if err := examPlanService.ChangeStatus(planId.PlanId); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
