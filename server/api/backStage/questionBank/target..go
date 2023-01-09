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

type TargetApi struct {
}

var TargetService = service.ServiceGroupApp.QuestionBankServiceGroup.TargetService

// Create 创建靶场题
func (api *TargetApi) Create(c *gin.Context) {
	var req questionBankReq.TargetCreate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProblemType": {utils.NotEmpty()},
		"CanPractice": {utils.NotEmpty()},
		"CanExam":     {utils.NotEmpty()},
		"Title":       {utils.NotEmpty()},
		"Describe":    {utils.NotEmpty()},
		"ByteCode":    {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if err := TargetService.Create(&req.Target); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("创建失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("创建成功", c)
	}
}

// Delete 删除靶场题
func (api *TargetApi) Delete(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := TargetService.Delete(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("批量删除失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新靶场题
func (api *TargetApi) Update(c *gin.Context) {
	var req questionBankPo.Target
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TargetService.Update(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("更新失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("更新成功", c)
	}
}

// FindList  分页查找靶场题
func (api *TargetApi) FindList(c *gin.Context) {
	var pageInfo questionBankReq.TargetSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := TargetService.FindTargetList(pageInfo.TargetSearchCriteria, pageInfo.PageInfo); err != nil {
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

// FindDetail  获取靶场题详细
func (api *TargetApi) FindDetail(c *gin.Context) {
	var req questionBankReq.DetailFind
	_ = c.ShouldBindQuery(&req)
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}

	if Target, err := TargetService.FindDetail(req.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
		return
	} else {
		questionBankResp.OkWithDetailed(Target, "获取成功", c)
	}
}
