package exam

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	request2 "github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage/request"
	request3 "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"time"
)

type TargetExamApi struct {
}

var targetService = service.ServiceGroupApp.QuestionBankServiceGroup.TargetService
var targetOjService = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.TargetService
var targetExamService = service.ServiceGroupApp.ExammanageServiceGroup.TargetExamPaperService

func (examApi *ExamApi) FindTargetExamPlans(c *gin.Context) {
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
func (targetExamApi *TargetExamApi) GetTargetExamPaper(c *gin.Context) {
	var planId request3.ExamPlan
	_ = c.ShouldBindQuery(&planId)
	studentId := utils.GetStudentId(c)
	if isFinishPreExam, err, preExamIds := examPlanService.IsFinishPreExam(planId.PlanId, studentId); err != nil {
		response.FailWithMessage("前置计划查询出错", c)
	} else if isFinishPreExam == false {
		response.FailWithDetailed(preExamIds, "请先完成前置计划", c)
	} else {
		var examComing = request.ExamComing{
			StudentId: studentId,
			PlanId:    planId.PlanId,
		}
		PlanDetail, _ := examPlanService.GetExamPlan(planId.PlanId)
		if PlanDetail.StartTime.Unix() > time.Now().Unix() {
			response.FailWithMessageAndError(701, "还没开考呢,莫急", c)
		} else if PlanDetail.EndTime.Unix() < time.Now().Unix() {
			response.FailWithMessageAndError(702, "你来晚了,考试已经结束了", c)
		} else if examPaper, status, err := targetExamService.GetTargetExamPapers(examComing); err != nil {
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
}

// CommitExamPaper 提交试卷内容
func (targetExamApi *TargetExamApi) CommitTargetExamPaper(c *gin.Context) {
	var ExamCommit request.CommitTargetExamPaper
	_ = c.ShouldBindJSON(&ExamCommit)
	ExamCommit.StudentId = utils.GetStudentId(c)
	PlanDetail, _ := examPlanService.GetExamPlan(ExamCommit.PlanId)
	status, _ := statuServie.GetStatus(ExamCommit.StudentId, ExamCommit.PlanId)
	if time.Now().Unix() > PlanDetail.EndTime.Unix() {
		response.FailWithMessageAndError(704, "提交失败,考试已经结束了", c)
	} else if status.IsCommit {
		response.FailWithMessage("你已经提交过了", c)
	} else {
		if err := targetExamService.CommitTargetExamPapers(ExamCommit); err != nil {
			global.GVA_LOG.Error("试卷提交失败", zap.Error(err))
			response.FailWithMessage("试卷提交试卷失败", c)
		} else {
			go func() {
				fmt.Println("start,开始处理试卷")
				time.AfterFunc(time.Second*5, func() {
					wg.Add(1)
					//utils.ExecPapers(ExamCommit)
					defer wg.Done()
				})
				wg.Wait()
			}()
			response.OkWithData(gin.H{"examPaper": ExamCommit}, c)
		}
	}
}
