package exam

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/examType"
	"github.com/prl26/exam-system/server/model/examManage/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	request3 "github.com/prl26/exam-system/server/model/teachplan/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"github.com/prl26/exam-system/server/utils1"
	"go.uber.org/zap"
	"time"
)

type TargetExamApi struct {
}

var targetService = service.ServiceGroupApp.QuestionBankServiceGroup.TargetService
var targetOjService = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.TargetService
var targetExamService = service.ServiceGroupApp.ExammanageServiceGroup.TargetExamPaperService

// GetExamPaper 学生进入考试时获取试卷内容
func (targetExamApi *TargetExamApi) GetTargetExamPaper(c *gin.Context) {
	var planId request3.ExamPlan
	_ = c.ShouldBindQuery(&planId)
	studentId := utils.GetStudentId(c)
	ip := c.ClientIP()
	if bool, _ := examPlanService.CheckIsExamSt(planId.PlanId, studentId); bool == false {
		response.FailWithMessage("你不在参与此考试的范围中", c)
	} else {
		if isFinishPreExam, err, preExamIds := examPlanService.IsFinishPreExam(planId.PlanId, studentId); err != nil {
			response.FailWithMessage("前置计划查询出错", c)
		} else if isFinishPreExam == false {
			response.FailWithDetailedAndError(704, preExamIds, "请先完成前置计划", c)
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
			} else {
				if isReady, _ := examService.CheckIsReady(planId.PlanId); isReady == true {
					if examPaper, status, examScore, err := targetExamService.GetTargetExamPapersByRedis(examComing, ip); err != nil {
						global.GVA_LOG.Error("查询考试试卷失败", zap.Error(err))
						response.FailWithMessage("查询考试试卷失败", c)
					} else if status.IsCommit && PlanDetail.Type == examType.FinalExam {
						response.FailWithMessageAndError(703, "你已经提交过且通过该考试", c)
					} else if status.IsCommit && PlanDetail.Type == examType.ProceduralExam && *examScore.Score >= *PlanDetail.PassScore {
						response.FailWithMessageAndError(703, "你已经提交过了且通过该考试", c)
					} else {
						time3 := int64(time.Now().Sub(status.EnterTime).Seconds())
						time1 := *PlanDetail.Time*60 - time3
						time2 := int64(PlanDetail.EndTime.Sub(time.Now()).Seconds())
						var timeLimit int64
						if time1 <= time2 {
							timeLimit = time1
						} else {
							timeLimit = time2
						}
						response.OkWithData(gin.H{
							"examPaper": examPaper,
							"enterTime": status,
							"timeLimit": timeLimit,
						}, c)
					}
				} else {
					if examPaper, status, examScore, err := targetExamService.GetTargetExamPapers(examComing, ip); err != nil {
						global.GVA_LOG.Error("查询考试试卷失败", zap.Error(err))
						response.FailWithMessage("查询考试试卷失败", c)
					} else if status.IsCommit && PlanDetail.Type == examType.FinalExam {
						response.FailWithMessageAndError(703, "你已经提交过且通过该考试", c)
					} else if status.IsCommit && PlanDetail.Type == examType.ProceduralExam && *examScore.Score >= *PlanDetail.PassScore {
						response.FailWithMessageAndError(703, "你已经提交过了且通过该考试", c)
					} else {
						time3 := int64(time.Now().Sub(status.EnterTime).Seconds())
						time1 := *PlanDetail.Time*60 - time3
						time2 := int64(PlanDetail.EndTime.Sub(time.Now()).Seconds())
						var timeLimit int64
						if time1 <= time2 {
							timeLimit = time1
						} else {
							timeLimit = time2
						}
						response.OkWithData(gin.H{
							"examPaper": examPaper,
							"enterTime": status,
							"timeLimit": timeLimit,
						}, c)
					}
				}
			}
		}
	}
}
func (targetExamApi *TargetExamApi) ExamGenerateInstance(c *gin.Context) {
	var Instance request.TargetInstance
	_ = c.ShouldBindJSON(&Instance)
	byteCodeModel := targetService.GetByteCode(Instance.Id)
	if byteCodeModel == nil {
		questionBankResp.NotFind(c)
		return
	}
	salt, address, deployAddress, err := targetOjService.GenerateInstance(byteCodeModel.ByteCode)
	if err != nil {
		questionBankResp.ErrorHandle(c, fmt.Errorf("该题生成实例错误，请联系管理员检测"))
		return
	}
	studentId := utils.GetStudentId(c)
	targetService.ExamRecord(studentId, Instance.Id, address, Instance.PlanId)
	questionBankResp.OkWithDetailed(questionBankResp.TargetGenerateInstance{
		Address: address,
		//ByteCode: byteCodeModel.ByteCode,
		DeployAddress: deployAddress,
		Salt:          salt,
	}, "生成成功", c)
}

