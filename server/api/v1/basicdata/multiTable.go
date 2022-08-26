/*
*

	@author: qianyi  2022/8/24 19:00:00
	@note:
*/
package basicdata

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type MultiTableApi struct {
}

var multiTableService = service.ServiceGroupApp.BasicdataApiGroup.MultiTableService

// InitTeachClassStudent 教学班中 添加学生
// @Tags TeachClassStudent
// @Summary 教学班中 添加学生
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StuTeachClass true "添加TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/initTeachClassStudent [post]
func (multiTableServiceApi *MultiTableApi) InitTeachClassStudent(c *gin.Context) {
	var stuClassReq basicdataReq.StuTeachClass
	_ = c.ShouldBindJSON(&stuClassReq)
	err := multiTableService.InitTeachClassStudents(stuClassReq)
	if err != nil {
		global.GVA_LOG.Error("教学班中添加学生失败", zap.Error(err))
		response.FailWithMessage("教学班中添加学生失败", c)
	} else {
		response.OkWithMessage("教学班中添加学生成功", c)
	}
}

// DeleteTeachClassStudent 教学班 中移除学生
// @Tags TeachClassStudent
// @Summary 教学班 中移除学生
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StuTeachClass true "移除TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/deleteTeachClassStudent [post]
func (multiTableServiceApi *MultiTableApi) DeleteTeachClassStudent(c *gin.Context) {
	var stuClassReq basicdataReq.StuTeachClass
	_ = c.ShouldBindJSON(&stuClassReq)
	err := multiTableService.DeleteTeachClassStudents(stuClassReq)
	if err != nil {
		global.GVA_LOG.Error("教学班中移除学生失败", zap.Error(err))
		response.FailWithMessage("教学班中移除学生失败", c)
	} else {
		response.OkWithMessage("教学班中移除学生成功", c)
	}
}

// GetTeachClassStudentList 获取教学班中 学生列表
// @Tags TeachClassStudent
// @Summary 获取教学班中 学生列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.TeachClassStudent true "获取TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/getTeachClassStudentList [get]
func (multiTableServiceApi *MultiTableApi) GetTeachClassStudentList(c *gin.Context) {
	var pageInfo basicdataReq.TeachClassStudent
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		fmt.Println(err)
	}
	if list, total, err := multiTableService.GetTeachClassStudentInfo(pageInfo); err != nil {
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
