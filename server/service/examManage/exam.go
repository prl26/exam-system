package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/examManage/response"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type ExamService struct {
}

func (examServices *ExamService) FindTeachClass(id uint) (teachClassAndLesson []response.TeachAndLessons, err error) {
	var teachClassIds []basicdata.StudentAndTeachClass
	var nameOfLessons []string
	err = global.GVA_DB.Where("id = ?", id).Find(&teachClassIds).Error
	for i := 0; i < len(teachClassIds); i++ {
		var teachClass basicdata.TeachClass
		err = global.GVA_DB.Where("id = ?", teachClassIds[i]).Find(&teachClass).Error
		if err != nil {
			return
		}
		var Lessons basicdata.Lesson
		err = global.GVA_DB.Where("id = ?", teachClass.CourseId).Find(&Lessons).Error
		if err != nil {
			return
		}
		nameOfLessons[i] = Lessons.Name
	}
	return
}
func (examServices *ExamService) FindExamPlans(teachClassId uint) (examPlans []teachplan.ExamPlan, err error) {
	err = global.GVA_DB.Where("teach_class_id = ?", teachClassId).Find(&examPlans).Error
	return
}
