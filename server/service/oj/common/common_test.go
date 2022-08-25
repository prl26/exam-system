package common

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"log"
	"testing"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 0:15

 * @Note:

 **/
var server CommonService

func TestMain(m *testing.M) {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_DB = initialize.Gorm()
	m.Run()
}

func TestFindProgrammCase(t *testing.T) {
	programmCase, err := server.FindProgrammCase(1, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(programmCase)
}
