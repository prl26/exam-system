package testutils

import (
	"github.com/prl26/exam-system/server/global"
	"go.uber.org/zap"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/27 11:55

 * @Note:

 **/

func InitTest() {
	path := filepath.Join(getCurrentAbPath(), "../../config_test.yaml")
	global.GVA_VP = Viper(path) // 初始化Viper
	global.GVA_LOG = Zap()      // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = orm() // gorm连接数据库
}

func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
