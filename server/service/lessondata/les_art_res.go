package lessondata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/lessondata"
	lessondataReq "github.com/prl26/exam-system/server/model/lessondata/request"
)

type ArticleResourcesService struct {
}

// CreateArticleResources 创建ArticleResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleResourcesService *ArticleResourcesService) CreateArticleResources(articleResources lessondata.ArticleResources) (err error) {
	err = global.GVA_DB.Create(&articleResources).Error
	return err
}

// DeleteArticleResources 删除ArticleResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleResourcesService *ArticleResourcesService) DeleteArticleResources(articleResources lessondata.ArticleResources) (err error) {
	err = global.GVA_DB.Delete(&articleResources).Error
	return err
}

// DeleteArticleResourcesByIds 批量删除ArticleResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleResourcesService *ArticleResourcesService) DeleteArticleResourcesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]lessondata.ArticleResources{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateArticleResources 更新ArticleResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleResourcesService *ArticleResourcesService) UpdateArticleResources(articleResources lessondata.ArticleResources) (err error) {
	err = global.GVA_DB.Updates(&articleResources).Error
	return err
}

// GetArticleResources 根据id获取ArticleResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleResourcesService *ArticleResourcesService) GetArticleResources(id uint) (articleResources lessondata.ArticleResources, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&articleResources).Error
	return
}

// GetArticleResourcesInfoList 分页获取ArticleResources记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleResourcesService *ArticleResourcesService) GetArticleResourcesInfoList(info lessondataReq.ArticleResourcesSearch) (list []lessondata.ArticleResources, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&lessondata.ArticleResources{})
	var articleResourcess []lessondata.ArticleResources
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Author != "" {
		db = db.Where("author LIKE ?", "%"+info.Author+"%")
	}
	if info.IsReference != nil {
		db = db.Where("is_reference = ?", info.IsReference)
	}
	if info.ReferenceUrl != "" {
		db = db.Where("reference_url = ?", info.ReferenceUrl)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articleResourcess).Error
	return articleResourcess, total, err
}
