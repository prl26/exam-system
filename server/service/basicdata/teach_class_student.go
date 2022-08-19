package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeachClassStudentService struct {
}

// CreateTeachClassStudent 创建TeachClassStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassStudentService *TeachClassStudentService) CreateTeachClassStudent(teachClassStudent basicdata.TeachClassStudent) (err error) {
	err = global.GVA_DB.Create(&teachClassStudent).Error
	return err
}

// DeleteTeachClassStudent 删除TeachClassStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassStudentService *TeachClassStudentService)DeleteTeachClassStudent(teachClassStudent basicdata.TeachClassStudent) (err error) {
	err = global.GVA_DB.Delete(&teachClassStudent).Error
	return err
}

// DeleteTeachClassStudentByIds 批量删除TeachClassStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassStudentService *TeachClassStudentService)DeleteTeachClassStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.TeachClassStudent{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTeachClassStudent 更新TeachClassStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassStudentService *TeachClassStudentService)UpdateTeachClassStudent(teachClassStudent basicdata.TeachClassStudent) (err error) {
	err = global.GVA_DB.Save(&teachClassStudent).Error
	return err
}

// GetTeachClassStudent 根据id获取TeachClassStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassStudentService *TeachClassStudentService)GetTeachClassStudent(id uint) (teachClassStudent basicdata.TeachClassStudent, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teachClassStudent).Error
	return
}

// GetTeachClassStudentInfoList 分页获取TeachClassStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachClassStudentService *TeachClassStudentService)GetTeachClassStudentInfoList(info basicdataReq.TeachClassStudentSearch) (list []basicdata.TeachClassStudent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&basicdata.TeachClassStudent{})
    var teachClassStudents []basicdata.TeachClassStudent
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&teachClassStudents).Error
	return  teachClassStudents, total, err
}
