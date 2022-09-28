package utils

import (
	"path"
	"runtime"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/27 16:17

 * @Note:

 **/

// GetCurrentAbPath 获取项目路径
func GetCurrentAbPath() string {
	return path.Join(getCurrentAbPathByCaller(), "../")
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
