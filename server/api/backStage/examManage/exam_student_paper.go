package examManage

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"github.com/prl26/exam-system/server/utils1"
	"go.uber.org/zap"
	"time"
)

type ExamStudentPaperApi struct {
}

var examstudentPaperService = service.ServiceGroupApp.ExammanageServiceGroup.ExamStudentPaperService
var examService = service.ServiceGroupApp.ExammanageServiceGroup.ExamService

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
	var examstudentPaper examManageReq.ExamComing
	_ = c.ShouldBindQuery(&examstudentPaper)
	if reexamstudentPaper, _, err := examService.GetExamPapersAndScores(examstudentPaper, ""); err != nil {
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

// 状态监测
func (examstudentPaperApi *ExamStudentPaperApi) StatusMonitor(c *gin.Context) {
	var pageInfo examManageReq.StatusMonitor
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := examstudentPaperService.StudentPaperStatus(pageInfo); err != nil {
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

// 恢复学生考试资格
func (examstudentPaperApi *ExamStudentPaperApi) RecoverPower(c *gin.Context) {
	var plan teachplan.CoverRq
	_ = c.ShouldBindJSON(&plan)
	planDetail, _ := examService.GetPlanDetail(plan.PlanId)
	if planDetail.EndTime.Unix() < time.Now().Unix() {
		response.FailWithMessage("已经超过考试截至时间了", c)
	} else {
		if err := examstudentPaperService.RecoverStudentPower(plan.StudentId, plan.PlanId); err != nil {
			response.FailWithMessage("恢复失败", c)
		} else {
			response.OkWithMessage("恢复成功", c)
		}
	}
}

// 上报学生分数
func (examstudentPaperApi *ExamStudentPaperApi) ReportScore(c *gin.Context) {
	var st teachplan.CoverRqs
	_ = c.ShouldBindJSON(&st)
	var isCheck bool
	for _, v := range st.StudentIds {
		lst := teachplan.CoverRq{
			StudentId: v,
			PlanId:    st.PlanId,
		}
		if bool, _ := examstudentPaperService.CheckIsCommit(lst); bool == false {
			isCheck = false
			break
		}
	}
	if isCheck == false {
		response.FailWithMessage("选中的学生中有未提交的学生", c)
	} else {
		if err := examstudentPaperService.ReportScore(st); err != nil {
			global.GVA_LOG.Error("上报成绩失败!", zap.Error(err))
			response.FailWithMessage("上报失败", c)
		} else {
			response.OkWithMessage("上报成功", c)
		}
	}
}

// 试卷批阅
func (examstudentPaperApi *ExamStudentPaperApi) PaperReview(c *gin.Context) {
	var pageInfo examManageReq.PaperReview
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := examstudentPaperService.ReviewScore(pageInfo); err != nil {
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

// 成绩及答案的修正
func (examstudentPaperApi *ExamStudentPaperApi) PaperCheating(c *gin.Context) {
	var cheating examManageReq.PaperCheating
	_ = c.ShouldBindJSON(&cheating)
	if err := examstudentPaperService.PaperCheating(cheating); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithData(gin.H{
			"修改数据":    cheating.AnswerCheating,
			"message": "修改成功",
		}, c)
	}
}

// 单独上报学生分数
func (examstudentPaperApi *ExamStudentPaperApi) ReportStudentScore(c *gin.Context) {
	var st teachplan.CoverRq
	_ = c.ShouldBindJSON(&st)
	if bool, err := examstudentPaperService.CheckIsCommit(st); err != nil {
		response.FailWithMessage("查询试卷状态失败", c)
	} else if bool == false {
		response.FailWithMessage("该生还未提交试卷", c)
	} else {
		if err := examstudentPaperService.ReportStudentScore(st.PlanId, st.StudentId); err != nil {
			global.GVA_LOG.Error("上报成绩失败!", zap.Error(err))
			response.FailWithMessage("上报失败", c)
		} else {
			response.OkWithMessage("上报成功", c)
		}
	}

}

// 单个学生试卷重新批阅
func (examstudentPaperApi *ExamStudentPaperApi) ExecAgain(c *gin.Context) {
	var sp teachplan.CoverRq
	_ = c.ShouldBindJSON(&sp)
	if err := utils.ReExecPapers(sp); err != nil {
		global.GVA_LOG.Error("自动批阅出错啦!", zap.Error(err))
		response.FailWithMessage("自动批阅出错啦", c)
	} else {
		if err := utils1.ReExecTargetPapers(sp); err != nil {
			global.GVA_LOG.Error("自动批阅出错啦!", zap.Error(err))
			response.FailWithMessage("自动批阅出错啦"+err.Error(), c)
			return
		}
		response.OkWithMessage("批阅成功", c)
	}
}

// 考试计划下学生试卷重新批阅
func (examstudentPaperApi *ExamStudentPaperApi) AllExecAgain(c *gin.Context) {
	var sp teachplan.CoverRqs
	_ = c.ShouldBindJSON(&sp)
	for _, v := range sp.StudentIds {
		lsp := teachplan.CoverRq{
			StudentId: v,
			PlanId:    sp.PlanId,
		}
		if err := utils.ReExecPapers(lsp); err != nil {
			global.GVA_LOG.Error("自动批阅出错啦!", zap.Error(err))
			response.FailWithMessage("自动批阅出错啦", c)
		} else {
			if err := utils1.ReExecTargetPapers(lsp); err != nil {
				global.GVA_LOG.Error("自动批阅出错啦!", zap.Error(err))
				response.FailWithMessage("自动批阅出错啦", c)
			}
		}
	}
	response.OkWithMessage("批阅成功", c)
}

// 获取考试提交日志
func (examstudentPaperApi *ExamStudentPaperApi) GetCommitRecord(c *gin.Context) {
	var recordRq examManageReq.RecordRq
	_ = c.ShouldBindQuery(&recordRq)
	if recordList, err := examstudentPaperService.GetCommitRecord(recordRq); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{
			"recordList": recordList,
			"message":    "获取成功",
		}, c)
	}
}

// 通过指定的日志id恢复学生试卷
func (examstudentPaperApi *ExamStudentPaperApi) RecoverByRecord(c *gin.Context) {
	var recordRq examManageReq.RecordRq1
	_ = c.ShouldBindQuery(&recordRq)
	if err := examstudentPaperService.RecoverByRecord(recordRq.PlanId, recordRq.Student, recordRq.RecordId); err != nil {
		global.GVA_LOG.Error("恢复学生试卷失败!", zap.Error(err))
		response.FailWithMessage("恢复学生试卷失败", c)
	} else {
		response.OkWithMessage("恢复学生试卷成功", c)
	}
}

// 删除学生考卷
func (examstudentPaperApi *ExamStudentPaperApi) DeleteStudentAnswer(c *gin.Context) {
	var recordRq examManageReq.RecordRq
	_ = c.ShouldBindJSON(&recordRq)
	if err := examstudentPaperService.DeleteAnswer(recordRq.PlanId, recordRq.Student); err != nil {
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// 强制提交学生试卷
func (examstudentPaperApi *ExamStudentPaperApi) ForceCommitStudent(c *gin.Context) {
	var recordRq examManageReq.RecordRq
	_ = c.ShouldBindJSON(&recordRq)
	if err := examstudentPaperService.ForceCommitStudent(recordRq.PlanId, recordRq.Student); err != nil {
		response.FailWithMessage("强制提交失败", c)
	} else {
		response.OkWithMessage("强制提交成功", c)
	}
}
