package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	basicdataReq "github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
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
	err = global.GVA_DB.Save(&student).Error
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
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.Id_card != "" {
		db = db.Where("id_card = ?", info.Id_card)
	}
	if info.College_id != nil {
		db = db.Where("college_id = ?", info.College_id)
	}
	if info.Professional_id != nil {
		db = db.Where("professional_id = ?", info.Professional_id)
	}
	if info.Class_id != nil {
		db = db.Where("class_id = ?", info.Class_id)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&students).Error
	return students, total, err
}
