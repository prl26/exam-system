/*
*

	@author: qianyi  2022/8/24 18:56:00
	@note:
*/
package basicdata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/basicdata/request"
)

type MultiTableService struct {
}

// InitTeachClassStudents 向教学班中 加入学生的关联（单一加入）
func (multiTableService *MultiTableService) InitTeachClassStudents(TeachClassId uint, students []*basicdata.Student) error {

	var teachClass basicdata.TeachClass

	teachClass.ID = TeachClassId

	err := global.GVA_DB.Model(&teachClass).Association("Student").Append(students)

	return err
}

func (multiTableService *MultiTableService) AssociationStudents(students []*basicdata.Student, teachClassID uint) error {

	var teachClass basicdata.TeachClass
	teachClass.ID = teachClassID

	err := global.GVA_DB.Model(&teachClass).Association("Student").Append(students)

	return err
}

func (multiTableService *MultiTableService) AssociationStudent(student *basicdata.Student, teachClassID uint) error {

	var teachClass basicdata.TeachClass
	teachClass.ID = teachClassID

	err := global.GVA_DB.Model(&teachClass).Association("Student").Append(student)

	return err
}

// DeleteTeachClassStudents 在教学班中 移除学生
func (multiTableService *MultiTableService) DeleteTeachClassStudents(info request.StuTeachClass) error {

	var teachClass basicdata.TeachClass
	teachClass.ID = info.TeachClassId

	n := len(info.StudentIds)

	students := make([]basicdata.Student, n)
	for i := 0; i < n; i++ {
		students[i].ID = info.StudentIds[i]
	}

	err := global.GVA_DB.Model(&teachClass).Association("Student").Delete(students)

	return err
}

// GetTeachClassStudentInfo 获取教学班的学生
func (multiTableService *MultiTableService) GetTeachClassStudentInfo(info request.TeachClassStudent) (list []basicdata.Student, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//var limit, offset int
	//
	//if info.PageSize <= 0 {
	//	limit = 10
	//} else {
	//	limit = info.PageSize
	//}
	//
	//if info.Page <= 0 {
	//	offset = 0
	//} else {
	//	offset = info.PageSize * (info.Page - 1)
	//}

	var teachClass basicdata.TeachClass
	teachClass.ID = info.TeachClassId

	var students []basicdata.Student
	var sLength []basicdata.Student

	db := global.GVA_DB

	_ = db.Model(&teachClass).Association("Student").Find(&sLength)
	total = int64(len(sLength))
	err = db.Limit(limit).Offset(offset).Where("id in (?)", db.Table("bas_student_teach_classes").
		Select("student_id").Where("teach_class_id = ?", teachClass.ID)).Find(&students).Error

	return students, total, err

}

// AddStudentByClass 将一个班的学生 整体加入教学班
func (multiTableService *MultiTableService) AddStudentByClass(req request.AddStudentByClass) (err error) {
	teachClassId := req.TeachClassId
	classId := req.ClassId

	var teachClass basicdata.TeachClass
	teachClass.ID = teachClassId

	students, err := multiTableService.FindStudentsByClassId(classId)
	if err != nil {
		return err
	}

	err = global.GVA_DB.Model(&teachClass).Association("Student").Append(students)
	if err != nil {
		return err
	}

	return nil
}

// DeleteStudentByClass 将一个班的学生 整体移除教学班
func (multiTableService *MultiTableService) DeleteStudentByClass(req request.AddStudentByClass) error {
	teachClassId := req.TeachClassId
	classId := req.ClassId

	var teachClass basicdata.TeachClass
	teachClass.ID = teachClassId

	students, err := multiTableService.FindStudentsByClassId(classId)
	if err != nil {
		return err
	}

	err = global.GVA_DB.Model(&teachClass).Association("Student").Delete(students)
	if err != nil {
		return err
	}

	return nil
}

// FindStudentsByClassId 寻找一个班的学生集合
func (multiTableService *MultiTableService) FindStudentsByClassId(classId uint) ([]basicdata.Student, error) {
	var students []basicdata.Student

	err := global.GVA_DB.Where("class_id = ?", classId).Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (MultiTableService *MultiTableService) FindStudentByStudentClassId(classId uint) (res []uint, err error) {
	err = global.GVA_DB.Select("student_id").Where("teach_class_id=?", classId).Model(&basicdata.StudentAndTeachClass{}).Find(&res).Error
	return
}
