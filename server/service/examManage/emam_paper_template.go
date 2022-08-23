package examManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/examManage"
	examManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/examManage/request"
)

type PaperTemplateService struct {
}

// CreatePaperTemplate 创建PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService) CreatePaperTemplate(Papertemplate examManage.PaperTemplate) (err error) {
	err = global.GVA_DB.Create(&Papertemplate).Error
	return err
}

// DeletePaperTemplate 删除PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService)DeletePaperTemplate(Papertemplate examManage.PaperTemplate) (err error) {
	err = global.GVA_DB.Delete(&Papertemplate).Error
	return err
}

// DeletePaperTemplateByIds 批量删除PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService)DeletePaperTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.PaperTemplate{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePaperTemplate 更新PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService)UpdatePaperTemplate(Papertemplate examManage.PaperTemplate) (err error) {
	err = global.GVA_DB.Save(&Papertemplate).Error
	return err
}

// GetPaperTemplate 根据id获取PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService)GetPaperTemplate(id uint) (Papertemplate examManage.PaperTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&Papertemplate).Error
	return
}

// GetPaperTemplateInfoList 分页获取PaperTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (PapertemplateService *PaperTemplateService)GetPaperTemplateInfoList(info examManageReq.PaperTemplateSearch) (list []examManage.PaperTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&examManage.PaperTemplate{})
    var Papertemplates []examManage.PaperTemplate
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.CourseId != 0 {
        db = db.Where("course_id = ?",info.CourseId)
    }
    if info.UserId != 0 {
        db = db.Where("user_id = ?",info.UserId)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&Papertemplates).Error
	return  Papertemplates, total, err
}
