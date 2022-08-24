/*
*

	@author: qianyi  2022/8/24 19:00:00
	@note:
*/
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type MultiTableServiceApi struct {
}

var multiTableService = service.ServiceGroupApp.BasicdataApiGroup.MultiTableService

// 更新学生教学班关联表
func (multiTableServiceApi *MultiTableServiceApi) UpdateTeachClassStudent(c *gin.Context) {

	// 接收 教学班id 和要与之关联的学生id的 数组

}
