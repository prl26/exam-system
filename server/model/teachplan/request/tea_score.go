package request

import (
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type ScoreSearch struct {
	teachplan.Score
	request.PageInfo
}
