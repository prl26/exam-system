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

type TeachAttendanceRecordApi struct {
}

var teachAttendanceRecordService = service.ServiceGroupApp.TeachplanServiceGroup.TeachAttendanceRecordService

// CreateTeachAttendanceRecord 创建TeachAttendanceRecord
// @Tags TeachAttendanceRecord
// @Summary 创建TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.TeachAttendanceRecord true "创建TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachAttendanceRecord/createTeachAttendanceRecord [post]
func (teachAttendanceRecordApi *TeachAttendanceRecordApi) CreateTeachAttendanceRecord(c *gin.Context) {
	var teachAttendanceRecord teachplan.TeachAttendanceRecord
	_ = c.ShouldBindJSON(&teachAttendanceRecord)
	if err := teachAttendanceRecordService.CreateTeachAttendanceRecord(teachAttendanceRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeachAttendanceRecord 删除TeachAttendanceRecord
// @Tags TeachAttendanceRecord
// @Summary 删除TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.TeachAttendanceRecord true "删除TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachAttendanceRecord/deleteTeachAttendanceRecord [delete]
func (teachAttendanceRecordApi *TeachAttendanceRecordApi) DeleteTeachAttendanceRecord(c *gin.Context) {
	var teachAttendanceRecord teachplan.TeachAttendanceRecord
	_ = c.ShouldBindJSON(&teachAttendanceRecord)
	if err := teachAttendanceRecordService.DeleteTeachAttendanceRecord(teachAttendanceRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeachAttendanceRecordByIds 批量删除TeachAttendanceRecord
// @Tags TeachAttendanceRecord
// @Summary 批量删除TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teachAttendanceRecord/deleteTeachAttendanceRecordByIds [delete]
func (teachAttendanceRecordApi *TeachAttendanceRecordApi) DeleteTeachAttendanceRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := teachAttendanceRecordService.DeleteTeachAttendanceRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeachAttendanceRecord 更新TeachAttendanceRecord
// @Tags TeachAttendanceRecord
// @Summary 更新TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body teachplan.TeachAttendanceRecord true "更新TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teachAttendanceRecord/updateTeachAttendanceRecord [put]
func (teachAttendanceRecordApi *TeachAttendanceRecordApi) UpdateTeachAttendanceRecord(c *gin.Context) {
	var teachAttendanceRecord teachplan.TeachAttendanceRecord
	_ = c.ShouldBindJSON(&teachAttendanceRecord)
	if err := teachAttendanceRecordService.UpdateTeachAttendanceRecord(teachAttendanceRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeachAttendanceRecord 用id查询TeachAttendanceRecord
// @Tags TeachAttendanceRecord
// @Summary 用id查询TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplan.TeachAttendanceRecord true "用id查询TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teachAttendanceRecord/findTeachAttendanceRecord [get]
func (teachAttendanceRecordApi *TeachAttendanceRecordApi) FindTeachAttendanceRecord(c *gin.Context) {
	var teachAttendanceRecord teachplan.TeachAttendanceRecord
	_ = c.ShouldBindQuery(&teachAttendanceRecord)
	if reteachAttendanceRecord, err := teachAttendanceRecordService.GetTeachAttendanceRecord(teachAttendanceRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteachAttendanceRecord": reteachAttendanceRecord}, c)
	}
}

// GetTeachAttendanceRecordList 分页获取TeachAttendanceRecord列表
// @Tags TeachAttendanceRecord
// @Summary 分页获取TeachAttendanceRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query teachplanReq.TeachAttendanceRecordSearch true "分页获取TeachAttendanceRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachAttendanceRecord/getTeachAttendanceRecordList [get]
func (teachAttendanceRecordApi *TeachAttendanceRecordApi) GetTeachAttendanceRecordList(c *gin.Context) {
	var pageInfo teachplanReq.TeachAttendanceRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := teachAttendanceRecordService.GetTeachAttendanceRecordInfoList(pageInfo); err != nil {
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
