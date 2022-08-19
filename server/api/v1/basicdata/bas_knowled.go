package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type KnowledgeApi struct {
}

var knowledgeService = service.ServiceGroupApp.BasicdataApiGroup.KnowledgeService


// CreateKnowledge 创建Knowledge
// @Tags Knowledge
// @Summary 创建Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Knowledge true "创建Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /knowledge/createKnowledge [post]
func (knowledgeApi *KnowledgeApi) CreateKnowledge(c *gin.Context) {
	var knowledge basicdata.Knowledge
	_ = c.ShouldBindJSON(&knowledge)
	if err := knowledgeService.CreateKnowledge(knowledge); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteKnowledge 删除Knowledge
// @Tags Knowledge
// @Summary 删除Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Knowledge true "删除Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /knowledge/deleteKnowledge [delete]
func (knowledgeApi *KnowledgeApi) DeleteKnowledge(c *gin.Context) {
	var knowledge basicdata.Knowledge
	_ = c.ShouldBindJSON(&knowledge)
	if err := knowledgeService.DeleteKnowledge(knowledge); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteKnowledgeByIds 批量删除Knowledge
// @Tags Knowledge
// @Summary 批量删除Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /knowledge/deleteKnowledgeByIds [delete]
func (knowledgeApi *KnowledgeApi) DeleteKnowledgeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := knowledgeService.DeleteKnowledgeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateKnowledge 更新Knowledge
// @Tags Knowledge
// @Summary 更新Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Knowledge true "更新Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /knowledge/updateKnowledge [put]
func (knowledgeApi *KnowledgeApi) UpdateKnowledge(c *gin.Context) {
	var knowledge basicdata.Knowledge
	_ = c.ShouldBindJSON(&knowledge)
	if err := knowledgeService.UpdateKnowledge(knowledge); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindKnowledge 用id查询Knowledge
// @Tags Knowledge
// @Summary 用id查询Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.Knowledge true "用id查询Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /knowledge/findKnowledge [get]
func (knowledgeApi *KnowledgeApi) FindKnowledge(c *gin.Context) {
	var knowledge basicdata.Knowledge
	_ = c.ShouldBindQuery(&knowledge)
	if reknowledge, err := knowledgeService.GetKnowledge(knowledge.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reknowledge": reknowledge}, c)
	}
}

// GetKnowledgeList 分页获取Knowledge列表
// @Tags Knowledge
// @Summary 分页获取Knowledge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.KnowledgeSearch true "分页获取Knowledge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /knowledge/getKnowledgeList [get]
func (knowledgeApi *KnowledgeApi) GetKnowledgeList(c *gin.Context) {
	var pageInfo basicdataReq.KnowledgeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := knowledgeService.GetKnowledgeInfoList(pageInfo); err != nil {
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
