package examManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	request3 "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type ExamPaperApi struct {
}

var examPaperService = service.ServiceGroupApp.ExammanageServiceGroup.ExamPaperService
var PaperTemplateItemService = service.ServiceGroupApp.ExammanageServiceGroup.PaperTemplateItemService
var examStatusService = service.ServiceGroupApp.ExammanageServiceGroup.ExamStatusService
var multiTableService = service.ServiceGroupApp.BasicdataApiGroup.MultiTableService

// CreateExamPaperByRand 创建ExamPaper
// @Tags ExamPaper
// @Summary 随机创建ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.ExamPaper true "创建ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaper/createExamPaper [post]
func (examPaperApi *ExamPaperApi) CreateExamPaperByRand(c *gin.Context) {
	var examPaper examManage.ExamPaper
	_ = c.ShouldBindJSON(&examPaper)
	verify := utils.Rules{
		"Name":       {utils.NotEmpty()},
		"PlanId":     {utils.NotEmpty()},
		"TemplateId": {utils.NotEmpty()},
	}
	if err := utils.Verify(examPaper, verify); err != nil {
		response.CheckHandle(c, err)
		return
	}
	userId := utils.GetUserID(c)
	examPaper.UserId = &userId
	numOfPapers := c.Query("numOfPapers")
	n, _ := strconv.Atoi(numOfPapers)
	if n == 0 {
		response.FailWithMessage("试卷份数不能为0", c)
	}
	for i := 0; i < n; i++ {
		if err := examPaperService.CreateExamPaper(examPaper); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("试卷创建失败", c)
		}
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExamPaper 删除ExamPaper
// @Tags ExamPaper
// @Summary 删除ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.ExamPaper true "删除ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPaper/deleteExamPaper [delete]
func (examPaperApi *ExamPaperApi) DeleteExamPaper(c *gin.Context) {
	var examPaper examManage.ExamPaper
	_ = c.ShouldBindJSON(&examPaper)
	if err := examPaperService.DeleteExamPaper(examPaper); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExamPaperByIds 批量删除ExamPaper
// @Tags ExamPaper
// @Summary 批量删除ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /examPaper/deleteExamPaperByIds [delete]
func (examPaperApi *ExamPaperApi) DeleteExamPaperByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := examPaperService.DeleteExamPaperByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExamPaper 更新ExamPaper
// @Tags ExamPaper
// @Summary 更新ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.ExamPaper true "更新ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examPaper/updateExamPaper [put]
func (examPaperApi *ExamPaperApi) UpdateExamPaper(c *gin.Context) {
	var examPaper examManage.ExamPaper
	_ = c.ShouldBindJSON(&examPaper)
	if bool, err := examPaperService.UpdateExamPaper(examPaper); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		if bool == true {
			err = examPaperService.DeletePaperMerge(examPaper)
			if err != nil {
				response.FailWithMessage("删除失败", c)
			}
			err = examPaperService.CreateExamPaper(examPaper)
			if err != nil {
				response.FailWithMessage("更新试卷失败", c)
			}
		}
		response.OkWithMessage("更新成功", c)
	}
}

// FindExamPaper 用id查询ExamPaper
// @Tags ExamPaper
// @Summary 用id查询ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query frontExamManage.ExamPaper true "用id查询ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examPaper/findExamPaper [get]
func (examPaperApi *ExamPaperApi) FindExamPaper(c *gin.Context) {
	var examPaper examManage.ExamPaper
	_ = c.ShouldBindQuery(&examPaper)
	if reexamPaper, examPaperTitle, err := examPaperService.GetExamPaper(examPaper.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{
			"reexamPaper":    reexamPaper,
			"examPaperTitle": examPaperTitle,
		}, c)
	}
}

// GetExamPaperList 分页获取ExamPaper列表
// @Tags ExamPaper
// @Summary 分页获取ExamPaper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManageReq.ExamPaperSearch true "分页获取ExamPaper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaper/getExamPaperList [get]
func (examPaperApi *ExamPaperApi) GetExamPaperList(c *gin.Context) {
	var pageInfo examManageReq.ExamPaperSearch
	_ = c.ShouldBindQuery(&pageInfo)
	userId := utils.GetUserID(c)
	authorityId := utils.GetUserAuthorityID(c)
	fmt.Println(pageInfo)
	if list, total, err := examPaperService.GetExamPaperInfoList(pageInfo, userId, authorityId); err != nil {
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

// 判断学生最近是否有考试
func (examPaperApi *ExamPaperApi) SetStudentsToRedis(c *gin.Context) {
	students, err := examStatusService.GaSStudentsOfExam()
	if err != nil {
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithData(gin.H{"data": students}, c)
	}
}

//试卷分发
func (examPaperApi *ExamPaperApi) PaperDistribution(c *gin.Context) {
	var planId examManageReq.PaperDistribution
	_ = c.ShouldBindQuery(&planId)
	status, err := examPaperService.GetPlanStatus(planId.PlanId)
	if err != nil {
		response.FailWithMessageAndError(601, "试卷分发失败", c)
	} else if status {
		response.FailWithMessage("试卷已经分发了", c)
	} else {
		if number, err := examPaperService.GetPaperNum(planId.PlanId); err != nil {
			response.FailWithMessage("试卷分发失败", c)
		} else if len(number) == 0 {
			response.FailWithMessageAndError(602, "需要先为计划生成试卷", c)
		} else {
			err = examPaperService.PaperDistribution(planId.PlanId, number)
			if err != nil {
				response.FailWithMessage("试卷分发失败", c)
			} else {
				response.OkWithMessage("试卷分发成功", c)
			}
		}
	}
}

//导出成绩表
func (examPaperApi *ExamPaperApi) ExportPaper(c *gin.Context) {
	var excelInfo request3.Excel
	_ = c.ShouldBindJSON(&excelInfo)
	if strings.Index(excelInfo.FileName, "..") > -1 {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	//todaystr1 := time.Now().Format("2006-01-02-f15:04:05")
	filePath := global.GVA_CONFIG.Excel.Dir + excelInfo.FileName
	respath := "/static/" + excelInfo.FileName
	infoList, _ := examService.GetExamScoreToExcel(excelInfo.PlanId)
	err := examService.ExportPaperScore(infoList, filePath)
	if err != nil {
		global.GVA_LOG.Error("转换Excel失败!", zap.Error(err))
		response.FailWithMessage("转换Excel失败", c)
		return
	} else {
		c.Writer.Header().Add("Content-Disposition", "attachment; filepath="+filePath)
		//c.File(filePath)
		response.OkWithData(gin.H{
			"filepath": respath,
		}, c)
	}
}
func (examPaperApi *ExamPaperApi) ExportMultiPaper(c *gin.Context) {
	var excelInfo request3.MultiExcel
	_ = c.ShouldBindJSON(&excelInfo)
	if strings.Index(excelInfo.FileName, "..") > -1 {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	//todaystr1 := time.Now().Format("2006-01-02-f15:04:05")
	filePath := global.GVA_CONFIG.Excel.Dir + excelInfo.FileName
	respath := "/static/" + excelInfo.FileName
	infoList, _ := examService.GetPlanList(excelInfo.TeachPlanId)
	err := examService.ExportMultiPaperScore(infoList, filePath)
	if err != nil {
		global.GVA_LOG.Error("转换Excel失败!", zap.Error(err))
		response.FailWithMessage("转换Excel失败", c)
		return
	} else {
		c.Writer.Header().Add("Content-Disposition", "attachment; filepath="+filePath)
		//c.File(filePath)
		response.OkWithData(gin.H{
			"filepath": respath,
		}, c)
	}
}
