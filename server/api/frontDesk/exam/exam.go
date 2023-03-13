package exam

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	request2 "github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/examType"
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
var termService = service.ServiceGroupApp.BasicdataApiGroup.TermService

// FindExamPlans 查询该学生 某个教学班 下所有的考试计划
func (examApi *ExamApi) FindExamPlans(c *gin.Context) {
	var teachClassId request2.GetByTeachClassId
	_ = c.ShouldBindQuery(&teachClassId)
	sid := utils.GetStudentId(c)
	if examPlans, err := examService.FindExamPlans(teachClassId.TeachClassId, sid); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"examPlans": examPlans}, c)
	}
}
func (examApi *ExamApi) FindTargetExamPlans(c *gin.Context) {
	var teachClassId request2.GetByTeachClassId
	_ = c.ShouldBindQuery(&teachClassId)
	sId := utils.GetStudentId(c)
	if examPlans, err := examService.FindTargetExamPlans(teachClassId.TeachClassId, sId); err != nil {
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
	ip := c.ClientIP()
	if bool, _ := examPlanService.CheckIsExamSt(planId.PlanId, studentId); bool == false {
		response.FailWithMessage("你不在参与此考试的范围中", c)
	} else {
		if isFinishPreExam, err, preExamIds := examPlanService.IsFinishPreExam(planId.PlanId, studentId); err != nil {
			response.FailWithMessage("前置计划查询出错", c)
		} else if isFinishPreExam == false {
			response.FailWithDetailedAndError(704, preExamIds, "请先完成并通过前置计划", c)
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
				if isReady, _ := examService.CheckIsReady(planId.PlanId); isReady == false {
					if exampaper, status, examScore, err := examService.GetExamPapersBySql(examComing, ip); err != nil {
						global.GVA_LOG.Error("查询考试试卷失败", zap.Error(err))
						response.FailWithMessage("查询考试试卷失败", c)
					} else if status.IsCommit && PlanDetail.Type == examType.FinalExam {
						response.FailWithMessageAndError(703, "你已经提交过且通过该考试", c)
					} else if status.IsCommit && PlanDetail.Type == examType.ProceduralExam && *examScore.Score >= *PlanDetail.PassScore {
						response.FailWithMessageAndError(703, "你已经提交过且通过该考试", c)
					} else {
						var remainTime float64
						time1 := PlanDetail.EndTime.Sub(time.Now()).Minutes()
						if time1 < float64(*PlanDetail.Time) {
							remainTime = time1
						} else {
							remainTime = float64(*PlanDetail.Time)
						}
						response.OkWithData(gin.H{
							"examPaper":  exampaper,
							"enterTime":  status.EnterTime,
							"timeRemain": remainTime,
						}, c)
					}
				} else {
					if examPaperId, singleChoice, multiChoice, judge, blank, program, status, examScore, err := examService.GetExamPapers(examComing, ip); err != nil {
						global.GVA_LOG.Error("查询考试试卷失败", zap.Error(err))
						response.FailWithMessage("查询考试试卷失败", c)
					} else if status.IsCommit && PlanDetail.Type == examType.FinalExam {
						response.FailWithMessageAndError(703, "你已经提交过且通过该考试", c)
					} else if status.IsCommit && PlanDetail.Type == examType.ProceduralExam && *examScore.Score >= *PlanDetail.PassScore {
						response.FailWithMessageAndError(703, "你已经提交过了且通过该考试", c)
					} else {
						response.OkWithData(gin.H{
							"examPaperId":           examPaperId,
							"singleChoiceComponent": singleChoice,
							"multiChoiceComponent":  multiChoice,
							"judgeComponent":        judge,
							"blankComponent":        blank,
							"programComponent":      program,
							"enterTime":             status.EnterTime,
						}, c)
					}
				}
			}
		}
	}
}

