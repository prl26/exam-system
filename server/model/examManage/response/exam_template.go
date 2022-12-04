package response

type TemplateDetail struct {
	ChapterId   uint   `json:"chapterId" form:"chapterId" gorm:"chapter_id"`
	ChapterName string `json:"chapterName" form:"chapterName" gorm:"chapter_name"`
	ProblemType uint   `json:"problemType" form:"problemType"`
	Num         uint   `json:"num"`
}
type Template struct {
	Choice  []TemplateDetail `json:"choice"`
	Judge   []TemplateDetail `json:"judge"`
	Blank   []TemplateDetail `json:"blank"`
	Program []TemplateDetail `json:"program"`
}
