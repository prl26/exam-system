package teachplan

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/teachplan"
	teachplanReq "github.com/prl26/exam-system/server/model/teachplan/request"
	teachplanResp "github.com/prl26/exam-system/server/model/teachplan/response"
	"time"
)

type TeachAttendanceService struct {
}

// CreateTeachAttendance 创建TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) CreateTeachAttendance(teachAttendance teachplan.TeachAttendance, students []uint) (err error) {
	err = global.GVA_DB.Create(&teachAttendance).Error
	if err != nil {
		return err
	} else {
		if len(students) > 0 {
			attendances := make([]*teachplan.TeachAttendanceRecord, 0, len(students))
			for i := 0; i < len(students); i++ {
				studentId := students[i]
				attendanceId := int(teachAttendance.ID)
				attendanceRecord := &teachplan.TeachAttendanceRecord{
					StudentId:    studentId,
					Status:       0,
					AttendanceId: &attendanceId,
				}
				attendances = append(attendances, attendanceRecord)
			}
			global.GVA_DB.Create(&attendances)
		}
	}
	return err
}

// DeleteTeachAttendance 删除TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) DeleteTeachAttendance(teachAttendance teachplan.TeachAttendance) (err error) {
	err = global.GVA_DB.Delete(&teachAttendance).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&teachplan.TeachAttendanceRecord{}, "attenance_id=?", teachAttendance.ID).Error
	return err
}

// DeleteTeachAttendanceByIds 批量删除TeachAttendance记录
// Author [piexlmax](https://github.com/piexlmax)
func (teachAttendanceService *TeachAttendanceService) DeleteTeachAttendanceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]teachplan.TeachAttendance{}, "id in ?", ids.Ids).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&[]teachplan.TeachAttendanceRecord{}, "attendance_id in ?", ids.Ids).Error
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
func (teachAttendanceService *TeachAttendanceService) GetTeachAttendance(id uint, info request.PageInfo) (list []teachplanResp.AttendanceDetail, total int64, doneTotal int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	err = global.GVA_DB.Model(&teachplan.TeachAttendanceRecord{}).Where("attendance_id=?", id).Count(&total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&teachplan.TeachAttendanceRecord{}).Where("attendance_id=? and attendance > 0", id).Count(&doneTotal).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Raw("select a.*,b.name from tea_attendance_record a left join bas_student b on a.student_id=b.id where a.attendance_id=? limit ? offset ?", id, limit, offset).Find(&list).Error
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
	if info.TeachClassId != 0 {
		db = db.Where("teach_class_id = ?", info.TeachClassId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&teachAttendances).Error
	return teachAttendances, total, err
}

func (t TeachAttendanceService) Attendance(studentId uint, ip string, attendanceId uint, status uint) (int64, error) {
	now := time.Now()
	r := global.GVA_DB.Where("attendance_id=? and student_id=?", attendanceId, studentId).Updates(&teachplan.TeachAttendanceRecord{
		Ip:             ip,
		AttendanceTime: &now,
		Status:         status,
	})
	if r.Error != nil {
		return 0, r.Error
	}

	return r.RowsAffected, nil
}
