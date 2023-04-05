package basicdata

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/examManage/response"
)

func (t *TeachClassService) FindTeachClass(id uint, termId int) (teachClassAndLessons []response.TeachAndLessons, err error) {
	var teachClassIds []basicdata.StudentAndTeachClassAndTerm
	if termId != 0 {
		err = global.GVA_DB.Raw("SELECT teach_class_id,student_id FROM `bas_student_teach_classes` as s JOIN bas_teach_class as t ON t.term_id = ? and t.id = s.teach_class_id and t.course_id != 25 and s.student_id = ? GROUP BY teach_class_id", termId, id).Find(&teachClassIds).Error
	} else {
		err = global.GVA_DB.Raw("SELECT teach_class_id,student_id FROM `bas_student_teach_classes` as s JOIN bas_teach_class as t ON t.id = s.teach_class_id and t.course_id != 25 and s.student_id = ? GROUP BY teach_class_id", id).Find(&teachClassIds).Error
	}
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
func (t *TeachClassService) FindAllTeachClass(id uint, termId int) (teachClassAndLessons []response.TeachAndLessons, err error) {
	var teachClassIds []basicdata.StudentAndTeachClassAndTerm
	if termId != 0 {
		err = global.GVA_DB.Raw("SELECT teach_class_id,student_id FROM `bas_student_teach_classes` as s JOIN bas_teach_class as t ON t.term_id =? and t.id = s.teach_class_id  and s.student_id = ? GROUP BY teach_class_id", termId, id).Find(&teachClassIds).Error
	} else {
		err = global.GVA_DB.Raw("SELECT teach_class_id,student_id FROM `bas_student_teach_classes` as s JOIN bas_teach_class as t ON t.id = s.teach_class_id  and s.student_id = ? GROUP BY teach_class_id", id).Find(&teachClassIds).Error
	}
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

func (t *TeachClassService) FindTargetTeachClass(id uint, termId int) (teachClassAndLessons []response.TeachAndLessons, err error) {
	var teachClassIds []basicdata.StudentAndTeachClassAndTerm
	if termId != 0 {
		err = global.GVA_DB.Raw("SELECT teach_class_id,student_id FROM `bas_student_teach_classes` as s JOIN bas_teach_class as t ON t.term_id =? and t.id = s.teach_class_id and t.course_id = 25 and s.student_id = ? GROUP BY teach_class_id", termId, id).Find(&teachClassIds).Error
	} else {
		err = global.GVA_DB.Raw("SELECT teach_class_id,student_id FROM `bas_student_teach_classes` as s JOIN bas_teach_class as t ON t.id = s.teach_class_id and t.course_id = 25 and s.student_id = ? GROUP BY teach_class_id", id).Find(&teachClassIds).Error
	}
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
