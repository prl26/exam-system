package multipleChoice

import (
	"github.com/prl26/exam-system/server/model/questionBank/po"
	testutils "github.com/prl26/exam-system/server/utils/test"
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
	testutils.InitTest()
}

func TestCheck(t *testing.T) {
	question := &po.MultipleChoice{
		MostOptions: getIntPoint(2),
		Answer:      "1,2",
	}
	log.Println(server.check(question, []int{1}))
}

func getUintPoint(a uint) *uint {
	return &a
}

func getIntPoint(a int) *int {
	return &a
}
