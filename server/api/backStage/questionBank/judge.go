package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"

	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
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
		questionBankResp.CheckHandle(c, err)
		return
	}
	if err := judgeService.Create(&req.Judge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("创建失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("创建成功", c)
	}
}

// Delete 删除判断题
func (api *JudgeApi) Delete(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := judgeService.Delete(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("批量删除失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新判断题
func (api *JudgeApi) Update(c *gin.Context) {
	var req questionBankPo.Judge
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
		questionBankResp.ErrorHandle(c, fmt.Errorf("更新失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("更新成功", c)
	}
}

// FindList  分页查找判断题
func (api *JudgeApi) FindJudgeList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankJudgeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := judgeService.FindJudgeList(pageInfo.JudgeSearchCriteria, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
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
		questionBankResp.CheckHandle(c, err)
		return
	}

	if judge, err := judgeService.FindDetail(req.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
		return
	} else {
		questionBankResp.OkWithDetailed(judge, "获取成功", c)
	}
}
