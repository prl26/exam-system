package questionBank

import (
	"github.com/gin-gonic/gin"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
)

type OjApi struct {
}

var cService = &service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService
var goService = &service.ServiceGroupApp.OjServiceServiceGroup.GoLanguage
var programOjService = &service.ServiceGroupApp.OjServiceServiceGroup.ProgramService

func (p *OjApi) Compile(c *gin.Context) {
	var req questionBankReq.Compile
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"Code":       {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	compile, t, err := programOjService.Compile(req.Code, req.LanguageId)
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	} else {
		questionBankResp.OkWithDetailed(questionBankResp.Compile{
			FileId:         compile,
			ExpirationTime: *t,
		}, "编译成功", c)
	}
}

func (p *OjApi) Execute(c *gin.Context) {
	var req questionBankReq.Execute
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"FileId":     {utils.NotEmpty()},
		"LanguageId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if output, executeSituation, err := programOjService.Execute(req.LanguageId, req.FileId, req.Input, req.LanguageLimit); err != nil {
		questionBankResp.ErrorHandle(c, err)
	} else {
		questionBankResp.OkWithDetailed(questionBankResp.Execute{
			Output:           output,
			ExecuteSituation: *executeSituation,
		}, "获取运行结果成功", c)
	}
}
