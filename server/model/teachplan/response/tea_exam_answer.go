package response

type HistoryItem struct {
	Exist  bool   `json:"exist"`
	Score  uint   `json:"score"`
	Answer string `json:"answer"`
}
type History struct {
	History map[uint]*HistoryItem `json:"history"`
}
