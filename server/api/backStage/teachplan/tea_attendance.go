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
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"strconv"
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
	key = fmt.Sprintf("http://%s/check-in?code=%s", global.GVA_CONFIG.FrontDeskAddress, key)

	//key = url.QueryEscape(key)

	code := utils.GenerateQRCode(key)
	response.OkWithData(teachplanResp.GenerateQRCode{
		QRCodeURL:  code,
		ExpireTime: timeStr,
		Minute:     req.Minute,
	}, c)
}

func (a *TeachAttendanceApi) GetAttendanceExcel(c *gin.Context) {
	teachClassId, _ := strconv.Atoi(c.Query("teachClassId"))
	attendance, err := teachAttendanceService.GetAllTeachAttendances(uint(teachClassId))
	if err != nil {
		global.GVA_LOG.Error("获取全部的考勤情况失败" + err.Error())
		return
	}
	simple, err := MultiTableService.GetAllTeachClassStudentSimple(teachClassId)
	if err != nil {
		global.GVA_LOG.Error("导出考勤记录: 获取教学计划的所以学生失败" + err.Error())
		return
	}

	zongtiqingkuang := "总体情况"
	xiangxiqingkuang := "详细情况"
	excel := a.attendanceExcelBase(zongtiqingkuang, xiangxiqingkuang)

	for i := 0; i < len(simple); i++ {
		studentId := simple[i].ID
		studentName := simple[i].Name
		excel.SetCellInt(xiangxiqingkuang, fmt.Sprintf("A%d", i+2), int(studentId))
		excel.SetCellStr(xiangxiqingkuang, fmt.Sprintf("B%d", i+2), studentName)
	}

	for i := 0; i < len(attendance); i++ {
		attendanceId := uint(attendance[i].ID)
		detail, err := teachAttendanceService.GetAllTeachAttendanceDetail(uint(attendanceId))
		if err != nil {
			return
		}
		attendanceTime := attendance[i].CreatedAt.Format("2006-01-02")
		thisRecordAccount := len(detail)
		thisAttendanceAccount := 0
		thisAbsenceAccount := 0
		table := make(map[uint]uint) // studentId -> status
		for j := 0; j < len(detail); j++ {
			this := detail[j]
			studentId := this.StudentId
			table[studentId] = this.Status
			if this.Status != 0 {
				thisAttendanceAccount++
			} else {
				thisAbsenceAccount++
			}
		}
		excel.SetCellStr(xiangxiqingkuang, fmt.Sprintf("%c1", 'C'+i), attendanceTime)
		for j := 0; j < len(simple); j++ {
			value := -1
			if v, ok := table[simple[j].ID]; !ok {
				value = -1
			} else {
				value = int(v)
			}
			excel.SetCellInt(xiangxiqingkuang, fmt.Sprintf("%c%d", 'C'+i, 2+j), value)
		}

		excel.SetCellStr(zongtiqingkuang, fmt.Sprintf("A%d", i+2), attendanceTime)
		excel.SetCellInt(zongtiqingkuang, fmt.Sprintf("B%d", i+2), thisAttendanceAccount)
		excel.SetCellInt(zongtiqingkuang, fmt.Sprintf("C%d", i+2), thisAbsenceAccount)
		excel.SetCellInt(zongtiqingkuang, fmt.Sprintf("D%d", i+2), thisRecordAccount)
	}

	buffer, err := excel.WriteToBuffer()
	if err != nil {
		global.GVA_LOG.Error("将表格写入buffer错误" + err.Error())
		return
	}
	response.FileByReader(c, "考勤情况.xlsx", buffer)
	return
}

func (a *TeachAttendanceApi) attendanceExcelBase(overallSituationSheet string, specificSituationSheet string) *excelize.File {
	excel := excelize.NewFile()

	excel.NewSheet(overallSituationSheet)
	excel.SetCellStr(overallSituationSheet, "A1", "考勤时间")
	excel.SetCellStr(overallSituationSheet, "B1", "考勤人数")
	excel.SetCellStr(overallSituationSheet, "C1", "缺勤人数")
	excel.SetCellStr(overallSituationSheet, "D1", "记录数量")
	xiangxiqingkuang := "详细情况"
	excel.NewSheet(xiangxiqingkuang)
	excel.SetCellStr(xiangxiqingkuang, "A1", "学生学号")
	excel.SetCellStr(xiangxiqingkuang, "B1", "学生姓名")
	excel.DeleteSheet("Sheet1")
	return excel
}
