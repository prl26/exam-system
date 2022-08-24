package main

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
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
	a := 1
	fakeChoiceQuestion(&a, 20)

}

func fakeChoiceQuestion(chapterId *int, n int) {
	choices := make([]*questionBank.MultipleChoice, n)
	options := make([]*questionBank.Options, 0, 4*n)
	merge := make([]*questionBank.ChapterMerge, 0, n)
	for i := 0; i < n; i++ {
		choices[i] = &questionBank.MultipleChoice{}
		randInt := RandInt(10, 21)
		choices[i].Title = randChineseStr(int(randInt) - 7)
		choices[i].Describe = randChineseStr(int(randInt)) + "?"
		mostOptions := 1
		choices[i].MostOptions = &mostOptions
		correct := int(RandInt(1, 4))
		choices[i].Answer = fmt.Sprintf("%d", correct)
	}
	global.GVA_DB.Create(choices)
	for _, choice := range choices {
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
		k := &questionBank.ChapterMerge{}
		canExam := RandInt(0, 1)
		k.CanExam = &canExam
		canPractice := RandInt(0, 1)
		k.CanPractice = &canPractice
		difficulty := RandInt(0, 3)
		k.Difficulty = &difficulty
		k.QuestionId = &u
		k.ChapterId = chapterId
		// 选择题TYPE的id有需求可以改一下
		k.QuestionType = &u
		merge = append(merge, k)
	}
	global.GVA_DB.Create(options)
	global.GVA_DB.Create(merge)
}

func randChineseStr(n int) string {
	a := make([]rune, n)
	for i := range a {
		a[i] = rune(RandInt(19968, 20000))

	}
	return string(a)
}
func RandInt(min, max int64) int {
	return int(min + rand.Int63n(max-min))
}
