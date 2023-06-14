package examManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/teachplan"
	request3 "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
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
		"Name":   {utils.NotEmpty()},
		"PlanId": {utils.NotEmpty()},
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
	} else {
		templateId, err := examPaperService.FindTemplateId(examPaper)
		if templateId == 0 || err != nil {
			response.FailWithMessage("请先绑定模板", c)
		} else {
			for i := 0; i < n; i++ {
				if err := examPaperService.CreateExamPaper(examPaper); err != nil {
					global.GVA_LOG.Error("创建失败!", zap.Error(err))
					response.FailWithMessage("试卷创建失败", c)
				}
			}
			response.OkWithMessage("创建成功", c)
		}
	}
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
	if count := examPaperService.CheckPaperIsUsed(examPaper.ID); count != 0 {
		response.FailWithMessage("该试卷已被使用,删除失败", c)
	} else {
		if err := examPaperService.DeleteExamPaper(examPaper); err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		} else {
			response.OkWithMessage("删除成功", c)
		}
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
	for _, v := range IDS.Ids {
		if count := examPaperService.CheckPaperIsUsed(v); count != 0 {
			response.FailWithMessage(fmt.Sprintf("id为%d 的试卷已经被使用", v), c)
		}
	}
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

//迟到学生试卷分发
func (examPaperApi *ExamPaperApi) LateStudentDistribute(c *gin.Context) {
	var plan teachplan.ExamPlan
	_ = c.ShouldBindJSON(&plan)
	if diffArray, err := examPaperService.FindLateJoinStd(plan.ID); err != nil {
		response.FailWithMessage("查询未分发试卷的学生出错", c)
	} else {
		if number, err := examPaperService.GetPaperNum(plan.ID); err != nil {
			response.FailWithMessage("试卷分发失败", c)
		} else if len(number) == 0 {
			response.FailWithMessageAndError(602, "需要先为计划生成试卷", c)
		} else {
			err = examPaperService.LateStdsDistribution(plan.ID, diffArray, number)
			if err != nil {
				response.FailWithMessage("试卷补发失败", c)
			} else {
				response.OkWithMessage("试卷补发成功", c)
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
	respath := "/static/excel/" + excelInfo.FileName
	c.Writer.Header().Add("Content-Disposition", "attachment; filepath="+filePath)
	//c.File(filePath)
	response.OkWithData(gin.H{
		"filepath": respath,
	}, c)
	quesNum, _ := examService.GetPaperQuesNum(excelInfo.PlanId)
	infoList, _ := examService.GetExamScoreToExcel(excelInfo.PlanId)
	studentList, _ := examService.GetStudentList(excelInfo.PlanId)
	err := examService.ExportPaperScore(excelInfo.PlanId, studentList, infoList, filePath, quesNum[0].Num)
	if err != nil {
		global.GVA_LOG.Error("转换Excel失败!", zap.Error(err))
		return
	}
	//else {
	//	c.Writer.Header().Add("Content-Disposition", "attachment; filepath="+filePath)
	//	//c.File(filePath)
	//	response.OkWithData(gin.H{
	//		"filepath": respath,
	//	}, c)
	//}
}
func (examPaperApi *ExamPaperApi) ExportPaperToHtml(c *gin.Context) {
	var excelInfo request3.Excel
	_ = c.ShouldBindJSON(&excelInfo)
	if strings.Index(excelInfo.FileName, "..") > -1 {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	_, err := examService.ExportPaperToHtml(excelInfo.PlanId, excelInfo.FileName)
	if err != nil {
		global.GVA_LOG.Error("生成zip失败!", zap.Error(err))
		return
	} else {
		path := "/static/html/zip/" + fmt.Sprintf("%s.zip", excelInfo.FileName)
		response.OkWithData(gin.H{
			"filepath": path,
		}, c)
	}
}
func (examPaperApi *ExamPaperApi) ExportPaperToHtmlToCheck(c *gin.Context) {
	var excelInfo request3.Excel
	_ = c.ShouldBindJSON(&excelInfo)
	if strings.Index(excelInfo.FileName, "..") > -1 {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	var paper []examManage.ExamPaper
	global.GVA_DB.Model(examManage.ExamPaper{}).Where("plan_id = ?", excelInfo.PlanId).Find(&paper)
	if len(paper) == 1 {
		reexamPaper, examPaperTitle, err := examPaperService.GetExamPaper1(paper[0].ID)
		if _, err = examService.ExportPaperToHtmlToCheck(excelInfo.PlanId, excelInfo.FileName, reexamPaper, examPaperTitle); err != nil {
			global.GVA_LOG.Error("生成html失败!", zap.Error(err))
			return
		} else {
			path := "/static/html/" + excelInfo.FileName
			response.OkWithData(gin.H{
				"filepath": path,
			}, c)
		}
	} else {
		response.FailWithMessage("暂时不支持有多份试卷", c)
	}

}
func (examPaperApi *ExamPaperApi) ExportPaperToHtmlToCheckWithOutAnswer(c *gin.Context) {
	var excelInfo request3.Excel
	_ = c.ShouldBindJSON(&excelInfo)
	if strings.Index(excelInfo.FileName, "..") > -1 {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	var paper []examManage.ExamPaper
	global.GVA_DB.Model(examManage.ExamPaper{}).Where("plan_id = ?", excelInfo.PlanId).Find(&paper)
	if len(paper) == 1 {
		reexamPaper, examPaperTitle, err := examPaperService.GetExamPaper1(paper[0].ID)
		if _, err = examService.ExportPaperToHtmlToCheck(excelInfo.PlanId, excelInfo.FileName, reexamPaper, examPaperTitle); err != nil {
			global.GVA_LOG.Error("生成html失败!", zap.Error(err))
			return
		} else {
			path := "/static/html/" + excelInfo.FileName
			response.OkWithData(gin.H{
				"filepath": path,
			}, c)
		}
	} else {
		response.FailWithMessage("暂时不支持有多份试卷", c)
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
	respath := "/static/excel/" + excelInfo.FileName
	c.Writer.Header().Add("Content-Disposition", "attachment; filepath="+filePath)
	//c.File(filePath)
	infoList, _ := examService.GetPlanList(excelInfo.TeachPlanId)
	studentList, _ := examService.GetStudentListByTeachPlan(excelInfo.TeachPlanId)
	err := examService.ExportMultiPaperScore(studentList, infoList, filePath)
	if err != nil {
		global.GVA_LOG.Error("转换Excel失败!", zap.Error(err))
		return
	} else {
		c.Writer.Header().Add("Content-Disposition", "attachment; filepath="+filePath)
		//c.File(filePath)
		response.OkWithData(gin.H{
			"filepath": respath,
		}, c)
	}
}

//进入考试准备阶段
func (examPaperApi *ExamPaperApi) SetExamPre(c *gin.Context) {
	var plan teachplan.ExamPlan
	_ = c.ShouldBindJSON(&plan)
	plandetail, _ := examService.GetPlanDetail(plan.ID)
	unix1 := time.Now().Add(24 * time.Hour)
	if isDistributed, _ := examService.CheckIsDistributed(plan.ID); isDistributed == false {
		response.FailWithMessage("需要先分配试卷", c)
	} else if plandetail.StartTime.Unix() > unix1.Unix() {
		response.FailWithMessage("进入考试准备阶段需要在开考的前一天内", c)
	} else if plandetail.StartTime.Unix() < time.Now().Unix() {
		response.FailWithMessage("考试已经开考了", c)
	} else if plandetail.EndTime.Unix() < time.Now().Unix() {
		response.FailWithMessage("考试已经结束了", c)
	} else {
		err := examService.SetExamPre(plan.ID)
		if err != nil {
			response.FailWithMessage("进入考试准备阶段失败", c)
		} else {
			response.OkWithMessage("成功进入考试准备阶段", c)
		}
	}
}
