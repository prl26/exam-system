package basicdata

import (
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type CollegeApi struct {
}

var collegeService = service.ServiceGroupApp.BasicdataApiGroup.CollegeService

// CreateCollege 创建College
// @Tags College
// @Summary 创建College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.College true "创建College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /college/createCollege [post]
func (collegeApi *CollegeApi) CreateCollege(c *gin.Context) {
	var college basicdata.College
	_ = c.ShouldBindJSON(&college)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(college, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := collegeService.CreateCollege(college); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCollege 删除College
// @Tags College
// @Summary 删除College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.College true "删除College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /college/deleteCollege [delete]
func (collegeApi *CollegeApi) DeleteCollege(c *gin.Context) {
	var college basicdata.College
	_ = c.ShouldBindJSON(&college)
	if err := collegeService.DeleteCollege(college); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCollegeByIds 批量删除College
// @Tags College
// @Summary 批量删除College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /college/deleteCollegeByIds [delete]
func (collegeApi *CollegeApi) DeleteCollegeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := collegeService.DeleteCollegeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCollege 更新College
// @Tags College
// @Summary 更新College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.College true "更新College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /college/updateCollege [put]
func (collegeApi *CollegeApi) UpdateCollege(c *gin.Context) {
	var college basicdata.College
	_ = c.ShouldBindJSON(&college)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(college, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := collegeService.UpdateCollege(college); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCollege 用id查询College
// @Tags College
// @Summary 用id查询College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.College true "用id查询College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /college/findCollege [get]
func (collegeApi *CollegeApi) FindCollege(c *gin.Context) {
	var college basicdata.College
	_ = c.ShouldBindQuery(&college)
	if recollege, err := collegeService.GetCollege(college.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recollege": recollege}, c)
	}
}

// GetCollegeList 分页获取College列表
// @Tags College
// @Summary 分页获取College列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.CollegeSearch true "分页获取College列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /college/getCollegeList [get]
func (collegeApi *CollegeApi) GetCollegeList(c *gin.Context) {
	var pageInfo basicdataReq.CollegeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := collegeService.GetCollegeInfoList(pageInfo); err != nil {
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
