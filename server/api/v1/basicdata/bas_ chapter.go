package basicdata

import (
	"exam-system/global"
	"exam-system/model/basicdata"
	basicdataReq "exam-system/model/basicdata/request"
	"exam-system/model/common/request"
	"exam-system/model/common/response"
	"exam-system/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChapterApi struct {
}

var chapterService = service.ServiceGroupApp.BasicdataApiGroup.ChapterService

// CreateChapter 创建Chapter
// @Tags Chapter
// @Summary 创建Chapter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Chapter true "创建Chapter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chapter/createChapter [post]
func (chapterApi *ChapterApi) CreateChapter(c *gin.Context) {
	var chapter basicdata.Chapter
	_ = c.ShouldBindJSON(&chapter)
	if err := chapterService.CreateChapter(chapter); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChapter 删除Chapter
// @Tags Chapter
// @Summary 删除Chapter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Chapter true "删除Chapter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chapter/deleteChapter [delete]
func (chapterApi *ChapterApi) DeleteChapter(c *gin.Context) {
	var chapter basicdata.Chapter
	_ = c.ShouldBindJSON(&chapter)
	if err := chapterService.DeleteChapter(chapter); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChapterByIds 批量删除Chapter
// @Tags Chapter
// @Summary 批量删除Chapter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Chapter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /chapter/deleteChapterByIds [delete]
func (chapterApi *ChapterApi) DeleteChapterByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := chapterService.DeleteChapterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChapter 更新Chapter
// @Tags Chapter
// @Summary 更新Chapter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Chapter true "更新Chapter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chapter/updateChapter [put]
func (chapterApi *ChapterApi) UpdateChapter(c *gin.Context) {
	var chapter basicdata.Chapter
	_ = c.ShouldBindJSON(&chapter)
	if err := chapterService.UpdateChapter(chapter); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChapter 用id查询Chapter
// @Tags Chapter
// @Summary 用id查询Chapter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.Chapter true "用id查询Chapter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chapter/findChapter [get]
func (chapterApi *ChapterApi) FindChapter(c *gin.Context) {
	var chapter basicdata.Chapter
	_ = c.ShouldBindQuery(&chapter)
	if rechapter, err := chapterService.GetChapter(chapter.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rechapter": rechapter}, c)
	}
}

// GetChapterList 分页获取Chapter列表
// @Tags Chapter
// @Summary 分页获取Chapter列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.ChapterSearch true "分页获取Chapter列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chapter/getChapterList [get]
func (chapterApi *ChapterApi) GetChapterList(c *gin.Context) {
	var pageInfo basicdataReq.ChapterSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := chapterService.GetChapterInfoList(pageInfo); err != nil {
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
