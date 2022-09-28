package testutils

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"path/filepath"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/27 11:55

 * @Note:

 **/

func InitTest() {
	testPath := filepath.Join(utils.GetCurrentAbPath(), "config_test.yaml")
	global.GVA_VP = setViper(testPath)
	global.GVA_LOG = setLogger()
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = setOrm()
}
