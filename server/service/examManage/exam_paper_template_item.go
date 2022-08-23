package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    examManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/examManage/request"
)

type PaperTemplateItemService struct {
}

// CreatePaperTemplateItem 创建PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService) CreatePaperTemplateItem(paperTemplateItem examManage.PaperTemplateItem) (err error) {
	err = global.GVA_DB.Create(&paperTemplateItem).Error
	return err
}

// DeletePaperTemplateItem 删除PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService)DeletePaperTemplateItem(paperTemplateItem examManage.PaperTemplateItem) (err error) {
	err = global.GVA_DB.Delete(&paperTemplateItem).Error
	return err
}

// DeletePaperTemplateItemByIds 批量删除PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService)DeletePaperTemplateItemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.PaperTemplateItem{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePaperTemplateItem 更新PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService)UpdatePaperTemplateItem(paperTemplateItem examManage.PaperTemplateItem) (err error) {
	err = global.GVA_DB.Save(&paperTemplateItem).Error
	return err
}

// GetPaperTemplateItem 根据id获取PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService)GetPaperTemplateItem(id uint) (paperTemplateItem examManage.PaperTemplateItem, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&paperTemplateItem).Error
	return
}

// GetPaperTemplateItemInfoList 分页获取PaperTemplateItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (paperTemplateItemService *PaperTemplateItemService)GetPaperTemplateItemInfoList(info examManageReq.PaperTemplateItemSearch) (list []examManage.PaperTemplateItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&examManage.PaperTemplateItem{})
    var paperTemplateItems []examManage.PaperTemplateItem
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Chapter != "" {
        db = db.Where("chapter = ?",info.Chapter)
    }
    if info.ProblemType != nil {
        db = db.Where("problem_type = ?",info.ProblemType)
    }
    if info.Difficulty != nil {
        db = db.Where("difficulty = ?",info.Difficulty)
    }
    if info.Num != nil {
        db = db.Where("num = ?",info.Num)
    }
    if info.Score != nil {
        db = db.Where("score = ?",info.Score)
    }
    if info.TemplateId != nil {
        db = db.Where("template_id = ?",info.TemplateId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&paperTemplateItems).Error
	return  paperTemplateItems, total, err
}
