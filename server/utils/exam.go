package utils

import "github.com/prl26/exam-system/server/global"

func Check(id *int) bool {
	var sum int
	global.GVA_DB.Raw("SELECT sum(num*score) from exam_paper_template_item where template_id = ?", id).Scan(&sum)
	if sum == 100 {
		return true
	}
	return false
}
