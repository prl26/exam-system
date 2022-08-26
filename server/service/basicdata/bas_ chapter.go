package basicdata

import (
	"exam-system/global"
	"exam-system/model/basicdata"
	basicdataReq "exam-system/model/basicdata/request"
	"exam-system/model/common/request"
)

type ChapterService struct {
}

// CreateChapter 创建Chapter记录
// Author [piexlmax](https://github.com/piexlmax)
func (chapterService *ChapterService) CreateChapter(chapter basicdata.Chapter) (err error) {
	err = global.GVA_DB.Create(&chapter).Error
	return err
}

// DeleteChapter 删除Chapter记录
// Author [piexlmax](https://github.com/piexlmax)
func (chapterService *ChapterService) DeleteChapter(chapter basicdata.Chapter) (err error) {
	err = global.GVA_DB.Delete(&chapter).Error
	return err
}

// DeleteChapterByIds 批量删除Chapter记录
// Author [piexlmax](https://github.com/piexlmax)
func (chapterService *ChapterService) DeleteChapterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.Chapter{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateChapter 更新Chapter记录
// Author [piexlmax](https://github.com/piexlmax)
func (chapterService *ChapterService) UpdateChapter(chapter basicdata.Chapter) (err error) {
	err = global.GVA_DB.Save(&chapter).Error
	return err
}

// GetChapter 根据id获取Chapter记录
// Author [piexlmax](https://github.com/piexlmax)
func (chapterService *ChapterService) GetChapter(id uint) (chapter basicdata.Chapter, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chapter).Error
	return
}

// GetChapterInfoList 分页获取Chapter记录
// Author [piexlmax](https://github.com/piexlmax)
func (chapterService *ChapterService) GetChapterInfoList(info basicdataReq.ChapterSearch) (list []basicdata.Chapter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&basicdata.Chapter{})
	var chapters []basicdata.Chapter
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.LessonId != "" {
		db = db.Where("lesson_id = ?", info.LessonId)
	}
	if info.Order != nil {
		db = db.Where("order = ?", info.Order)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&chapters).Error
	return chapters, total, err
}
