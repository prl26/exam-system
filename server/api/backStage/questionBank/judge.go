package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/questionBank"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type JudgeApi struct {
}

var judgeService = service.ServiceGroupApp.QuestionBankServiceGroup.JudgeService

// Create 创建判断题
func (api *JudgeApi) Create(c *gin.Context) {
	var req questionBankReq.JudgeCreate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProblemType": {utils.NotEmpty()},
		"CanPractice": {utils.NotEmpty()},
		"CanExam":     {utils.NotEmpty()},
		"Title":       {utils.NotEmpty()},
		"Describe":    {utils.NotEmpty()},
		"IsRight":     {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := judgeService.Create(&req.Judge, req.LessonSupports); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除判断题
func (api *JudgeApi) Delete(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := judgeService.Delete(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新判断题
func (api *JudgeApi) Update(c *gin.Context) {
	var req questionBank.Judge
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := judgeService.Update(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindList  分页查找判断题
func (api *JudgeApi) FindJudgeList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankJudgeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := judgeService.FindJudgeList(pageInfo); err != nil {
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

// FindDetail  获取判断题详细
func (api *JudgeApi) FindDetail(c *gin.Context) {
	var req questionBankReq.DetailFind
	_ = c.ShouldBindQuery(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resp := questionBankResp.JudgeDetail{}
	if err := judgeService.FindDetail(&resp.Judge, req.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	if resp.Judge.ID != 0 {
		if err := questionBankService.FindCourseSupport(&resp.CourseSupport, resp.Judge.ID, questionType.JUDGE); err != nil {
			global.GVA_LOG.Error("获取语言支持失败", zap.Error(err))
			response.FailWithMessage("获取语言支持失败", c)
			return
		}
		response.OkWithData(resp, c)
	} else {
		response.FailWithMessage("找不到该编程题", c)
		return
	}
}
