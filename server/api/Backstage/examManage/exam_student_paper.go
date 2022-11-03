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
)

type ExamStudentPaperApi struct {
}

var examstudentPaperService = service.ServiceGroupApp.ExammanageServiceGroup.ExamStudentPaperService

// CreateExamStudentPaper 创建ExamStudentPaper
// @Tags ExamStudentPaper
// @Summary 创建ExamStudentPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.ExamStudentPaper true "创建ExamStudentPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examstudentPaper/createExamStudentPaper [post]
func (examstudentPaperApi *ExamStudentPaperApi) CreateExamStudentPaper(c *gin.Context) {
	var examstudentPaper examManageReq.ExamComing
	_ = c.ShouldBindJSON(&examstudentPaper)
	if st, err := examstudentPaperService.CreateExamStudentPaper(examstudentPaper); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		//response.OkWithMessage("创建成功", c)
		response.OkWithData(gin.H{"examPaper": st}, c)
	}
}

// DeleteExamStudentPaper 删除ExamStudentPaper
// @Tags ExamStudentPaper
// @Summary 删除ExamStudentPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.ExamStudentPaper true "删除ExamStudentPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examstudentPaper/deleteExamStudentPaper [delete]
func (examstudentPaperApi *ExamStudentPaperApi) DeleteExamStudentPaper(c *gin.Context) {
	var examstudentPaper examManage.ExamStudentPaper
	_ = c.ShouldBindJSON(&examstudentPaper)
	if err := examstudentPaperService.DeleteExamStudentPaper(examstudentPaper); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)

	}
}

// DeleteExamStudentPaperByIds 批量删除ExamStudentPaper
// @Tags ExamStudentPaper
// @Summary 批量删除ExamStudentPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExamStudentPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /examstudentPaper/deleteExamStudentPaperByIds [delete]
func (examstudentPaperApi *ExamStudentPaperApi) DeleteExamStudentPaperByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := examstudentPaperService.DeleteExamStudentPaperByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExamStudentPaper 更新ExamStudentPaper
// @Tags ExamStudentPaper
// @Summary 更新ExamStudentPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body frontExamManage.ExamStudentPaper true "更新ExamStudentPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examstudentPaper/updateExamStudentPaper [put]
func (examstudentPaperApi *ExamStudentPaperApi) UpdateExamStudentPaper(c *gin.Context) {
	var examstudentPaper examManage.ExamStudentPaper
	_ = c.ShouldBindJSON(&examstudentPaper)
	if err := examstudentPaperService.UpdateExamStudentPaper(examstudentPaper); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExamStudentPaper 用id查询ExamStudentPaper
// @Tags ExamStudentPaper
// @Summary 用id查询ExamStudentPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query frontExamManage.ExamStudentPaper true "用id查询ExamStudentPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examstudentPaper/findExamStudentPaper [get]
func (examstudentPaperApi *ExamStudentPaperApi) FindExamStudentPaper(c *gin.Context) {
	var examstudentPaper examManage.ExamStudentPaper
	_ = c.ShouldBindQuery(&examstudentPaper)
	if reexamstudentPaper, err := examstudentPaperService.GetExamStudentPaper(examstudentPaper.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexamstudentPaper": reexamstudentPaper}, c)
	}
}

// GetExamStudentPaperList 分页获取ExamStudentPaper列表
// @Tags ExamStudentPaper
// @Summary 分页获取ExamStudentPaper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManageReq.ExamStudentPaperSearch true "分页获取ExamStudentPaper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examstudentPaper/getExamStudentPaperList [get]
func (examstudentPaperApi *ExamStudentPaperApi) GetExamStudentPaperList(c *gin.Context) {
	var pageInfo examManageReq.ExamStudentPaperSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := examstudentPaperService.GetExamStudentPaperInfoList(pageInfo); err != nil {
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
