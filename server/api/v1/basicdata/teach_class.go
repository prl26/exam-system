package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TeachClassApi struct {
}

var teachClassService = service.ServiceGroupApp.BasicdataApiGroup.TeachClassService

// CreateTeachClass 创建TeachClass
// @Tags TeachClass
// @Summary 创建TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.TeachClass true "创建TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClass/createTeachClass [post]
func (teachClassApi *TeachClassApi) CreateTeachClass(c *gin.Context) {
	var teachClass basicdata.TeachClass
	_ = c.ShouldBindJSON(&teachClass)
	verify := utils.Rules{
		"CourseId":              {utils.NotEmpty()},
		"TermId":                {utils.NotEmpty()},
		"Belong_class_id":       {utils.NotEmpty()},
		"Name":                  {utils.NotEmpty()},
		"Attendance_proportion": {utils.NotEmpty()},
	}
	if err := utils.Verify(teachClass, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teachClassService.CreateTeachClass(teachClass); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeachClass 删除TeachClass
// @Tags TeachClass
// @Summary 删除TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.TeachClass true "删除TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachClass/deleteTeachClass [delete]
func (teachClassApi *TeachClassApi) DeleteTeachClass(c *gin.Context) {
	var teachClass basicdata.TeachClass
	_ = c.ShouldBindJSON(&teachClass)
	if err := teachClassService.DeleteTeachClass(teachClass); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeachClassByIds 批量删除TeachClass
// @Tags TeachClass
// @Summary 批量删除TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teachClass/deleteTeachClassByIds [delete]
func (teachClassApi *TeachClassApi) DeleteTeachClassByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := teachClassService.DeleteTeachClassByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeachClass 更新TeachClass
// @Tags TeachClass
// @Summary 更新TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.TeachClass true "更新TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teachClass/updateTeachClass [put]
func (teachClassApi *TeachClassApi) UpdateTeachClass(c *gin.Context) {
	var teachClass basicdata.TeachClass
	_ = c.ShouldBindJSON(&teachClass)
	verify := utils.Rules{
		"CourseId":              {utils.NotEmpty()},
		"TermId":                {utils.NotEmpty()},
		"Belong_class_id":       {utils.NotEmpty()},
		"Name":                  {utils.NotEmpty()},
		"Attendance_proportion": {utils.NotEmpty()},
	}
	if err := utils.Verify(teachClass, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teachClassService.UpdateTeachClass(teachClass); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeachClass 用id查询TeachClass
// @Tags TeachClass
// @Summary 用id查询TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.TeachClass true "用id查询TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teachClass/findTeachClass [get]
func (teachClassApi *TeachClassApi) FindTeachClass(c *gin.Context) {
	var teachClass basicdata.TeachClass
	_ = c.ShouldBindQuery(&teachClass)
	if reteachClass, err := teachClassService.GetTeachClass(teachClass.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteachClass": reteachClass}, c)
	}
}

// GetTeachClassList 分页获取TeachClass列表
// @Tags TeachClass
// @Summary 分页获取TeachClass列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.TeachClassSearch true "分页获取TeachClass列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClass/getTeachClassList [get]
func (teachClassApi *TeachClassApi) GetTeachClassList(c *gin.Context) {
	var pageInfo basicdataReq.TeachClassSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := teachClassService.GetTeachClassInfoList(pageInfo); err != nil {
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
