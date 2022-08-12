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

type TeachClassStudentApi struct {
}

var teachClassStudentService = service.ServiceGroupApp.BasicdataApiGroup.TeachClassStudentService


// CreateTeachClassStudent 创建TeachClassStudent
// @Tags TeachClassStudent
// @Summary 创建TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.TeachClassStudent true "创建TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/createTeachClassStudent [post]
func (teachClassStudentApi *TeachClassStudentApi) CreateTeachClassStudent(c *gin.Context) {
	var teachClassStudent basicdata.TeachClassStudent
	_ = c.ShouldBindJSON(&teachClassStudent)
    verify := utils.Rules{
        "Student_id":{utils.NotEmpty()},
        "Teach_class_id":{utils.NotEmpty()},
    }
	if err := utils.Verify(teachClassStudent, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := teachClassStudentService.CreateTeachClassStudent(teachClassStudent); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeachClassStudent 删除TeachClassStudent
// @Tags TeachClassStudent
// @Summary 删除TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.TeachClassStudent true "删除TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachClassStudent/deleteTeachClassStudent [delete]
func (teachClassStudentApi *TeachClassStudentApi) DeleteTeachClassStudent(c *gin.Context) {
	var teachClassStudent basicdata.TeachClassStudent
	_ = c.ShouldBindJSON(&teachClassStudent)
	if err := teachClassStudentService.DeleteTeachClassStudent(teachClassStudent); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeachClassStudentByIds 批量删除TeachClassStudent
// @Tags TeachClassStudent
// @Summary 批量删除TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teachClassStudent/deleteTeachClassStudentByIds [delete]
func (teachClassStudentApi *TeachClassStudentApi) DeleteTeachClassStudentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := teachClassStudentService.DeleteTeachClassStudentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeachClassStudent 更新TeachClassStudent
// @Tags TeachClassStudent
// @Summary 更新TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.TeachClassStudent true "更新TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teachClassStudent/updateTeachClassStudent [put]
func (teachClassStudentApi *TeachClassStudentApi) UpdateTeachClassStudent(c *gin.Context) {
	var teachClassStudent basicdata.TeachClassStudent
	_ = c.ShouldBindJSON(&teachClassStudent)
      verify := utils.Rules{
          "Student_id":{utils.NotEmpty()},
          "Teach_class_id":{utils.NotEmpty()},
      }
    if err := utils.Verify(teachClassStudent, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := teachClassStudentService.UpdateTeachClassStudent(teachClassStudent); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeachClassStudent 用id查询TeachClassStudent
// @Tags TeachClassStudent
// @Summary 用id查询TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.TeachClassStudent true "用id查询TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teachClassStudent/findTeachClassStudent [get]
func (teachClassStudentApi *TeachClassStudentApi) FindTeachClassStudent(c *gin.Context) {
	var teachClassStudent basicdata.TeachClassStudent
	_ = c.ShouldBindQuery(&teachClassStudent)
	if reteachClassStudent, err := teachClassStudentService.GetTeachClassStudent(teachClassStudent.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteachClassStudent": reteachClassStudent}, c)
	}
}

// GetTeachClassStudentList 分页获取TeachClassStudent列表
// @Tags TeachClassStudent
// @Summary 分页获取TeachClassStudent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.TeachClassStudentSearch true "分页获取TeachClassStudent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/getTeachClassStudentList [get]
func (teachClassStudentApi *TeachClassStudentApi) GetTeachClassStudentList(c *gin.Context) {
	var pageInfo basicdataReq.TeachClassStudentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := teachClassStudentService.GetTeachClassStudentInfoList(pageInfo); err != nil {
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
