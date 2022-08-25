/*
*

	@author: qianyi  2022/8/24 18:56:00
	@note:
*/
package basicdata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/basicdata/request"
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
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&basicdata.Student{})

	var teachClass basicdata.TeachClass
	teachClass.ID = info.TeachClassId

	var students []basicdata.Student

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Model(&teachClass).Association("Student").Find(&students)
	return students, total, err

}