// CommitExamPaper 提交试卷内容
func (targetExamApi *TargetExamApi) CommitTargetExamPaper(c *gin.Context) {
	var ExamCommit request.CommitTargetExamPaper
	_ = c.ShouldBindJSON(&ExamCommit)
	ExamCommit.StudentId = utils.GetStudentId(c)
	PlanDetail, _ := examPlanService.GetExamPlan(ExamCommit.PlanId)
	status, _ := statuServie.GetStatus(ExamCommit.StudentId, ExamCommit.PlanId)
	minutes := *PlanDetail.Time
	unix1 := status.EnterTime.Add(time.Duration(minutes) * time.Minute)
	examScore, _ := statuServie.GetScore(ExamCommit.StudentId, ExamCommit.PlanId)
	if PlanDetail.IsLimitTime == true && time.Now().Unix() > unix1.Add(2*time.Minute).Unix() {
		response.FailWithMessageAndError(704, "超出考试时间", c)
	} else if time.Now().Unix() > PlanDetail.EndTime.Add(2*time.Minute).Unix() {
		response.FailWithMessageAndError(704, "提交失败,考试已经结束了", c)
	} else if status.IsCommit && PlanDetail.Type == examType.FinalExam {
		response.FailWithMessageAndError(703, "你已经提交过且通过该考试", c)
	} else if status.IsCommit && PlanDetail.Type == examType.ProceduralExam && *examScore.Score >= *PlanDetail.PassScore {
		response.FailWithMessageAndError(703, "你已经提交过了且通过该考试", c)
	} else {
		if err := targetExamService.CommitTargetExamPapers(ExamCommit); err != nil {
			global.GVA_LOG.Error("试卷提交失败", zap.Error(err))
			response.FailWithMessage("试卷提交试卷失败", c)
		} else {
			go func() {
				wg.Add(1)
				utils1.ExecTarget(ExamCommit)
				defer wg.Done()
				wg.Wait()
			}()
			response.OkWithData(gin.H{"examPaper": ExamCommit}, c)
		}
	}
}
func (targetExamApi *TargetExamApi) GetTargetExamScore(c *gin.Context) {
	var ScoreSearch request.ExamStudentScore
	_ = c.ShouldBindQuery(&ScoreSearch)
	StudentId := utils.GetStudentId(c)
	if scoreList, total, err := targetExamService.GetTargetExamScore(ScoreSearch, StudentId); err != nil {
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

// 获取考试做题情况
func (targetExamApi *TargetExamApi) GetTargetExamingScore(c *gin.Context) {
	var ExamCommit request.CommitTargetExamPaper
	_ = c.ShouldBindJSON(&ExamCommit)
	ExamCommit.StudentId = utils.GetStudentId(c)
	var scoreList []examManage.TargetExamingScore
	for _, v := range ExamCommit.TargetComponent {
		address, _ := targetService.QueryExamRecord(ExamCommit.StudentId, v.QuestionId, ExamCommit.PlanId)
		score, _ := targetOjService.QueryScore(address)
		list := examManage.TargetExamingScore{
			MergeId: v.MergeId,
			Answer:  address,
			Score:   score,
		}
		scoreList = append(scoreList, list)
	}
	response.OkWithData(gin.H{
		"scoreList": scoreList,
	}, c)
}
