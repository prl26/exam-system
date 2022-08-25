package teachplan

import (
	"exam-system/global"
	"exam-system/model/common/request"
	"exam-system/model/teachplan"
	teachplanReq "exam-system/model/teachplan/request"
)

type TeachAttendanceService struct {
}

// CreateTeachAttendance 创建TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) CreateTeachAttendance(teachAttendance teachplan.TeachAttendance) (err error) {
	err = global.GVA_DB.Create(&teachAttendance).Error
	return err
}

// DeleteTeachAttendance 删除TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) DeleteTeachAttendance(teachAttendance teachplan.TeachAttendance) (err error) {
	err = global.GVA_DB.Delete(&teachAttendance).Error
	return err
}

// DeleteTeachAttendanceByIds 批量删除TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) DeleteTeachAttendanceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]teachplan.TeachAttendance{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeachAttendance 更新TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) UpdateTeachAttendance(teachAttendance teachplan.TeachAttendance) (err error) {
	err = global.GVA_DB.Save(&teachAttendance).Error
	return err
}

// GetTeachAttendance 根据id获取TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) GetTeachAttendance(id uint) (teachAttendance teachplan.TeachAttendance, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teachAttendance).Error
	return
}

// GetTeachAttendanceInfoList 分页获取TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) GetTeachAttendanceInfoList(info teachplanReq.TeachAttendanceSearch) (list []teachplan.TeachAttendance, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&teachplan.TeachAttendance{})
	var teachAttendances []teachplan.TeachAttendance
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TeachId != nil {
		db = db.Where("teach_id = ?", info.TeachId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&teachAttendances).Error
	return teachAttendances, total, err
}
