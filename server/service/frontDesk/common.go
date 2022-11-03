package frontDesk

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/system"
)

type CommonService struct {
}

func (commonService *CommonService) FindTeachClass(id uint) (teachClassAndLessons []response.TeachAndLessons, err error) {
	var teachClassIds []basicdata.StudentAndTeachClass
	err = global.GVA_DB.Table("bas_student_teach_classes").Where("student_id = ?", id).Find(&teachClassIds).Error
	if err != nil {
		return
	}
	for i := 0; i < len(teachClassIds); i++ {
		var teachClass basicdata.TeachClass
		var user system.SysUser
		err = global.GVA_DB.Where("id = ?", teachClassIds[i].TeachClassId).Find(&teachClass).Error
		if err != nil {
			return
		}

		var Lessons basicdata.Lesson
		err = global.GVA_DB.Where("id = ?", teachClass.CourseId).Find(&Lessons).Error
		if err != nil {
			return
		}
		err = global.GVA_DB.Where("id = ?", teachClass.TeacherId).Find(&user).Error
		teachClassAndLesson := response.TeachAndLessons{
			TeachClassId:   teachClass.ID,
			TeachClassName: teachClass.Name,
			NameOfLesson:   Lessons.Name,
			LessonId:       uint(*teachClass.CourseId),
			TeacherName:    user.NickName,
		}
		teachClassAndLessons = append(teachClassAndLessons, teachClassAndLesson)
	}
	return
}
