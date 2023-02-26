package response

import "github.com/prl26/exam-system/server/model/examManage"

type RecordRp struct {
	ExamRecord  examManage.ExamRecord        `json:"examRecord"`
	RecorcMerge []examManage.ExamRecordMerge `json:"recorcMerge"`
}
