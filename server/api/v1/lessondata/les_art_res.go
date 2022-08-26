package lessondata

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/common/response"
	"exam-system/model/lessondata"
	lessondataReq "exam-system/model/lessondata/request"
	"exam-system/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleResourcesApi struct {
}

var articleResourcesService = service.ServiceGroupApp.LessondataServiceGroup.ArticleResourcesService

// CreateArticleResources 创建ArticleResources
// @Tags ArticleResources
// @Summary 创建ArticleResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ArticleResources true "创建ArticleResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /articleResources/createArticleResources [post]
func (articleResourcesApi *ArticleResourcesApi) CreateArticleResources(c *gin.Context) {
	var articleResources lessondata.ArticleResources
	_ = c.ShouldBindJSON(&articleResources)
	if err := articleResourcesService.CreateArticleResources(articleResources); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticleResources 删除ArticleResources
// @Tags ArticleResources
// @Summary 删除ArticleResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ArticleResources true "删除ArticleResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /articleResources/deleteArticleResources [delete]
func (articleResourcesApi *ArticleResourcesApi) DeleteArticleResources(c *gin.Context) {
	var articleResources lessondata.ArticleResources
	_ = c.ShouldBindJSON(&articleResources)
	if err := articleResourcesService.DeleteArticleResources(articleResources); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticleResourcesByIds 批量删除ArticleResources
// @Tags ArticleResources
// @Summary 批量删除ArticleResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ArticleResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /articleResources/deleteArticleResourcesByIds [delete]
func (articleResourcesApi *ArticleResourcesApi) DeleteArticleResourcesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := articleResourcesService.DeleteArticleResourcesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArticleResources 更新ArticleResources
// @Tags ArticleResources
// @Summary 更新ArticleResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lessondata.ArticleResources true "更新ArticleResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /articleResources/updateArticleResources [put]
func (articleResourcesApi *ArticleResourcesApi) UpdateArticleResources(c *gin.Context) {
	var articleResources lessondata.ArticleResources
	_ = c.ShouldBindJSON(&articleResources)
	if err := articleResourcesService.UpdateArticleResources(articleResources); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticleResources 用id查询ArticleResources
// @Tags ArticleResources
// @Summary 用id查询ArticleResources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondata.ArticleResources true "用id查询ArticleResources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /articleResources/findArticleResources [get]
func (articleResourcesApi *ArticleResourcesApi) FindArticleResources(c *gin.Context) {
	var articleResources lessondata.ArticleResources
	_ = c.ShouldBindQuery(&articleResources)
	if rearticleResources, err := articleResourcesService.GetArticleResources(articleResources.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearticleResources": rearticleResources}, c)
	}
}

// GetArticleResourcesList 分页获取ArticleResources列表
// @Tags ArticleResources
// @Summary 分页获取ArticleResources列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lessondataReq.ArticleResourcesSearch true "分页获取ArticleResources列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /articleResources/getArticleResourcesList [get]
func (articleResourcesApi *ArticleResourcesApi) GetArticleResourcesList(c *gin.Context) {
	var pageInfo lessondataReq.ArticleResourcesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := articleResourcesService.GetArticleResourcesInfoList(pageInfo); err != nil {
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
