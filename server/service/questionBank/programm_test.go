package questionBank

import (
	"fmt"
	"github.com/prl26/exam-system/server/model/questionBank"
	responese "github.com/prl26/exam-system/server/model/questionBank/response"
	testutils "github.com/prl26/exam-system/server/utils/test"
	"testing"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 15:10

 * @Note:

 **/

func testMain(m *testing.M) {
	testutils.InitTest()
	m.Run()
}

var programmService ProgrammService

func TestProgrammService_FindProgramDetail(t *testing.T) {
	var s questionBank.Programm
	err := programmService.FindProgramDetail(&s, 500)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestProgrammService_EditProgramDetail(t *testing.T) {
	var s questionBank.Programm
	err := programmService.FindProgramDetail(&s, 500)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Title += "?"
	err = programmService.EditProgrammDetail(&s)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestProgrammService_DeleteProgrammDetail(t *testing.T) {
	err := programmService.DeleteProgramm([]int{901})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestProgrammService_FindProgrammCases(t *testing.T) {
	var s []questionBank.ProgrammCase
	err := programmService.FindProgrammCases(&s, 500, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
func TestProgrammService_FindLanguageSupport(t *testing.T) {
	var s []responese.LanguageSupport
	err := programmService.FindLanguageSupport(&s, 500)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
func TestProgrammService_EditProgrammCases(t *testing.T) {
	var s []questionBank.ProgrammCase
	err := programmService.FindProgrammCases(&s, 500, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	s[0].Name += "?"
	err = programmService.EditProgrammCases(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestProgrammService_AddProgrammCase(t *testing.T) {

}

func TestProgrammService_DeleteProgrammCase(t *testing.T) {

}

func TestProgrammService_EditLanguageSupport(t *testing.T) {
	var s []responese.LanguageSupport
	err := programmService.FindLanguageSupport(&s, 500)
	if err != nil {
		fmt.Println(err)
		return
	}
	s[0].ReferenceAnswer += " "
	merge := questionBank.ProgrammLanguageMerge{}
	merge.ID = s[0].ID
	merge.LanguageId = s[0].LanguageId
	merge.ReferenceAnswer = s[0].ReferenceAnswer + " "
	merge.DefaultCode = s[0].DefaultCode
	merge.DefaultCode = s[0].DefaultCode
	err = programmService.EditLanguageSupport(&merge)
	if err != nil {
		fmt.Println(err)
		return
	}
}
