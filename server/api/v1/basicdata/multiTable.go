/*
*

	@author: qianyi  2022/8/24 19:00:00
	@note:
*/
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MultiTableApi struct {
}

var multiTableService = service.ServiceGroupApp.BasicdataApiGroup.MultiTableService

// CreateTeachClass 创建TeachClass
// @Tags TeachClass
// @Summary 创建TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.TeachClass true "创建TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClass/createTeachClass [post]

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
	var stuClassReq request.StuTeachClass

	_ = c.ShouldBindJSON(&stuClassReq)
	err := multiTableService.UpdateTeachClassStudents(stuClassReq)
	if err != nil {
		global.GVA_LOG.Error("更新学生教学班关联表失败", zap.Error(err))
		response.FailWithMessage("更新学生教学班关联表失败", c)
	} else {
		response.OkWithMessage("更新学生教学班关联表成功", c)
	}

}
