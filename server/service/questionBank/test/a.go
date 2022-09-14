package main

import (
	"github.com/prl26/exam-system/server/core"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/initialize"
	"github.com/prl26/exam-system/server/service/questionBank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/14 15:34

 * @Note:

 **/

var server questionBank.QuestionBankService

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_DB = initialize.Gorm()
	server.FindQuestions(1)
}
