package testutils

import (
	"github.com/prl26/exam-system/server/global"
	"gorm.io/gorm"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/27 12:02

 * @Note:

 **/

func orm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}
