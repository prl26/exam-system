package teachplan

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/common/response"
	"exam-system/model/teachplan"
	teachplanReq "exam-system/model/teachplan/request"
	"exam-system/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TeachAttendanceApi struct {
}

var teachAttendanceService = service.ServiceGroupApp.TeachplanServiceGroup.TeachAttendanceService

// CreateTeachAttendance 创建TeachAttendance
// @Tags TeachAttendance
// @Summary 创建TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.TeachAttendance true "创建TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachAttendance/createTeachAttendance [post]
func (teachAttendanceApi *TeachAttendanceApi) CreateTeachAttendance(c *gin.Context) {
	var teachAttendance teachplan.TeachAttendance
	_ = c.ShouldBindJSON(&teachAttendance)
	if err := teachAttendanceService.CreateTeachAttendance(teachAttendance); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeachAttendance 删除TeachAttendance
// @Tags TeachAttendance
// @Summary 删除TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.TeachAttendance true "删除TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachAttendance/deleteTeachAttendance [delete]
func (teachAttendanceApi *TeachAttendanceApi) DeleteTeachAttendance(c *gin.Context) {
	var teachAttendance teachplan.TeachAttendance
	_ = c.ShouldBindJSON(&teachAttendance)
	if err := teachAttendanceService.DeleteTeachAttendance(teachAttendance); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeachAttendanceByIds 批量删除TeachAttendance
// @Tags TeachAttendance
// @Summary 批量删除TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teachAttendance/deleteTeachAttendanceByIds [delete]
func (teachAttendanceApi *TeachAttendanceApi) DeleteTeachAttendanceByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := teachAttendanceService.DeleteTeachAttendanceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeachAttendance 更新TeachAttendance
// @Tags TeachAttendance
// @Summary 更新TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.TeachAttendance true "更新TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teachAttendance/updateTeachAttendance [put]
func (teachAttendanceApi *TeachAttendanceApi) UpdateTeachAttendance(c *gin.Context) {
	var teachAttendance teachplan.TeachAttendance
	_ = c.ShouldBindJSON(&teachAttendance)
	if err := teachAttendanceService.UpdateTeachAttendance(teachAttendance); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeachAttendance 用id查询TeachAttendance
// @Tags TeachAttendance
// @Summary 用id查询TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplan.TeachAttendance true "用id查询TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teachAttendance/findTeachAttendance [get]
func (teachAttendanceApi *TeachAttendanceApi) FindTeachAttendance(c *gin.Context) {
	var teachAttendance teachplan.TeachAttendance
	_ = c.ShouldBindQuery(&teachAttendance)
	if reteachAttendance, err := teachAttendanceService.GetTeachAttendance(teachAttendance.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteachAttendance": reteachAttendance}, c)
	}
}

// GetTeachAttendanceList 分页获取TeachAttendance列表
// @Tags TeachAttendance
// @Summary 分页获取TeachAttendance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplanReq.TeachAttendanceSearch true "分页获取TeachAttendance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachAttendance/getTeachAttendanceList [get]
func (teachAttendanceApi *TeachAttendanceApi) GetTeachAttendanceList(c *gin.Context) {
	var pageInfo teachplanReq.TeachAttendanceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := teachAttendanceService.GetTeachAttendanceInfoList(pageInfo); err != nil {
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
