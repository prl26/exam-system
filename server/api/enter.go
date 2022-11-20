package api

import (
	"github.com/prl26/exam-system/server/api/backStage"

	"github.com/prl26/exam-system/server/api/frontDesk"

)

type ApiGroups struct {
	backStage.BackStage
	frontDesk.FrontDesk
}

var ApiGroupApp = new(ApiGroups)
