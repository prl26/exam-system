package questionBank

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/model/common/response"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
)

type OjApi struct {
}

//
var (
	judgeService          = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.JudgeService
	programService        = &service.ServiceGroupApp.QuestionBankServiceGroup.OjService.ProgramService
	supplyBlankService    = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.SupplyBlankService
	multipleChoiceService = service.ServiceGroupApp.QuestionBankServiceGroup.OjService.MultipleChoiceService
)

func (*OjApi) CheckJudge(c *gin.Context) {
	var r questionBankReq.CheckJudge
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"Id": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, err := judgeService.Check(r.Id, r.Answer)
	if err != nil {
		response.FailWithMessage("找不到该题目", c)
		return
	}
	response.OkWithData(result, c)
	return
}

func (*OjApi) CheckProgram(c *gin.Context) {
	var r questionBankReq.CheckProgramm
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"Id":         {utils.NotEmpty()},
		"Code":       {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	}

	program, _, err := programService.CheckProgram(r.Id, r.Code, r.LanguageId)
	if err != nil {
		questionBankResp.ErrorHandle(c, err)
		return
	} else {
		response.OkWithData(program, c)
	}
}

func (*OjApi) CheckSupplyBlank(c *gin.Context) {
	var r questionBankReq.CheckSupplyBlank
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"Id":      {utils.NotEmpty()},
		"Answers": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, _, err := supplyBlankService.Check(r.Id, r.Answers)
	if err != nil {
		response.FailWithMessage("找不到该题目", c)
		return
	}
	response.OkWithData(result, c)
	return
}

func (*OjApi) CheckMultipleChoice(c *gin.Context) {
	var r questionBankReq.CheckMultipleChoice
	_ = c.ShouldBindJSON(&r)
	verify := utils.Rules{
		"Id":      {utils.NotEmpty()},
		"Answers": {utils.NotEmpty()},
	}
	if err := utils.Verify(r, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, err := multipleChoiceService.Check(r.Id, r.Answers)
	if err != nil {
		response.FailWithMessage("找不到该题目", c)
		return
	}
	response.OkWithData(result, c)
	return
}
