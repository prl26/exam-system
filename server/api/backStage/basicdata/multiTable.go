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
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type MultiTableApi struct {
}

var multiTableService = service.ServiceGroupApp.BasicdataApiGroup.MultiTableService
var scoreService = service.ServiceGroupApp.TeachplanServiceGroup.ScoreService

// InitTeachClassStudent 教学班中 添加学生
// @Tags TeachClassStudent
// @Summary 教学班中 添加学生
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StuTeachClass true "添加TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /teachClassStudent/initTeachClassStudent [post]
func (multiTableServiceApi *MultiTableApi) InitTeachClassStudent(c *gin.Context) {
	var stuClassReq basicdataReq.StuTeachClass
	_ = c.ShouldBindJSON(&stuClassReq)

	tid := int(stuClassReq.TeachClassId)
	termId := int(stuClassReq.TermId)

	n := len(stuClassReq.StudentIds)
	students := make([]*basicdata.Student, n)
	scoreStudents := make([]*teachplan.Score, n)
	new := make([]*teachplan.Score, n)

	for i := 0; i < n; i++ {
		id := int(stuClassReq.StudentIds[i])
		students[i].ID = stuClassReq.StudentIds[i]
		scoreStudents[i].StudentId = &id
		scoreStudents[i].TeachClassId = &tid
		scoreStudents[i].TermId = &termId

		score := scoreService.QueryScoreByStudent(scoreStudents[i])
		if score.StudentId == nil && score.TeachClassId == nil {
			// 不存在这个 数据 则创建索引
			new = append(new, scoreStudents[i])
		}
	}

	err := multiTableService.InitTeachClassStudents(stuClassReq.TeachClassId, students)
	_ = scoreService.CreateScores(new)

	if err != nil {
		global.GVA_LOG.Error("添加学生失败", zap.Error(err))
		response.FailWithMessage("添加学生失败", c)
	} else {
		response.OkWithMessage("添加学生成功", c)
	}
}

// AddStudentByClass 教学班中 添加指定班级的所有学生
// @Tags TeachClassStudent
// @Summary 教学班中 添加指定班级的所有学生
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddStudentByClass true "添加TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /teachClassStudent/addStudentByClass [post]
func (multiTableServiceApi *MultiTableApi) AddStudentByClass(c *gin.Context) {
	var param basicdataReq.AddStudentByClass
	_ = c.ShouldBindJSON(&param)
	err := multiTableService.AddStudentByClass(param)
	if err != nil {
		global.GVA_LOG.Error("添加学生失败", zap.Error(err))
		response.FailWithMessage("添加学生失败", c)
	} else {
		response.OkWithMessage("添加学生成功", c)
	}
}

// DeleteStudentByClass 教学班中 移除指定班级的所有学生
// @Tags TeachClassStudent
// @Summary 教学班中 移除指定班级的所有学生
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddStudentByClass true "添加TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachClassStudent/deleteStudentByClass [post]
func (multiTableServiceApi *MultiTableApi) DeleteStudentByClass(c *gin.Context) {
	var param basicdataReq.AddStudentByClass
	_ = c.ShouldBindJSON(&param)
	err := multiTableService.DeleteStudentByClass(param)
	if err != nil {
		global.GVA_LOG.Error("移除学生失败", zap.Error(err))
		response.FailWithMessage("移除学生失败", c)
	} else {
		response.OkWithMessage("移除学生成功", c)
	}
}

// DeleteTeachClassStudent 教学班中 移除学生
// @Tags TeachClassStudent
// @Summary 教学班中 移除学生
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StuTeachClass true "移除TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"移除成功"}"
// @Router /teachClassStudent/deleteTeachClassStudent [post]
func (multiTableServiceApi *MultiTableApi) DeleteTeachClassStudent(c *gin.Context) {
	var stuClassReq basicdataReq.StuTeachClass
	_ = c.ShouldBindJSON(&stuClassReq)
	err := multiTableService.DeleteTeachClassStudents(stuClassReq)
	if err != nil {
		global.GVA_LOG.Error("移除学生失败", zap.Error(err))
		response.FailWithMessage("移除学生失败", c)
	} else {
		response.OkWithMessage("移除学生成功", c)
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
