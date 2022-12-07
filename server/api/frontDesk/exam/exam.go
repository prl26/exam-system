package exam

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	request2 "github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	response2 "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	request3 "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"sync"
	"time"
)

type ExamApi struct {
}

var wg sync.WaitGroup
var examService = service.ServiceGroupApp.ExammanageServiceGroup.ExamService
var statuServie = service.ServiceGroupApp.ExammanageServiceGroup.ExamStatusService
var examPlanService = service.ServiceGroupApp.TeachplanServiceGroup.ExamPlanService
var programService = &service.ServiceGroupApp.QuestionBankServiceGroup.OjService.ProgramService

// FindExamPlans 查询该学生 某个教学班 下所有的考试计划
func (examApi *ExamApi) FindExamPlans(c *gin.Context) {
	var teachClassId request2.GetByTeachClassId
	_ = c.ShouldBindQuery(&teachClassId)
	if examPlans, err := examService.FindExamPlans(teachClassId.TeachClassId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"examPlans": examPlans}, c)
	}
}

// GetExamPaper 学生进入考试时获取试卷内容
func (examApi *ExamApi) GetExamPaper(c *gin.Context) {
	var planId request3.ExamPlan
	_ = c.ShouldBindQuery(&planId)
	studentId := utils.GetStudentId(c)
	var examComing = request.ExamComing{
		StudentId: studentId,
		PlanId:    planId.PlanId,
	}
	PlanDetail, _ := examPlanService.GetExamPlan(planId.PlanId)
	if PlanDetail.StartTime.Unix() > time.Now().Unix() {
		response.FailWithMessageAndError(701, "还没开考呢,莫急", c)
	} else if PlanDetail.EndTime.Unix() < time.Now().Unix() {
		response.FailWithMessageAndError(702, "你来晚了,考试已经结束了", c)
	} else if examPaper, status, err := examService.GetExamPapers(examComing); err != nil {
		global.GVA_LOG.Error("查询考试试卷失败", zap.Error(err))
		response.FailWithMessage("查询考试试卷失败", c)
	} else if status.IsCommit {
		response.FailWithMessageAndError(703, "你已经提交过了", c)
	} else {
		response.OkWithData(gin.H{
			"examPaper": examPaper,
			"enterTime": status.EnterTime,
		}, c)
	}
}

// CommitExamPaper 提交试卷内容
func (examApi *ExamApi) CommitExamPaper(c *gin.Context) {
	var ExamCommit examManage.CommitExamPaper
	_ = c.ShouldBindJSON(&ExamCommit)
	ExamCommit.StudentId = utils.GetStudentId(c)
	PlanDetail, _ := examPlanService.GetExamPlan(ExamCommit.PlanId)
	status, _ := statuServie.GetStatus(ExamCommit.StudentId, ExamCommit.PlanId)
	if time.Now().Unix() > PlanDetail.EndTime.Unix() {
		response.FailWithMessageAndError(704, "提交失败,考试已经结束了", c)
	} else if status.IsCommit {
		response.FailWithMessage("你已经提交过了", c)
	} else {
		if err := examService.CommitExamPapers(ExamCommit); err != nil {
			global.GVA_LOG.Error("试卷提交失败", zap.Error(err))
			response.FailWithMessage("试卷提交试卷失败", c)
		} else {
			response.OkWithData(gin.H{"examPaper": ExamCommit}, c)
			go func() {
				fmt.Println("start")
				time.AfterFunc(time.Second*5, func() {
					wg.Add(1)
					utils.ExecPapers(ExamCommit)
					defer wg.Done()
				})
				wg.Wait()
			}()
		}
	}
}

// 考试的时候提交编程题
func (examApi *ExamApi) CommitProgram(c *gin.Context) {
	var program examManage.CommitProgram
	_ = c.ShouldBindJSON(&program)
	program.StudentId = utils.GetStudentId(c)
	resp := make(chan response2.SubmitResponse)
	var err error
	go func() {
		checkProgram, score, e := programService.CheckProgram(program.QuestionId, program.Code, program.LanguageId)
		err = e
		resp <- response2.SubmitResponse{Submit: checkProgram, Score: score}
	}()
	program.StudentId = utils.GetStudentId(c)
	PlanDetail, _ := examPlanService.GetExamPlan(program.PlanId)
	err = examService.CommitProgram(program)
	if err != nil {
		response.FailWithMessage("更新失败", c)
	} else {
		if time.Now().Unix() > PlanDetail.EndTime.Unix() {
			response.FailWithMessageAndError(704, "提交失败,考试已经结束了", c)
		} else {
			result := <-resp
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				response2.ErrorHandle(c, err)
				return
			}
			err := utils.ExecProgram(program, result.Score)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				response.Fail(c)
				return
			}
			response.OkWithData(result.Score, c)
		}
	}
}

//获取考试分数
func (ExamApi *ExamApi) GetExamScore(c *gin.Context) {
	var ScoreSearch request.ExamStudentScore
	_ = c.ShouldBindQuery(&ScoreSearch)
	StudentId := utils.GetStudentId(c)
	if scoreList, total, err := examService.GetExamScore(ScoreSearch, StudentId); err != nil {
		global.GVA_LOG.Error("查询成绩失败", zap.Error(err))
		response.FailWithMessage("查询成绩失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     scoreList,
			Total:    total,
			Page:     ScoreSearch.Page,
			PageSize: ScoreSearch.PageSize,
		}, "获取成功", c)
	}
}
