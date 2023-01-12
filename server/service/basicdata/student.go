package basicdata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/utils"
	"strconv"
)

type StudentService struct {
}

// CreateStudent 创建Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) CreateStudent(student basicdata.Student) (err error) {
	err = global.GVA_DB.Create(&student).Error
	return err
}

// DeleteStudent 删除Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) DeleteStudent(student basicdata.Student) (err error) {
	err = global.GVA_DB.Delete(&student).Error
	return err
}

// DeleteStudentByIds 批量删除Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) DeleteStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]basicdata.Student{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateStudent 更新Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) UpdateStudent(student basicdata.Student) (err error) {
	err = global.GVA_DB.Updates(&student).Error
	return err
}

// GetStudent 根据id获取Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) GetStudent(id uint) (student basicdata.Student, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

// GetStudentInfoList 分页获取Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) GetStudentInfoList(info basicdataReq.StudentSearch) (list []basicdata.Student, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&basicdata.Student{})
	var students []basicdata.Student
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.IdCard != "" {
		db = db.Where("id_card LIKE ?", "%d"+info.IdCard+"%d")
	}
	if info.CollegeName != "" {
		db = db.Where("college_name = ?", info.CollegeName)
	}
	if info.ProfessionalName != "" {
		db = db.Where("professional_name = ?", info.ProfessionalName)
	}
	if info.ClassId != 0 {
		db = db.Where("class_id = ?", info.ClassId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&students).Error
	return students, total, err
}

func (studentService *StudentService) CreateStudents(students []*basicdata.Student) error {
	err := global.GVA_DB.Create(&students).Error
	return err
}

func (studentService *StudentService) QueryStudentById(id uint) basicdata.Student {
	var student basicdata.Student
	global.GVA_DB.Where("id = ?", id).First(&student)
	return student
}

// ResetStudentsPassword 重置密码为学号
func (studentService *StudentService) ResetStudentsPassword(idReq request.IdReq) error {
	var student basicdata.Student
	id := idReq.Id
	resetpwd := utils.BcryptHash(strconv.Itoa(id))
	return global.GVA_DB.Model(&student).Where("id = ?", id).Update("password", resetpwd).Error
}
