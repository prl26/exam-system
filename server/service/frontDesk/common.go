package frontDesk

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/system"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type CommonService struct {
}

func (commonService *CommonService) FindTeachClass(id uint) (teachClassAndLesson []response.TeachAndLessons, err error) {
	var teachClassIds []basicdata.StudentAndTeachClass
	err = global.GVA_DB.Where("id = ?", id).Find(&teachClassIds).Error
	for i := 0; i < len(teachClassIds); i++ {
		var teachClass basicdata.TeachClass
		var user system.SysUser
		err = global.GVA_DB.Where("id = ?", teachClassIds[i]).Find(&teachClass).Error
		if err != nil {
			return
		}
		teachClassAndLesson[i].TeachClassId = teachClass.ID
		teachClassAndLesson[i].TeachClassName = teachClass.Name
		teachClassAndLesson[i].LessonId = uint(*teachClass.CourseId)
		var Lessons basicdata.Lesson
		err = global.GVA_DB.Where("id = ?", teachClass.CourseId).Find(&Lessons).Error
		if err != nil {
			return
		}
		teachClassAndLesson[i].NameOfLesson = Lessons.Name
		err = global.GVA_DB.Where("id = ?", teachClass.TeacherId).Find(&user).Error
		teachClassAndLesson[i].TeacherName = user.NickName
	}
	return
}
func (commonService *CommonService) FindExamPlans(teachClassId uint) (examPlans []teachplan.ExamPlan, err error) {
	err = global.GVA_DB.Where("teach_class_id = ?", teachClassId).Find(&examPlans).Error
	return
}
