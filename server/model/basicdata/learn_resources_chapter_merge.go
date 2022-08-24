// 自动生成模板LearnResourcesChapterMerge
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LearnResourcesChapterMerge 结构体
type LearnResourcesChapterMerge struct {
	global.GVA_MODEL
	LearnResourcesId *int `json:"learnResourcesId" form:"learn_resources_id" gorm:"column:learn_resources_id;comment:课程资源id;"`
	ChapterId        *int `json:"chapterId" form:"chapter_id" gorm:"column:chapter_id;comment:章节id;"`
}

// TableName LearnResourcesChapterMerge 表名
func (LearnResourcesChapterMerge) TableName() string {
	return "bas_learn_resources_chapter_merge"
}
