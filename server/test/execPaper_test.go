package test

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	examManage2 "github.com/prl26/exam-system/server/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func openDB() (db *gorm.DB) {
	dsn := "root:cuit@123456@tcp(139.9.249.149)/Lectures?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	},
	), &gorm.Config{})
	return
}
func TestExecPapers(t *testing.T) {
	global.GVA_DB = openDB()
	string1 := []string{"A"}
	string2 := []string{"B"}
	string3 := []string{"66666"}

	chioce := []examManage.MultipleChoiceCommit{
		{MergeId: 2481, QuestionId: 240, Answer: string2},
		{MergeId: 2484, QuestionId: 145, Answer: string2},
		{MergeId: 2486, QuestionId: 81, Answer: string2},
		{MergeId: 2490, QuestionId: 94, Answer: string1},
		{MergeId: 2492, QuestionId: 150, Answer: string1},
		{MergeId: 2494, QuestionId: 213, Answer: string1},
	}
	judge := []examManage.JudgeCommit{
		{MergeId: 2478, QuestionId: 14, Answer: false},
		{MergeId: 2480, QuestionId: 2, Answer: false},
		{MergeId: 2485, QuestionId: 5, Answer: false},
		{MergeId: 2487, QuestionId: 35, Answer: false},
		{MergeId: 2491, QuestionId: 8, Answer: false},
	}
	blank := []examManage.BlankCommit{
		{MergeId: 2488, QuestionId: 30, Answer: string3},
		{MergeId: 2493, QuestionId: 5, Answer: string3},
	}
	examCommit := examManage.CommitExamPaper{
		StudentId:            2020131039,
		PlanId:               27,
		PaperId:              211,
		MultipleChoiceCommit: chioce,
		JudgeCommit:          judge,
		BlankCommit:          blank,
	}
	examManage2.ExecPapers(examCommit)

}
