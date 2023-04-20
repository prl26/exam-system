package teachplan

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
	teachplanResp "github.com/prl26/exam-system/server/model/teachplan/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"net/url"
	"time"
)

type TeachAttendanceApi struct {
}

var teachAttendanceService = service.ServiceGroupApp.TeachplanServiceGroup.TeachAttendanceService
var MultiTableService = service.ServiceGroupApp.BasicdataApiGroup.MultiTableService

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

	//var teachClassStudent = request2.TeachClassStudent{
	//	TeachClassId: uint(a),
	//	PageInfo: request.PageInfo{
	//		Page:     1,
	//		PageSize: 1000,
	//	},
	//}
	students, err := MultiTableService.FindStudentByStudentClassId(teachAttendance.TeachClassId)
	if err != nil {
		response.FailWithMessage("查询失败", c)
	}
	if err := teachAttendanceService.CreateTeachAttendance(teachAttendance, students); err != nil {
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
	var teachAttendance teachplanReq.AttendanceDetail
	_ = c.ShouldBindQuery(&teachAttendance)
	if list, total, doneTotal, err := teachAttendanceService.GetTeachAttendance(teachAttendance.AttendanceId, teachAttendance.PageInfo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"List":      list,
			"Total":     total,
			"Page":      teachAttendance.Page,
			"PageSize":  teachAttendance.PageSize,
			"doneTotal": doneTotal,
		}, "获取成功", c)
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

func (TeachAttendanceApi *TeachAttendanceApi) Supplement(c *gin.Context) {
	var req teachplanReq.Supplement
	_ = c.ShouldBindJSON(&req)
	ip, _ := c.RemoteIP()
	if n, err := teachAttendanceService.Attendance(req.StudentId, ip.String(), req.AttendanceId, 2); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		if n == 0 {
			response.CheckHandle(c, fmt.Errorf("该学员不存在于该考勤当中"))
			return
		}
		response.OkWithMessage("补签成功", c)
		return
	}

}

func (TeachAttendanceApi *TeachAttendanceApi) GenerateQRCode(c *gin.Context) {
	var req teachplanReq.GenerateQrCode
	_ = c.ShouldBindJSON(&req)
	t := time.Now().Add(time.Duration(req.Minute) * time.Minute)
	timeStr := utils.TimeToString(t)
	str := fmt.Sprintf("%d,%s", req.AttendanceId, timeStr)
	key := utils.Crypto(str)
	key = fmt.Sprintf("%s/check-in?code=%s", global.GVA_CONFIG.FrontDeskAddress, key)

	key = url.QueryEscape(key)

	code := utils.GenerateQRCode(key)
	response.OkWithData(teachplanResp.GenerateQRCode{
		QRCodeURL:  code,
		ExpireTime: timeStr,
		Minute:     req.Minute,
	}, c)
}
