package response

import "github.com/prl26/exam-system/server/model/examManage"

type AllPaperMerge struct {
	PaperMerge []examManage.PaperQuestionMerge `json:"paperMerge"`
}
