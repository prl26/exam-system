package response

import (
	ojBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"time"
)

type Execute struct {
	Output string `json:"output"` // 标准输出
	ojBo.ExecuteSituation
}

type Compile struct {
	FileId         string    `json:"fileId"`
	ExpirationTime time.Time `json:"expirationTime"`
}

type SubmitResponse struct {
	Score  uint           `json:"score"`
	Submit []*ojBo.Submit `json:"submit"`
}
