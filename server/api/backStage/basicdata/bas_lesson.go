package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"go.uber.org/zap"
)

type LessonApi struct {
}

var lessonService = service.ServiceGroupApp.BasicdataApiGroup.LessonService

// CreateLesson 创建Lesson
// @Tags Lesson
// @Summary 创建Lesson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Lesson true "创建Lesson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lesson/createLesson [post]
func (lessonApi *LessonApi) CreateLesson(c *gin.Context) {
	var lesson basicdata.Lesson
	_ = c.ShouldBindJSON(&lesson)
	lesson.OpenQuestionBank = true
	if err := lessonService.CreateLesson(lesson); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLesson 删除Lesson
// @Tags Lesson
// @Summary 删除Lesson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Lesson true "删除Lesson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lesson/deleteLesson [delete]
func (lessonApi *LessonApi) DeleteLesson(c *gin.Context) {
	var lesson basicdata.Lesson
	_ = c.ShouldBindJSON(&lesson)
	if err := lessonService.DeleteLesson(lesson); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLessonByIds 批量删除Lesson
// @Tags Lesson
// @Summary 批量删除Lesson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Lesson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lesson/deleteLessonByIds [delete]
func (lessonApi *LessonApi) DeleteLessonByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := lessonService.DeleteLessonByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLesson 更新Lesson
// @Tags Lesson
// @Summary 更新Lesson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Lesson true "更新Lesson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lesson/updateLesson [put]
func (lessonApi *LessonApi) UpdateLesson(c *gin.Context) {
	var lesson basicdata.Lesson
	_ = c.ShouldBindJSON(&lesson)
	if err := lessonService.UpdateLesson(lesson); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLesson 用id查询Lesson
// @Tags Lesson
// @Summary 用id查询Lesson
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.Lesson true "用id查询Lesson"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lesson/findLesson [get]
func (lessonApi *LessonApi) FindLesson(c *gin.Context) {
	var lesson basicdata.Lesson
	_ = c.ShouldBindQuery(&lesson)
	if relesson, err := lessonService.GetLesson(lesson.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relesson": relesson}, c)
	}
}

// GetLessonList 分页获取Lesson列表
// @Tags Lesson
// @Summary 分页获取Lesson列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.LessonSearch true "分页获取Lesson列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lesson/getLessonList [get]
func (lessonApi *LessonApi) GetLessonList(c *gin.Context) {
	var pageInfo basicdataReq.LessonSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := lessonService.GetLessonInfoList(pageInfo); err != nil {
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
