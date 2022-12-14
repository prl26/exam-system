package teachplan

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/teachplan"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
)

type TeachAttendanceService struct {
}

// CreateTeachAttendance 创建TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) CreateTeachAttendance(teachAttendance teachplan.TeachAttendance, students []basicdata.Student) (err error) {
	err = global.GVA_DB.Create(&teachAttendance).Error
	if err != nil {
		return err
	} else {

		if err != nil {
			return err
		} else {
			if len(students) > 0 {
				for i := 0; i < len(students); i++ {
					studentId := int(students[i].ID)
					attendanceId := int(teachAttendance.ID)
					attendanceRecord := teachplan.TeachAttendanceRecord{
						GVA_MODEL:    global.GVA_MODEL{},
						StudentId:    &studentId,
						Longitute:    nil,
						Latitude:     nil,
						Status:       0,
						AttendanceId: &attendanceId,
					}
					global.GVA_DB.Create(&attendanceRecord)
				}
			}
		}
	}
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
	err = global.GVA_DB.Updates(&teachAttendance).Error
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
	if info.TeachClassId != nil {
		db = db.Where("teach_id = ?", info.TeachClassId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&teachAttendances).Error
	return teachAttendances, total, err
}
