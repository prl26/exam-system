package oj

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/model/common/response"
	ojReq "github.com/prl26/exam-system/server/model/oj/request"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/26 20:57

 * @Note:

 **/

type JudgeApi struct{}

var judgeService = service.ServiceGroupApp.OjServiceServiceGroup.JudgeService

// CheckJudge 检验判断题
// @Tags OJ
// @Summary 检验判断题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CheckJudge true "检验选择题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oj/judge/check [post]
func (*JudgeApi) CheckJudge(c *gin.Context) {
	var r ojReq.CheckJudge
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
