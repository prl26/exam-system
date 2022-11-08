package questionBank

import (
	background "github.com/prl26/exam-system/server/api/Backstage/questionBank/background"
	frontDesk "github.com/prl26/exam-system/server/api/frontDesk/questionBank"
)

type ApiGroup struct {
	background.Background
	frontDesk.FrontDesk
}