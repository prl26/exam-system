package exam

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	request2 "github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
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

// FindExamPlans 查询该学生 某个教学班 下所有的教学计划
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
	status, _ := statuServie.GetStatus(ExamCommit.StudentId, ExamCommit.PlanId)
	if status.IsCommit {
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

//获取考试分数
func (ExamApi *ExamApi) GetExamScore(c *gin.Context) {
	StudentId := utils.GetStudentId(c)
	if score, err := examService.GetExamScore(StudentId); err != nil {
		global.GVA_LOG.Error("查询成绩失败", zap.Error(err))
		response.FailWithMessage("查询成绩失败", c)
	} else {
		response.OkWithData(gin.H{"score": score}, c)
	}
}
