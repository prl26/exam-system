package request

import (
	"exam-system/model/common/request"
	"exam-system/model/teachplan"
)

type ScoreSearch struct {
	teachplan.Score
	request.PageInfo
}
