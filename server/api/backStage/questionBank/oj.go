package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	questionBankEnum "github.com/prl26/exam-system/server/model/questionBank/enum"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
)

type OjApi struct {
}

var cService = service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService

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
	switch req.LanguageId {
	case questionBankEnum.C_LANGUAGE:
		compile, t, err := cService.Compile(req.Code)
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		questionBankResp.OkWithDetailed(questionBankResp.Compile{
			FileId:         compile,
			ExpirationTime: *t,
		}, "编译成功", c)
	default:
		questionBankResp.CheckHandle(c, fmt.Errorf("编程语言输入错误"))
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
	switch req.LanguageId {
	case questionBankEnum.C_LANGUAGE:
		compile, t, err := cService.Execute(req.FileId, req.Input, req.LanguageLimit)
		if err != nil {
			questionBankResp.ErrorHandle(c, err)
			return
		}
		questionBankResp.OkWithDetailed(questionBankResp.Execute{
			Output:           compile,
			ExecuteSituation: *t,
		}, "获取运行结果成功", c)
	default:
		questionBankResp.CheckHandle(c, fmt.Errorf("编程语言输入错误"))
	}
}
