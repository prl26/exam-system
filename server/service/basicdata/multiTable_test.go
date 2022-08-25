/*
*

	@author: qianyi  2022/8/25 15:53:00
	@note:
*/
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"log"
	"testing"
)

func Test1(t *testing.T) {
	global.GVA_VP = core.Viper() // 初始化Viper
	//global.GVA_LOG = core.Zap()  // 初始化zap日志库
	//zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	var m MultiTableService
	var info = request.TeachClassStudent{
		TeachClassId: 1,
	}
	info.PageInfo.Page = 1
	info.PageInfo.PageSize = 10

	students, i, err := m.GetTeachClassStudentInfo(info)
	if err != nil {
		log.Printf("%q,%d,%q", students, i, err)
	}
}
