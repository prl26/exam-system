package router

import (
	"github.com/prl26/exam-system/server/router/backStage"
	"github.com/prl26/exam-system/server/router/frontDesk"
)

type RouterGroup struct {
	BackStage    backStage.BackStage
	FrontDesk    frontDesk.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
