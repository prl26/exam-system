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

type SupplyBlankApi struct {
}

var supplyBlankService = service.ServiceGroupApp.QuestionBankServiceGroup.SupplyBlankService

// Create 创建填空题
func (api *SupplyBlankApi) Create(c *gin.Context) {
	var req questionBankReq.SupplyBlankCreate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProblemType": {utils.NotEmpty()},
		"CanPractice": {utils.NotEmpty()},
		"CanExam":     {utils.NotEmpty()},
		"Title":       {utils.NotEmpty()},
		"Describe":    {utils.NotEmpty()},
		"Answers":     {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	supplyBlank := questionBankPo.SupplyBlank{}
	supplyBlank.BasicModel = req.BasicModel
	supplyBlank.IsOrder = req.IsOrder
	supplyBlank.SupplyBlankModel = req.SupplyBlankModel
	if a, b, err := req.Answers.GetAnswersAndProportions(); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		supplyBlank.Answer = a
		supplyBlank.Proportion = b
		num := len(req.Answers)
		supplyBlank.Num = &num
	}
	if err := supplyBlankService.Create(&supplyBlank); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("创建失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("创建成功", c)
	}
}

// Delete 删除填空题
func (api *SupplyBlankApi) Delete(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := supplyBlankService.Delete(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新填空题
func (api *SupplyBlankApi) Update(c *gin.Context) {
	var req questionBankReq.SupplyBlankUpdate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	supplyBlank := questionBankPo.SupplyBlank{}
	supplyBlank.ID = req.Id
	supplyBlank.SupplyBlankModel = req.SupplyBlankModel
	if a, b, err := req.Answers.GetAnswersAndProportions(); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		supplyBlank.Answer = a
		supplyBlank.Proportion = b
		num := len(req.Answers)
		supplyBlank.Num = &num
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
		return
	}
	if err := supplyBlankService.UpdateQuestionBankSupplyBlank(supplyBlank); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("更新失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("更新成功", c)
	}
}

// FindList  分页查找填空题
func (api *SupplyBlankApi) FindList(c *gin.Context) {
	var pageInfo questionBankReq.QuestionBankSupplyBlankSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := supplyBlankService.FindList(pageInfo); err != nil {
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

// FindDetail  获取填空题详细
func (api *SupplyBlankApi) FindDetail(c *gin.Context) {
	var req questionBankReq.DetailFind
	_ = c.ShouldBindQuery(&req)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}

	if data, err := supplyBlankService.FindDetail(req.Id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
	} else {
		if data == nil {
			questionBankResp.NotFind(c)
			return
		}
		//global.GVA_MODEL
		//	Chapter   basicdata.Chapter
		//	Knowledge basicdata.Knowledge
		//	questionBankPo.BasicModel
		detail := questionBankResp.SupplyBlankDetail{}
		detail.Answers.Deserialization(data.Answer, data.Proportion)
		detail.GVA_MODEL = data.GVA_MODEL
		detail.Chapter = data.Chapter
		detail.Knowledge = data.Knowledge
		detail.SupplyBlankModel = data.SupplyBlankModel
		questionBankResp.OkWithDetailed(detail, "获取成功", c)
	}
}
