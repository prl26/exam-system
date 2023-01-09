package lessondata

type Knowledge struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ChapterId   uint   `json:"chapterId" form:"chapterId"`
	LessonId    uint   `json:"lessonId" form:"lessonId"`
}

func (Knowledge) TableName() string {
	return "bas_knowledge"
}
