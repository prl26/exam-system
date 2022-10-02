package questionBank

import (
	background "github.com/prl26/exam-system/server/api/v1/questionBank/background"
	frontDesk "github.com/prl26/exam-system/server/api/v1/questionBank/frontDesk"
)

type ApiGroup struct {
	background.Background
	frontDesk.FrontDesk
}
