/*
*

	@author: qianyi  2022/8/24 18:56:00
	@note:
*/
package basicdata

import (
	"exam-system/global"
	"exam-system/model/basicdata"
	"exam-system/model/basicdata/request"
)

type MultiTableService struct {
}

// 向教学班中 加入学生的关联
func (multiTableService *MultiTableService) InitTeachClassStudents(info request.StuTeachClass) error {

	var teachClass basicdata.TeachClass
	teachClass.ID = info.TeachClassId

	n := len(info.StudentIds)
	students := make([]basicdata.Student, n)

	for i := 0; i < n; i++ {
		students[i].ID = info.StudentIds[i]
	}

	err := global.GVA_DB.Model(&teachClass).Association("Student").Append(students)

	return err
}

// 在教学班中 移除学生
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

// 获取教学班的学生
func (multiTableService *MultiTableService) GetTeachClassStudentInfo(info request.TeachClassStudent) (list []basicdata.Student, total int64, err error) {

	var limit, offset int

	if info.PageSize <= 0 {
		limit = 10
	} else {
		limit = info.PageSize
	}

	if info.Page <= 0 {
		offset = 0
	} else {
		offset = info.PageSize * (info.Page - 1)
	}

	var teachClass basicdata.TeachClass
	teachClass.ID = info.TeachClassId

	var students []basicdata.Student

	db := global.GVA_DB

	//err = db.Model(&teachClass).Association("Student").Find(&students) 高端写法但是无法分页
	db.Limit(limit).Offset(offset).Where("id in (?)", db.Table("bas_student_teach_classes").
		Select("student_id").Where("teach_class_id = ?", teachClass.ID)).Find(&students).Count(&total)

	return students, total, err

}
