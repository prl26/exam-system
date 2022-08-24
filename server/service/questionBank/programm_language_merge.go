package questionBank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	questionBankReq "github.com/flipped-aurora/gin-vue-admin/server/model/questionBank/request"
)

type ProgrammLanguageMergeService struct {
}

// CreateProgrammLanguageMerge 创建ProgrammLanguageMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (programmLanguageMergeService *ProgrammLanguageMergeService) CreateProgrammLanguageMerge(programmLanguageMerge questionBank.ProgrammLanguageMerge) (err error) {
	err = global.GVA_DB.Create(&programmLanguageMerge).Error
	return err
}

// DeleteProgrammLanguageMerge 删除ProgrammLanguageMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (programmLanguageMergeService *ProgrammLanguageMergeService) DeleteProgrammLanguageMerge(programmLanguageMerge questionBank.ProgrammLanguageMerge) (err error) {
	err = global.GVA_DB.Delete(&programmLanguageMerge).Error
	return err
}

// DeleteProgrammLanguageMergeByIds 批量删除ProgrammLanguageMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (programmLanguageMergeService *ProgrammLanguageMergeService) DeleteProgrammLanguageMergeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]questionBank.ProgrammLanguageMerge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProgrammLanguageMerge 更新ProgrammLanguageMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (programmLanguageMergeService *ProgrammLanguageMergeService) UpdateProgrammLanguageMerge(programmLanguageMerge questionBank.ProgrammLanguageMerge) (err error) {
	err = global.GVA_DB.Save(&programmLanguageMerge).Error
	return err
}

// GetProgrammLanguageMerge 根据id获取ProgrammLanguageMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (programmLanguageMergeService *ProgrammLanguageMergeService) GetProgrammLanguageMerge(id uint) (programmLanguageMerge questionBank.ProgrammLanguageMerge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&programmLanguageMerge).Error
	return
}

// GetProgrammLanguageMergeInfoList 分页获取ProgrammLanguageMerge记录
// Author [piexlmax](https://github.com/piexlmax)
func (programmLanguageMergeService *ProgrammLanguageMergeService) GetProgrammLanguageMergeInfoList(info questionBankReq.ProgrammLanguageMergeSearch) (list []questionBank.ProgrammLanguageMerge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&questionBank.ProgrammLanguageMerge{})
	var programmLanguageMerges []questionBank.ProgrammLanguageMerge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.LanguageId != nil {
		db = db.Where("language_id = ?", info.LanguageId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&programmLanguageMerges).Error
	return programmLanguageMerges, total, err
}
