package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	examManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/examManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaperTemplateApi struct {
}

var PapertemplateService = service.ServiceGroupApp.ExammanageServiceGroup.PaperTemplateService


// CreatePaperTemplate 创建PaperTemplate
// @Tags PaperTemplate
// @Summary 创建PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.PaperTemplate true "创建PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Papertemplate/createPaperTemplate [post]
func (PapertemplateApi *PaperTemplateApi) CreatePaperTemplate(c *gin.Context) {
	uid := int(utils.GetUserID(c))
	var Papertemplate examManage.PaperTemplate
	_ = c.ShouldBindJSON(&Papertemplate)
	Papertemplate.UserId = &uid
	if err := PapertemplateService.CreatePaperTemplate(Papertemplate); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		//for i:=0; i<len(Papertemplate.PaperTemplateItems);i++{
		//	b := &Papertemplate.PaperTemplateItems[i]
		//	valueofb := reflect.ValueOf(b).Elem()
		//	PaperTemplateItem := examManage.PaperTemplateItem{
		//		GVA_MODEL:   global.GVA_MODEL{},
		//		Chapter:     valueofb.FieldByName("Chapter").String(),
		//		ProblemType: int(reflect.ValueOf(b.ProblemType).Int()),
		//		Difficulty:  nil,
		//		Num:         nil,
		//		Score:       nil,
		//		TemplateId:  nil,
		//	}
		//
		//}
		for i:=0; i<len(Papertemplate.PaperTemplateItems);i++{
			b := Papertemplate.PaperTemplateItems[i]
			paperTemplateItem := examManage.PaperTemplateItem{
				GVA_MODEL:   global.GVA_MODEL{},
				Chapter:     b.Chapter,
				ProblemType: b.ProblemType,
				Difficulty:  b.Difficulty,
				Num:         b.Num,
				Score:       b.Score,
				TemplateId:  Papertemplate.UserId,
			}
			if err := paperTemplateItemService.CreatePaperTemplateItem(paperTemplateItem); err != nil {
				global.GVA_LOG.Error("创建失败!", zap.Error(err))
				response.FailWithMessage("创建失败", c)
			} else {
				response.OkWithMessage("创建成功", c)
			}
		}
	}
}

// DeletePaperTemplate 删除PaperTemplate
// @Tags PaperTemplate
// @Summary 删除PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.PaperTemplate true "删除PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Papertemplate/deletePaperTemplate [delete]
func (PapertemplateApi *PaperTemplateApi) DeletePaperTemplate(c *gin.Context) {
	var Papertemplate examManage.PaperTemplate
	_ = c.ShouldBindJSON(&Papertemplate)
	if err := PapertemplateService.DeletePaperTemplate(Papertemplate); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePaperTemplateByIds 批量删除PaperTemplate
// @Tags PaperTemplate
// @Summary 批量删除PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Papertemplate/deletePaperTemplateByIds [delete]
func (PapertemplateApi *PaperTemplateApi) DeletePaperTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := PapertemplateService.DeletePaperTemplateByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePaperTemplate 更新PaperTemplate
// @Tags PaperTemplate
// @Summary 更新PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body examManage.PaperTemplate true "更新PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Papertemplate/updatePaperTemplate [put]
func (PapertemplateApi *PaperTemplateApi) UpdatePaperTemplate(c *gin.Context) {
	var Papertemplate examManage.PaperTemplate
	_ = c.ShouldBindJSON(&Papertemplate)
	if err := PapertemplateService.UpdatePaperTemplate(Papertemplate); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPaperTemplate 用id查询PaperTemplate
// @Tags PaperTemplate
// @Summary 用id查询PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManage.PaperTemplate true "用id查询PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Papertemplate/findPaperTemplate [get]
func (PapertemplateApi *PaperTemplateApi) FindPaperTemplate(c *gin.Context) {
	var Papertemplate examManage.PaperTemplate
	_ = c.ShouldBindQuery(&Papertemplate)
	if rePapertemplate, err := PapertemplateService.GetPaperTemplate(Papertemplate.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePapertemplate": rePapertemplate}, c)
	}
}

// GetPaperTemplateList 分页获取PaperTemplate列表
// @Tags PaperTemplate
// @Summary 分页获取PaperTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query examManageReq.PaperTemplateSearch true "分页获取PaperTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Papertemplate/getPaperTemplateList [get]
func (PapertemplateApi *PaperTemplateApi) GetPaperTemplateList(c *gin.Context) {
	var pageInfo examManageReq.PaperTemplateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := PapertemplateService.GetPaperTemplateInfoList(pageInfo); err != nil {
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
