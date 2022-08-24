/*
*

	@author: qianyi  2022/8/24 19:00:00
	@note:
*/
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type MultiTableServiceApi struct {
}

//var student basicdata.Student
//_ = c.ShouldBindJSON(&student)
//if err := studentService.CreateStudent(student); err != nil {
//global.GVA_LOG.Error("创建失败!", zap.Error(err))
//response.FailWithMessage("创建失败", c)
//} else {
//response.OkWithMessage("创建成功", c)
//}

var multiTableService = service.ServiceGroupApp.BasicdataApiGroup.MultiTableService

// 更新学生教学班关联表
func (multiTableServiceApi *MultiTableServiceApi) UpdateTeachClassStudent(c *gin.Context) {

	// 接收 教学班id 和要与之关联的学生id的 数组
	var stuClassReq request.StuTeachClass

	err := c.ShouldBindJSON(&stuClassReq)
	if err != nil {
		c.Writer.Write([]byte("绑定参数出错"))
		return
	}

	err := multiTableService.UpdateTeachClassStudents(stuClassReq)
	if err != nil {
		return
	}

}
