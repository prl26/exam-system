package utils

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
)

func GetTermName(id uint) string {
	var result basicdata.Term
	global.GVA_DB.Model(basicdata.Term{}).Where("id = ?", id).Find(&result)
	return result.Name
}
func GetLessonName(id int) string {
	var result basicdata.Lesson
	global.GVA_DB.Model(basicdata.Lesson{}).Where("id = ?", id).Find(&result)
	return result.Name
}
func GetTeachPlanName(id uint) string {
	var result basicdata.TeachClass
	global.GVA_DB.Model(basicdata.TeachClass{}).Where("id = ?", id).Find(&result)
	return result.Name
}
