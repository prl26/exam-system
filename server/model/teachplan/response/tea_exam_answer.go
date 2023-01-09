package response

type HistoryItem struct {
	Exist bool `json:"exist"`
	Score uint `json:"score"`
}
type History struct {
	History map[uint]*HistoryItem `json:"history"`
}
