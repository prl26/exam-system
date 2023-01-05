package response

import "github.com/prl26/exam-system/server/model/basicdata"

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
type FrontResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}
type PageResultAndTerm struct {
	TermNow  basicdata.Term `json:"termNow"`
	List     interface{}    `json:"list"`
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}
