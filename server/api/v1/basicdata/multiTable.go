/*
*

	@author: qianyi  2022/8/24 19:00:00
	@note:
*/
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MultiTableApi struct {
}

var multiTableService = service.ServiceGroupApp.BasicdataApiGroup.MultiTableService

// InitTeachClassStudent 关联 学生与教学班
// @Tags TeachClassStudent
// @Summary 初始化关联 学生与教学班
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StuTeachClass true "创建TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/initTeachClassStudent [post]
func (multiTableServiceApi *MultiTableApi) InitTeachClassStudent(c *gin.Context) {
	var stuClassReq basicdataReq.StuTeachClass
	_ = c.ShouldBindJSON(&stuClassReq)
	err := multiTableService.UpdateTeachClassStudents(stuClassReq)
	if err != nil {
		global.GVA_LOG.Error("更新学生教学班关联表失败", zap.Error(err))
		response.FailWithMessage("更新学生教学班关联表失败", c)
	} else {
		response.OkWithMessage("更新学生教学班关联表成功", c)
	}

}

// GetTeachClassStudentList 获取一个教学班的学生
// @Tags TeachClassStudent
// @Summary 获取一个 教学班的学生
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StuTeachClass true "获取一个教学班的学生"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/getTeachClassStudentList [get]
//func (multiTableServiceApi *MultiTableApi) GetTeachClassStudentList(c *gin.Context) {
//	var pageInfo basicdataReq.TeachClassIdSearch
//	_ = c.ShouldBindQuery(&pageInfo)
//	if list, total, err := multiTableService.GetStudentInfo(); err != nil {
//		global.GVA_LOG.Error("获取失败!", zap.Error(err))
//		response.FailWithMessage("获取失败", c)
//	} else {
//		response.OkWithDetailed(response.PageResult{
//			List:     list,
//			Total:    total,
//			Page:     pageInfo.Page,
//			PageSize: pageInfo.PageSize,
//		}, "获取成功", c)
//	}
//}
