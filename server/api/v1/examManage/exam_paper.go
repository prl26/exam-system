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
	"strconv"
)

type ExamPaperApi struct {
}

var examPaperService = service.ServiceGroupApp.ExammanageServiceGroup.ExamPaperService
var PaperTemplateItemService = service.ServiceGroupApp.ExammanageServiceGroup.PaperTemplateItemService

// CreateExamPaperByRand 创建ExamPaper
// @Tags ExamPaper
// @Summary 随机创建ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.ExamPaper true "创建ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaper/createExamPaper [post]
func (examPaperApi *ExamPaperApi) CreateExamPaperByRand(c *gin.Context) {
	var examPaper examManage.ExamPaper
	_ = c.ShouldBindJSON(&examPaper)
	numOfPapers := c.Query("numOfPapers")
	n, _ := strconv.Atoi(numOfPapers)
	templateItems, err := examPaperService.GetTemplate(examPaper)
	if err != nil {
		response.FailWithMessage("查询试卷模板失败", c)
	} else {
		if err := examPaperService.CreateExamPaper(examPaper); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("试卷创建失败", c)
		} else {
			for i := 0; i < n; i++ {
				if err := PaperTemplateItemService.SetPaperQuestion(templateItems); err != nil {
					global.GVA_LOG.Error("创建失败!", zap.Error(err))
					response.FailWithMessage("创建试卷失败", c)
				} else {
					response.OkWithMessage("创建成功", c)
				}
			}
		}
	}
}

// DeleteExamPaper 删除ExamPaper
// @Tags ExamPaper
// @Summary 删除ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.ExamPaper true "删除ExamPaper"
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
// @Param data body examManage.ExamPaper true "更新ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examPaper/updateExamPaper [put]
func (examPaperApi *ExamPaperApi) UpdateExamPaper(c *gin.Context) {
	var examPaper examManage.ExamPaper
	_ = c.ShouldBindJSON(&examPaper)
	if err := examPaperService.UpdateExamPaper(examPaper); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExamPaper 用id查询ExamPaper
// @Tags ExamPaper
// @Summary 用id查询ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManage.ExamPaper true "用id查询ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examPaper/findExamPaper [get]
func (examPaperApi *ExamPaperApi) FindExamPaper(c *gin.Context) {
	var examPaper examManage.ExamPaper
	_ = c.ShouldBindQuery(&examPaper)
	if reexamPaper, err := examPaperService.GetExamPaper(examPaper.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexamPaper": reexamPaper}, c)
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
	if list, total, err := examPaperService.GetExamPaperInfoList(pageInfo); err != nil {
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
