package main

import (
	"fmt"
	"github.com/prl26/exam-system/server/core"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/initialize"
	"github.com/prl26/exam-system/server/model/questionBank"
	"go.uber.org/zap"
	"math/rand"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 3:22

 * @Note:

 **/
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	fakeChoiceQuestion(getPoint(1), 20)
	fakeSupplyBlankQuestion(getPoint(1), 20)
	fakeJudgeQuestion(getPoint(1), 20)
}

func buildBasicQuestion(chapterId *int, question *questionBank.BasicModel) {
	question.ChapterId = chapterId
	question.ProblemType = getPoint(RandInt(1, 3))
	question.CanPractice = getPoint(RandInt(0, 1))
	question.CanExam = getPoint(RandInt(0, 1))
	randInt := RandInt(10, 21)
	question.Describe = randChineseStr(randInt-4) + "?"
	question.Title = randChineseStr(randInt - 4)
}

func fakeSupplyBlankQuestion(chapterId *int, n int) {
	questions := make([]*questionBank.SupplyBlank, n)
	for i := 0; i < len(questions); i++ {
		questions[i] = &questionBank.SupplyBlank{}
		questions[i].IsOrder = getPoint(RandInt(0, 1))
		questions[i].Num = getPoint(RandInt(1, 3))
		buildBasicQuestion(chapterId, &questions[i].BasicModel)
	}
	global.GVA_DB.Create(questions)
}

func fakeJudgeQuestion(chapterId *int, n int) {
	questions := make([]*questionBank.Judge, n)
	for i := 0; i < n; i++ {
		questions[i] = &questionBank.Judge{}
		questions[i].IsRight = getPoint(RandInt(0, 1))
		buildBasicQuestion(chapterId, &questions[i].BasicModel)
	}
	global.GVA_DB.Create(questions)
}
func fakeChoiceQuestion(chapterId *int, n int) {
	questions := make([]*questionBank.MultipleChoice, n)
	options := make([]*questionBank.Options, 0, 4*n)
	for i := 0; i < n; i++ {
		questions[i] = &questionBank.MultipleChoice{}
		randInt := RandInt(10, 21)
		questions[i].Title = randChineseStr(int(randInt) - 7)
		questions[i].Describe = randChineseStr(randInt) + "?"
		mostOptions := 1
		questions[i].MostOptions = &mostOptions
		correct := RandInt(1, 4)
		questions[i].Answer = fmt.Sprintf("%d", correct)
		buildBasicQuestion(chapterId, &questions[i].BasicModel)
	}
	global.GVA_DB.Create(questions)
	for _, choice := range questions {
		randInt := RandInt(3, 5)
		u := int(choice.ID)
		for j := 0; j < 4; j++ {
			this := &questionBank.Options{}
			options = append(options, this)
			this.MultipleChoiceId = &u
			this.Describe = randChineseStr(int(randInt))
			order := j
			this.Orders = &order
		}
	}
	global.GVA_DB.Create(options)
}

func randChineseStr(n int) string {
	a := make([]rune, n)
	for i := range a {
		a[i] = rune(RandInt(19968, 20000))

	}
	return string(a)
}
func RandInt(min, max int64) int {
	return int(min + rand.Int63n(max-min+1))
}

func getPoint(a int) *int {
	return &a
}
