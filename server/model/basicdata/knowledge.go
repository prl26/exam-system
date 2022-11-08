package basicdata

type Knowledge struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ChapterId   uint   `json:"chapter_id"`
}

func (Knowledge) TableName() string {
	return "bas_knowledge"
}
