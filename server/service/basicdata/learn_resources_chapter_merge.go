package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type LearnResourcesChapterMergeService struct {
}

// CreateLearnResourcesChapterMerge 创建LearnResourcesChapterMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (learnResourcesChapterMergeService *LearnResourcesChapterMergeService) CreateLearnResourcesChapterMerge(learnResourcesChapterMerge basicdata.LearnResourcesChapterMerge) (err error) {
	err = global.GVA_DB.Create(&learnResourcesChapterMerge).Error
	return err
}

// DeleteLearnResourcesChapterMerge 删除LearnResourcesChapterMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (learnResourcesChapterMergeService *LearnResourcesChapterMergeService) DeleteLearnResourcesChapterMerge(learnResourcesChapterMerge basicdata.LearnResourcesChapterMerge) (err error) {
	err = global.GVA_DB.Delete(&learnResourcesChapterMerge).Error
	return err
}

// DeleteLearnResourcesChapterMergeByIds 批量删除LearnResourcesChapterMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (learnResourcesChapterMergeService *LearnResourcesChapterMergeService) DeleteLearnResourcesChapterMergeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.LearnResourcesChapterMerge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLearnResourcesChapterMerge 更新LearnResourcesChapterMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (learnResourcesChapterMergeService *LearnResourcesChapterMergeService) UpdateLearnResourcesChapterMerge(learnResourcesChapterMerge basicdata.LearnResourcesChapterMerge) (err error) {
	err = global.GVA_DB.Save(&learnResourcesChapterMerge).Error
	return err
}

// GetLearnResourcesChapterMerge 根据id获取LearnResourcesChapterMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (learnResourcesChapterMergeService *LearnResourcesChapterMergeService) GetLearnResourcesChapterMerge(id uint) (learnResourcesChapterMerge basicdata.LearnResourcesChapterMerge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&learnResourcesChapterMerge).Error
	return
}

// GetLearnResourcesChapterMergeInfoList 分页获取LearnResourcesChapterMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (learnResourcesChapterMergeService *LearnResourcesChapterMergeService) GetLearnResourcesChapterMergeInfoList(info basicdataReq.LearnResourcesChapterMergeSearch) (list []basicdata.LearnResourcesChapterMerge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&basicdata.LearnResourcesChapterMerge{})
	var learnResourcesChapterMerges []basicdata.LearnResourcesChapterMerge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.LearnResourcesId != nil {
		db = db.Where("learn_resources_id = ?", info.LearnResourcesId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&learnResourcesChapterMerges).Error
	return learnResourcesChapterMerges, total, err
}
