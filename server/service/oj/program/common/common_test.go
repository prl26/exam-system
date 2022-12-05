package common

import (
	"fmt"
	testutils "github.com/prl26/exam-system/server/utils/test"
	"log"
	"testing"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 0:15

 * @Note:

 **/
var server CommonService

func TestMain(m *testing.M) {
	testutils.InitTest()
}

func TestFindProgrammCase(t *testing.T) {
	programmCase, err := server.FindProgrammCase(1, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(programmCase)
}
