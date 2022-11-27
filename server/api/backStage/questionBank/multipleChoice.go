package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type MultipleChoiceApi struct {
}

var multipleChoiceService = service.ServiceGroupApp.QuestionBankServiceGroup.MultipleChoiceService

// Create 创建选择题
func (choiceApi *MultipleChoiceApi) Create(c *gin.Context) {
	var req questionBankReq.MultipleChoiceCreate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProblemType": {utils.NotEmpty()},
		"CanPractice": {utils.NotEmpty()},
		"CanExam":     {utils.NotEmpty()},
		"Title":       {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := multipleChoiceService.Create(&req.MultipleChoice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("创建失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("创建成功", c)
	}
}

// Delete 删除选择题
func (choiceApi *MultipleChoiceApi) Delete(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := multipleChoiceService.Delete(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("批量删除失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新选择题
func (choiceApi *MultipleChoiceApi) Update(c *gin.Context) {
	var req questionBankReq.MultipleChoiceUpdate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := multipleChoiceService.Update(req.MultipleChoice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("更新错误:%s", err.Error()))
		return
	} else {
		questionBankResp.OkWithMessage("更新成功", c)
	}
}

// FindDetail  获取选择题详细
func (choiceApi *MultipleChoiceApi) FindDetail(c *gin.Context) {
	var req questionBankReq.DetailFind
	_ = c.ShouldBindQuery(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if data, err := multipleChoiceService.FindDetail(req.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("查询错误:%s", err.Error()))
		return
	} else {
		if data != nil {
			questionBankResp.OkWithDetailed(data, "获取成功", c)
		} else {
			questionBankResp.NotFind(c)
		}
	}
}

// FindSingleChoice  分页查找单选题
func (choiceApi *MultipleChoiceApi) FindSingleChoice(c *gin.Context) {
	var pageInfo questionBankReq.MultipleChoiceList
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := multipleChoiceService.FindList(pageInfo.MultipleCriteria, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error(err.Error())
		questionBankResp.ErrorHandle(c, fmt.Errorf("查询错误:%s", err.Error()))
		return
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//// FindMultipleChoice 分页查找多选题
//func (choiceApi *MultipleChoiceApi) FindMultipleChoice(c *gin.Context) {
//	var pageInfo request2.MultipleChoiceList
//	_ = c.ShouldBindQuery(&pageInfo)
//	if list, total, err := multipleChoiceService.FindList(pageInfo, true); err != nil {
//		global.GVA_LOG.Error("获取失败!", zap.Error(err))
//		response.FailWithMessage("获取失败", c)
//	} else {
//		response.OkWithDetailed(response.PageResult{
//			List:     list,
//			Total:    total,
//			Page:     pageInfo.Page,
//			PageSize: pageInfo.PageSize,
//		}, "获取成功", c)
//	}
//}
