package response

import "github.com/prl26/exam-system/server/model/examManage"

type PaperStatus struct {
	Name                          string `json:"name"`
	IsReport                      bool   `json:"isReport"`
	examManage.StudentPaperStatus `json:"Status"`
}
