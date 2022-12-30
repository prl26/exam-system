package response

import (
	"github.com/prl26/exam-system/server/global"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

type RangTopicSimple struct {
	global.GVA_MODEL
	questionBank.SimpleModel
}
