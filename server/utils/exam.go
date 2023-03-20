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
func DiffArray(a []int64, b []int64) []int64 {
	var diffArray []int64
	temp := map[int64]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}
