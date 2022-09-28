package questionBank

import (
	"fmt"
	testutils "github.com/prl26/exam-system/server/utils/test"
	"testing"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/27 20:05

 * @Note:

 **/
func TestMain(t *testing.M) {
	testutils.InitTest()
	t.Run()
}

var questionBankService QuestionBankService

func TestFindQuestions(t *testing.T) {
	questions := questionBankService.FindQuestions(2)
	fmt.Println(questions)
}
