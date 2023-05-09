package lessondata

type Knowledge struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	Description   string `json:"description" from:"description"`
	ChapterId     uint   `json:"chapterId" form:"chapterId"`
	LessonId      uint   `json:"lessonId" form:"lessonId"`
	QuestionCount *int64 `json:"questionCount,omitempty"`
	DoneCount     *int64 `json:"doneCount,omitempty"`
}

func (Knowledge) TableName() string {
	return "bas_knowledge"
}
