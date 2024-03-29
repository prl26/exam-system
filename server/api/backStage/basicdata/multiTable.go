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
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
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
	courseId := int(stuClassReq.CourseId)

	n := len(stuClassReq.StudentIds)
	students := make([]*basicdata.Student, 0, n)
	news := make([]*teachplan.Score, 0, n)

	for i := 0; i < n; i++ {
		student := &basicdata.Student{}
		scoreStudent := &teachplan.Score{}

		student.ID = stuClassReq.StudentIds[i]
		student.Name = stuClassReq.StudentName

		id := int(stuClassReq.StudentIds[i])
		// 先添加学生
		students = append(students, student)

		scoreStudent.StudentId = &id
		scoreStudent.TeachClassId = &tid
		scoreStudent.TermId = &termId
		scoreStudent.CourseId = &courseId

		// 判断学生是否存在
		user := studentService.QueryStudentById(student.ID)

		if user.ID != 0 {
			score := scoreService.QueryScoreByStudent(scoreStudent)
			if score.ID == 0 {
				// 不存在这个 数据 则创建成绩索引
				news = append(news, scoreStudent)
			}
		} else {
			//response.FailWithMessage("学生不存在，请先创建学生", c)
			student.Password = utils.BcryptHash(strconv.Itoa(int(student.ID)))
			_ = studentService.CreateStudentX(student)
			news = append(news, scoreStudent)
		}

	}

	err := multiTableService.AssociationStudents(students, stuClassReq.TeachClassId)
	_ = scoreService.CreateScores(news)

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
