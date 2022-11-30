package utils

import "github.com/prl26/exam-system/server/model/examManage"

func Check(item []examManage.PaperTemplateItem) bool {
	var sum int
	for i := 0; i < len(item); i++ {
		score := *item[i].Score
		num := *item[i].Num
		temp := num * score
		sum += temp
	}
	if sum == 100 {
		return true
	}
	return false
}
