package teachplan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/teachplan"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    teachplanReq "github.com/flipped-aurora/gin-vue-admin/server/model/teachplan/request"
)

type TeachAttendanceRecordService struct {
}

// CreateTeachAttendanceRecord 创建TeachAttendanceRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceRecordService *TeachAttendanceRecordService) CreateTeachAttendanceRecord(teachAttendanceRecord teachplan.TeachAttendanceRecord) (err error) {
	err = global.GVA_DB.Create(&teachAttendanceRecord).Error
	return err
}

// DeleteTeachAttendanceRecord 删除TeachAttendanceRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceRecordService *TeachAttendanceRecordService)DeleteTeachAttendanceRecord(teachAttendanceRecord teachplan.TeachAttendanceRecord) (err error) {
	err = global.GVA_DB.Delete(&teachAttendanceRecord).Error
	return err
}

// DeleteTeachAttendanceRecordByIds 批量删除TeachAttendanceRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceRecordService *TeachAttendanceRecordService)DeleteTeachAttendanceRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]teachplan.TeachAttendanceRecord{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTeachAttendanceRecord 更新TeachAttendanceRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceRecordService *TeachAttendanceRecordService)UpdateTeachAttendanceRecord(teachAttendanceRecord teachplan.TeachAttendanceRecord) (err error) {
	err = global.GVA_DB.Save(&teachAttendanceRecord).Error
	return err
}

// GetTeachAttendanceRecord 根据id获取TeachAttendanceRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceRecordService *TeachAttendanceRecordService)GetTeachAttendanceRecord(id uint) (teachAttendanceRecord teachplan.TeachAttendanceRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teachAttendanceRecord).Error
	return
}

// GetTeachAttendanceRecordInfoList 分页获取TeachAttendanceRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceRecordService *TeachAttendanceRecordService)GetTeachAttendanceRecordInfoList(info teachplanReq.TeachAttendanceRecordSearch) (list []teachplan.TeachAttendanceRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&teachplan.TeachAttendanceRecord{})
    var teachAttendanceRecords []teachplan.TeachAttendanceRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StudentId != nil {
        db = db.Where("student_id = ?",info.StudentId)
    }
    if info.AttendanceId != nil {
        db = db.Where("attendance_id = ?",info.AttendanceId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&teachAttendanceRecords).Error
	return  teachAttendanceRecords, total, err
}
