package response

type TeachAndLessons struct {
	TeachClassId   uint   `json:"teachClassId"`
	TeachClassName string `json:"teachClassName"`
	NameOfLesson   string `json:"nameOfLesson"`
	LessonId       uint   `json:"lessonId"`
	TeacherName    string `json:"teacherName"`
}
