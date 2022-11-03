package frontDesk

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/examManage/response"
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
		var teachClassAndLesson response.TeachAndLessons
		err = global.GVA_DB.Raw("select c.id TeachClassId,c.`name` as TeachClassName,l.`name` as NameOfLesson,l.id as LessonId,u.nick_name as TeacherName from bas_teach_class c ,bas_lesson l,sys_users u where c.id = ? and l.id = c.course_id and u.id= c.teacher_id", teachClassIds[i].TeachClassId).
			Scan(&teachClassAndLesson).Error
		if err != nil {
			return
		}
		teachClassAndLessons = append(teachClassAndLessons, teachClassAndLesson)
	}
	return
}
