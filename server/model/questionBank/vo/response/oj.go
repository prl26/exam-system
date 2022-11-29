package response

import (
	"github.com/prl26/exam-system/server/model/oj"
	"time"
)

type Execute struct {
	Output string `json:"output"` // 标准输出
	oj.ExecuteSituation
}

type Compile struct {
	FileId         string    `json:"fileId"`
	ExpirationTime time.Time `json:"expirationTime"`
}
