package utils

import "github.com/prl26/exam-system/server/model/examManage"

func Check(Papertemplate examManage.PaperTemplate) bool {
	sum := 0
	for _, v := range Papertemplate.PaperTemplateItems {
		sum += *v.Score
	}
	if sum == 100 {
		return true
	}
	return false
}
