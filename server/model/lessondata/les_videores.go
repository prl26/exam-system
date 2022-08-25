// 自动生成模板VideoResources
package lessondata

import (
	"exam-system/global"
)

// VideoResources 结构体
type VideoResources struct {
	global.GVA_MODEL
	Url string `json:"url" form:"url" gorm:"column:url;comment:视频资源路径;size:255;"`
}

// TableName VideoResources 表名
func (VideoResources) TableName() string {
	return "les_video_resources"
}
