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

 * @Date:   2022/8/26 20:51

 * @Note:

 **/

type MultipleChoiceApi struct{}

var multipleChoiceService = service.ServiceGroupApp.OjServiceServiceGroup.MultipleChoiceService

// CheckMultipleChoice 检验选择题
// @Tags OJ
// @Summary 检验选择题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CheckMultipleChoice true "检验选择题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /oj/multipleChoice/check [post]
func (*MultipleChoiceApi) CheckMultipleChoice(c *gin.Context) {
	var r ojReq.CheckMultipleChoice
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
