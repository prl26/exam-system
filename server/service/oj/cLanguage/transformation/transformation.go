package main

import (
	"fmt"
	"github.com/prl26/exam-system/server/model/enum/language"
	"github.com/prl26/exam-system/server/model/enum/questionType"
	"github.com/prl26/exam-system/server/model/questionBank"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/27 16:13

 * @Note:

 **/

var ProgrammType uint = 2
var LanguageId = newPointInt(1)
var notCanExam = newPointInt(1)

func getDB(dsn string) *gorm.DB {
	mysqlFrom := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	if result, err := gorm.Open(mysql.New(mysqlFrom)); err != nil {
		panic(err)
	} else {
		return result
	}
}

type Result struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"describe" gorm:"column:describe"`
	ProblemType    int    `json:"problem_type"`
	Stage          int    `json:"stage"`
	Code           string `json:"code"`
	IsProgramBlank int    `json:"is_program_blank"`
	CanPractice    int
	CanExam        int
}

type TestCase struct {
	Input  string `json:"input" gorm:"column:TestCaseInput"`
	Output string `json:"output" gorm:"column:TestCaseOutput"`
	Score  int    `json:"score" gorm:"column:ScoreWeight"`
}

func main() {
	from := getDB("root:cuitexamloopmysql123@tcp(exam.cuit.edu.cn:3306)/stu_system?charset=utf8")
	to := getDB("root:chensida2318@tcp(47.108.150.32:3306)/gva?charset=utf8")
	transformation(from, to)
	defer func() {
		db, _ := from.DB()
		db.Close()
		db, _ = to.DB()
		db.Close()
	}()
}

func newPointInt(i int) *int {
	return &i
}

func transformation(from *gorm.DB, to *gorm.DB) {
	results := make([]Result, 0, 500)
	from.Raw("select `QuestionBh` as \"id\"," +
		"`name` as 'title'," +
		"`Description` as `describe`," +
		"CASE WHEN Difficulty='1000403' THEN '2'\nWHEN Difficulty='1000402' THEN '1'\nELSE '3' END AS 'problem_type' , " +
		"(Stage%10)+1 as `stage`," +
		"CASE WHEN JSON_VALID(SourceCode) THEN JSON_UNQUOTE(JSON_EXTRACT(SourceCode, \"$.key[0].code\"))  ELSE null END as `code`," +
		"IsProgramBlank='100001' as `is_program_blank`," +
		"Checked='100001' as 'can_exam', " +
		"Score as 'can_practice' " +
		" from stu_system.questions  " +
		" where CourseID=2000301 and QuestionType=1000206").Find(&results)
	cases := make([]*questionBank.ProgrammCase, 0, 2000)
	chapterMerges := make([]*questionBank.ChapterMerge, 0, 500)
	for i, result := range results {
		// 创建 programm
		programm := questionBank.Programm{}
		programm.Title = result.Title
		programm.Describe = result.Description
		programm.CanExam = &result.CanExam
		programm.CanPractice = &result.CanPractice
		programm.ProblemType = result.ProblemType
		if err := to.Create(&programm).Error; err != nil {
			panic(err)
		}

		// 创建 编程题 语言支持
		merge := questionBank.ProgrammLanguageMerge{}
		merge.LanguageId = language.C_LANGUAGE
		if result.IsProgramBlank == 0 {
			merge.ReferenceAnswer = result.Code
		} else {
			merge.DefaultCode = result.Code
		}
		merge.ProgrammId = programm.ID
		if err := to.Create(&merge).Error; err != nil {
			panic(err)
		}
		var testCase []TestCase
		if err := from.Table("testcase").Where("QuestionId", result.Id).Find(&testCase).Error; err != nil {
			panic(err)
		}

		// 加入测试用例集合
		for i, t := range testCase {
			thisCase := &questionBank.ProgrammCase{}
			thisCase.Name = fmt.Sprintf("测试用例-%d", i)
			thisCase.Input = t.Input
			thisCase.Output = t.Output
			thisCase.LanguageId = language.C_LANGUAGE
			thisCase.ProgrammId = programm.ID
			thisCase.Score = uint(t.Score)
			cases = append(cases, thisCase)
		}

		// 加入章节绑定
		chapterMerge := &questionBank.ChapterMerge{}
		chapterMerge.ChapterId = uint(result.Stage)
		chapterMerge.QuestionId = programm.ID
		chapterMerge.QuestionType = questionType.PROGRAM

		chapterMerges = append(chapterMerges, chapterMerge)

		fmt.Println("已经处理完", i+1, "个编程题勒！")
	}

	if err := to.CreateInBatches(cases, 50).Error; err != nil {
		panic(err)
	}
	if err := to.CreateInBatches(chapterMerges, 50).Error; err != nil {
		panic(err)
	}
	fmt.Println("已经处理完了", len(cases), "个编程用例")
}