//保存试卷内容
func (examApi *ExamApi) SaveExamPaper(c *gin.Context) {
	var ExamCommit examManage.CommitExamPaper2
	_ = c.ShouldBindJSON(&ExamCommit)
	ExamCommit.StudentId = utils.GetStudentId(c)
	PlanDetail, _ := examPlanService.GetExamPlan(ExamCommit.PlanId)
	status, _ := statuServie.GetStatus(ExamCommit.StudentId, ExamCommit.PlanId)
	examScore, _ := statuServie.GetScore(ExamCommit.StudentId, ExamCommit.PlanId)
	minutes := *PlanDetail.Time
	unix1 := status.EnterTime.Add(time.Duration(minutes) * time.Minute)
	if PlanDetail.IsLimitTime == true && time.Now().Unix() > unix1.Unix() {
		response.FailWithMessageAndError(704, "超出考试时间", c)
	} else if time.Now().Unix() > PlanDetail.EndTime.Unix() {
		response.FailWithMessageAndError(704, "保存失败,考试已经结束了", c)
	} else if status.IsCommit && PlanDetail.Type == examType.FinalExam {
		response.FailWithMessageAndError(703, "你已经提交过且通过该考试", c)
	} else if status.IsCommit && PlanDetail.Type == examType.ProceduralExam && *examScore.Score >= *PlanDetail.PassScore {
		response.FailWithMessageAndError(703, "你已经提交过了且通过该考试", c)
	} else {
		if err := examService.SaveExamPapers(ExamCommit); err != nil {
			global.GVA_LOG.Error("试卷保存失败", zap.Error(err))
			response.FailWithMessage("试卷提交失败", c)
		} else {
			response.OkWithMessage("试卷保存成功", c)
		}
	}
}
func (examApi *ExamApi) FindSaveExamPaper(c *gin.Context) {
	var ExamCommit examManage.CommitExamPaper
	_ = c.ShouldBindQuery(&ExamCommit)
	ExamCommit.StudentId = utils.GetStudentId(c)
	//if AllMergeId, err := examService.GetAllQues(ExamCommit.PlanId, ExamCommit.StudentId); err != nil {
	//	global.GVA_LOG.Error("试卷保存失败", zap.Error(err))
	//	response.FailWithMessage("试卷提交失败", c)
	//} else {
	if SavePaper, err := examService.GetAllQuesAnswer(ExamCommit.PlanId, ExamCommit.StudentId); err != nil {
		global.GVA_LOG.Error("获取保存的试卷失败", zap.Error(err))
		response.FailWithMessage("获取保存的试卷失败", c)
	} else {
		response.OkWithData(gin.H{
			"Answer":  SavePaper,
			"message": "获取成功",
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
	examScore, _ := statuServie.GetScore(ExamCommit.StudentId, ExamCommit.PlanId)
	minutes := *PlanDetail.Time
	unix1 := status.EnterTime.Add(time.Duration(minutes) * time.Minute)
	if bool, _ := examPlanService.CheckIsExamSt(ExamCommit.PlanId, ExamCommit.StudentId); bool == false {
		response.FailWithMessage("你不在参与此考试的范围中", c)
	} else {
		if PlanDetail.IsLimitTime == true && time.Now().Unix() > unix1.Unix() {
			response.FailWithMessageAndError(704, "超出考试时间", c)
		} else if time.Now().Unix() > PlanDetail.EndTime.Unix() {
			response.FailWithMessageAndError(704, "提交失败,考试已经结束了", c)
		} else if status.IsCommit && PlanDetail.Type == examType.FinalExam {
			response.FailWithMessageAndError(703, "你已经提交过且通过该考试", c)
		} else if status.IsCommit && PlanDetail.Type == examType.ProceduralExam && *examScore.Score >= *PlanDetail.PassScore {
			response.FailWithMessageAndError(703, "你已经提交过了且通过该考试", c)
		} else {
			if err := examService.CommitExamPapers(ExamCommit); err != nil {
				global.GVA_LOG.Error("试卷提交失败", zap.Error(err))
				response.FailWithMessage("试卷提交失败", c)
			} else {
				go func() {
					global.GVA_LOG.Info("start,开始处理试卷")
					time.AfterFunc(time.Second*5, func() {
						wg.Add(1)
						if err = examService.UpdateExamPapers(ExamCommit); err != nil {
							global.GVA_LOG.Error("更新试卷记录失败", zap.Error(err))
						}
						utils.ExecPapers(ExamCommit)
						defer wg.Done()
					})
					wg.Wait()
				}()
				response.OkWithData(gin.H{"examPaper": ExamCommit}, c)
			}
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
	errChan := make(chan error)
	go func() {
		checkProgram, score, _, e := programService.CheckProgram(program.QuestionId, program.Code, program.LanguageId)
		errChan <- e
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
			err = <-errChan
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

// 获取考试分数
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
