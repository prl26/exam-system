package basicdata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
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
	err = global.GVA_DB.Updates(&chapter).Error
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
	if info.LessonId != nil {
		db = db.Where("lesson_id = ?", info.LessonId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&chapters).Error
	return chapters, total, err
}

// AccessOrCreateByName 查询lessonId 课程下是否有名称为name的阶段  若有返回阶段ID 若无创建后返回阶段ID
func (chapterService *ChapterService) AccessOrCreateByName(name string, lessonId int) (uint, error) {
	chapter := basicdata.Chapter{}
	err := global.GVA_DB.Where("lesson_id=? and name=?", lessonId, name).Find(&chapter).Error
	if err != nil {
		return 0, err
	}
	if chapter.ID != 0 {
		return chapter.ID, nil
	}
	chapter.LessonId = &lessonId
	chapter.Name = name
	if err = global.GVA_DB.Create(&chapter).Error; err != nil {
		return 0, err
	}
	return chapter.ID, nil
}
