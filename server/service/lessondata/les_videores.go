package lessondata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/lessondata"
	lessondataReq "github.com/prl26/exam-system/server/model/lessondata/request"
)

type VideoResourcesService struct {
}

// CreateVideoResources 创建VideoResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoResourcesService *VideoResourcesService) CreateVideoResources(videoResources lessondata.VideoResources) (err error) {
	err = global.GVA_DB.Create(&videoResources).Error
	return err
}

// DeleteVideoResources 删除VideoResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoResourcesService *VideoResourcesService) DeleteVideoResources(videoResources lessondata.VideoResources) (err error) {
	err = global.GVA_DB.Delete(&videoResources).Error
	return err
}

// DeleteVideoResourcesByIds 批量删除VideoResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoResourcesService *VideoResourcesService) DeleteVideoResourcesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.VideoResources{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateVideoResources 更新VideoResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoResourcesService *VideoResourcesService) UpdateVideoResources(videoResources lessondata.VideoResources) (err error) {
	err = global.GVA_DB.Updates(&videoResources).Error
	return err
}

// GetVideoResources 根据id获取VideoResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoResourcesService *VideoResourcesService) GetVideoResources(id uint) (videoResources lessondata.VideoResources, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&videoResources).Error
	return
}

// GetVideoResourcesInfoList 分页获取VideoResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoResourcesService *VideoResourcesService) GetVideoResourcesInfoList(info lessondataReq.VideoResourcesSearch) (list []lessondata.VideoResources, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lessondata.VideoResources{})
	var videoResourcess []lessondata.VideoResources
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Url != "" {
		db = db.Where("url = ?", info.Url)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&videoResourcess).Error
	return videoResourcess, total, err
}
