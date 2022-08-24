package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProfessionalService struct {
}

// CreateProfessional 创建Professional记录
// Author [piexlmax](https://github.com/piexlmax)
func (professionalService *ProfessionalService) CreateProfessional(professional basicdata.Professional) (err error) {
	err = global.GVA_DB.Create(&professional).Error
	return err
}

// DeleteProfessional 删除Professional记录
// Author [piexlmax](https://github.com/piexlmax)
func (professionalService *ProfessionalService) DeleteProfessional(professional basicdata.Professional) (err error) {
	err = global.GVA_DB.Delete(&professional).Error
	return err
}

// DeleteProfessionalByIds 批量删除Professional记录
// Author [piexlmax](https://github.com/piexlmax)
func (professionalService *ProfessionalService) DeleteProfessionalByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.Professional{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProfessional 更新Professional记录
// Author [piexlmax](https://github.com/piexlmax)
func (professionalService *ProfessionalService) UpdateProfessional(professional basicdata.Professional) (err error) {
	err = global.GVA_DB.Save(&professional).Error
	return err
}

// GetProfessional 根据id获取Professional记录
// Author [piexlmax](https://github.com/piexlmax)
func (professionalService *ProfessionalService) GetProfessional(id uint) (professional basicdata.Professional, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&professional).Error
	return
}

// GetProfessionalInfoList 分页获取Professional记录
// Author [piexlmax](https://github.com/piexlmax)
func (professionalService *ProfessionalService) GetProfessionalInfoList(info basicdataReq.ProfessionalSearch) (list []basicdata.Professional, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&basicdata.Professional{})
	var professionals []basicdata.Professional
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.CollegeId != nil {
		db = db.Where("college_id = ?", info.CollegeId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&professionals).Error
	return professionals, total, err
}
