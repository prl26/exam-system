package system

import (
	"github.com/prl26/exam-system/server/config"
	"github.com/prl26/exam-system/server/global"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}
