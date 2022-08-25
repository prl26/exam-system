package multipleChoice

import (
	"exam-system/core"
	"exam-system/global"
	"exam-system/initialize"
	"exam-system/model/questionBank"
	"log"
	"testing"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 19:01

 * @Note:

 **/

var server *MultipleChoiceService

func TestMain(m *testing.M) {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_DB = initialize.Gorm()
	m.Run()
}

func TestCheck(t *testing.T) {
	question := &questionBank.MultipleChoice{
		MostOptions: getIntPoint(2),
		Answer:      "1,2",
	}
	log.Fatalln(server.check(question, []int{1}))
}

func getUintPoint(a uint) *uint {
	return &a
}

func getIntPoint(a int) *int {
	return &a
}
