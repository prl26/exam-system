package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/lessondata"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/enum/languageType"
	"github.com/prl26/exam-system/server/model/questionBank/enum/problemType"
	questionBankError "github.com/prl26/exam-system/server/model/questionBank/error"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

/*
*

  - @Author: AloneAtWar

  - @Date:   2022/9/27 16:13

  - @Note:

    *
*/
var courseName = "C语言程序设计" // 课程的名称

var truePtr = newPointInt(1)
var falsePtr = newPointInt(0)

func newPointInt(i int) *int {
	return &i
}

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

func CreateLesson(courseName string, db *gorm.DB) *basicdata.Lesson {
	_true := true
	var chapters []*basicdata.Chapter
	for i := 0; i < 10; i++ {
		chapters = append(chapters, &basicdata.Chapter{
			Name: fmt.Sprintf("第%d阶段", i+1),
		})
	}
	lesson := basicdata.Lesson{
		Name:             courseName,
		OpenQuestionBank: &_true,
		Chapters:         chapters,
	}
	_ = db.Create(&lesson).Error
	return &lesson
}

func CreateKnowledgeTable(courseId string, lesson *basicdata.Lesson, from *gorm.DB, to *gorm.DB) map[string]*lessondata.Knowledge {
	table := make(map[string]*lessondata.Knowledge)
	type knowledgepoint struct {
		KnowledgeBh   string
		Description   string
		KnowledgeName string
		Chapter       uint
	}
	var knowledgepoints []knowledgepoint
	_ = from.Where("CourseID=?", courseId).Order("Chapter").Find(knowledgepoints).Error
	for _, knowledge := range knowledgepoints {
		l := lessondata.Knowledge{
			Name:        knowledge.KnowledgeName,
			Description: knowledge.Description,
			ChapterId:   lesson.Chapters[knowledge.Chapter%10+1].ID,
			LessonId:    lesson.ID,
		}
		_ = to.Create(&l).Error
		table[knowledge.KnowledgeBh] = &l
	}
	return table
}

// 对于老题库表中的字段
type Questions struct {
	QuestionBh     string              //	主键
	CustomBh       string              //	自定义编号
	Name           string              //	标题
	Description    string              //	描述
	Difficulty     QuestionsDifficulty //  	难度
	Stage          uint                //  	阶段
	Score          uint                //  	是否可以练习 0为不能训练 1为可以训练
	KnowledgeBh    string              //  	知识点编号
	Checked        uint                //  	是否已审核
	Answer         string              //  	答案
	IsProgramBlank string              //  	是否是程序填空题
	Code           string              //  	代码
	Momo           string              // 	如果为程序填空题 填空内的真正答案会存在这里  // TODO 暂时未进行处理
	QuestionType   string              //   题目类型
}

func (q Questions) ToBasicQuestion() *questionBank.BasicModel {
	var canPractice *int
	if q.Score == 1 {
		canPractice = truePtr
	} else {
		canPractice = falsePtr
	}
	var isCheck *int
	if q.Checked == 100001 {
		isCheck = truePtr
	} else {
		isCheck = falsePtr
	}
	model := questionBank.BasicModel{
		SimpleModel: questionBank.SimpleModel{
			SerNo:       q.QuestionBh,
			ProblemType: q.Difficulty.toProblemType(),
			PracticeExamSupport: questionBank.PracticeExamSupport{
				IsCheck:     isCheck,
				CanPractice: canPractice,
				CanExam:     truePtr,
			},
			Title: q.Name,
		},
		Describe: q.Description,
	}
	return &model
}

func (q Questions) ToJudgeModel(db *gorm.DB) *questionBank.JudgeModel {
	baseQuestion := q.ToBasicQuestion()
	var isRight *bool
	answer := strings.Trim(q.Answer, " ")
	if answer == "1" {
		var _true = true
		isRight = &_true
	} else {
		var _false = false
		isRight = &_false
	}

	model := &questionBank.JudgeModel{
		IsRight:    isRight,
		BasicModel: *baseQuestion,
	}
	return model
}

func (q Questions) ToMultipleChoiceModel(db *gorm.DB) *questionBank.MultipleChoiceModel {
	baseQuestion := q.ToBasicQuestion()

	var mostOption = 1
	model := &questionBank.MultipleChoiceModel{
		BasicModel:   *baseQuestion,
		Answer:       q.Answer,
		MostOptions:  &mostOption,
		IsIndefinite: falsePtr,
	}
	return model
}

func (q Questions) ToSupplyBlankModel(db *gorm.DB) *questionBank.SupplyBlankModel {
	baseQuestion := q.ToBasicQuestion()
	type Blank struct {
		Answer     string
		Proportion string
	}
	var blanks []*Blank
	db.Raw("select * from apfill where QuestionBh=?", q.QuestionBh).Find(&blanks)
	var proportions []string
	var answers []string
	for _, blank := range blanks {
		proportions = append(proportions, blank.Proportion)
		answers = append(answers, blank.Answer)
	}
	num := len(blanks)
	proportion := strings.Join(proportions, ",")
	answer := strings.Join(answers, ",")

	model := &questionBank.SupplyBlankModel{
		BasicModel: *baseQuestion,
		IsOrder:    truePtr,
		Num:        &num,
		Proportion: proportion,
		Answer:     answer,
	}
	return model
}

func (q Questions) ToProgramModel(db *gorm.DB) *questionBank.ProgramModel {
	baseQuestion := q.ToBasicQuestion()
	type TestCase struct {
		Input  string `json:"input" gorm:"column:TestCaseInput"`
		Output string `json:"output" gorm:"column:TestCaseOutput"`
		Score  int    `json:"score" gorm:"column:ScoreWeight"`
	}
	var testCase []TestCase
	if err := db.Table("testcase").Where("QuestionId", q.QuestionBh).Find(&testCase).Error; err != nil {
		panic(err)
	}
	cases := questionBankBo.ProgramCases{}
	for i, t := range testCase {
		programCase := &questionBankBo.ProgramCase{
			Name:   fmt.Sprintf("第%d个用例", i+1),
			Score:  uint(t.Score),
			Input:  t.Input,
			Output: t.Output,
		}
		cases = append(cases, programCase)
	}
	programCase, err := cases.Serialize()
	if err != nil {
		if errors.As(err, questionBankError.ScoreError) {
			return nil
		}
		return nil
	}
	supports := questionBankBo.LanguageSupports{}
	supports = append(supports, &questionBankBo.LanguageSupport{
		LanguageId: languageType.C_LANGUAGE,
		LanguageLimit: questionBankBo.LanguageLimit{
			StrictMemoryLimit: nil,
			MemoryLimit:       nil,
			CpuLimit:          nil,
			ClockLimit:        nil,
			StackLimit:        nil,
			ProcLimit:         nil,
			CpuRateLimit:      nil,
			CpuSetLimit:       nil,
		},
	})
	languageSupport, languageSupportsBrief, _ := supports.Serialize()

	defaultCode := ""
	referenceAnswer := ""
	if q.IsProgramBlank == "100001" {
		defalutCodes := questionBankBo.DefaultCodes{}
		defalutCodes = append(defalutCodes, &questionBankBo.DefaultCode{
			LanguageId: languageType.C_LANGUAGE,
			Code:       q.Code,
		})
		defaultCode, _ = defalutCodes.Serialize()
	} else {
		referenceAnswers := questionBankBo.ReferenceAnswers{}
		referenceAnswers = append(referenceAnswers, &questionBankBo.ReferenceAnswer{
			LanguageId: languageType.C_LANGUAGE,
			Code:       q.Code,
		})
		referenceAnswer, _ = referenceAnswers.Serialize()
	}

	model := &questionBank.ProgramModel{
		BasicModel:            *baseQuestion,
		ProgramCases:          &programCase,
		LanguageSupports:      &languageSupport,
		ReferenceAnswers:      &referenceAnswer,
		DefaultCodes:          &defaultCode,
		LanguageSupportsBrief: &languageSupportsBrief,
	}
	return model
}

type QuestionsDifficulty uint

func (this QuestionsDifficulty) toProblemType() problemType.ProblemType {
	switch this {
	case 100402:
		return problemType.EASY
	case 100403:
		return problemType.MEDIUM
	case 100404:
		return problemType.HARD
	}
	log.Panicf("错误的QuestionsDifficulty: %d", this)
	return 0
}

func main() {
	from := getDB("root:cuitexamloopmysql123@tcp(exam.cuit.edu.cn:3306)/stu_system?charset=utf8")
	to := getDB("root:cuit@123456@tcp(139.9.249.149:3306)/gva?charset=utf8")
	defer func() {
		db, _ := from.DB()
		db.Close()
		db, _ = to.DB()
		db.Close()
	}()
	lesson := CreateLesson("C语言程序设计", to)
	table := CreateKnowledgeTable("2000301", lesson, from, to)
	judges := []*questionBank.Judge{}
	programs := []*questionBank.Program{}
	supplyBlanks := []*questionBank.SupplyBlank{}
	choices := []*questionBank.MultipleChoice{}
	questions := []*Questions{}
	_ = from.Where("CourseID=?", "200301").Find(&questions).Error
	for _, question := range questions {
		knowledge := table[question.KnowledgeBh]
		courseSupport := questionBank.CourseSupport{
			LessonId:    lesson.ID,
			ChapterId:   knowledge.ChapterId,
			KnowledgeId: knowledge.ID,
		}
		switch question.QuestionType {
		// 判断
		case "1000203":
			model := question.ToJudgeModel(from)
			judge := &questionBank.Judge{
				CourseSupport: courseSupport,
				JudgeModel:    *model,
			}
			judges = append(judges, judge)
		// 填空
		case "1000204":
			model := question.ToSupplyBlankModel(from)
			supplyBlank := &questionBank.SupplyBlank{
				CourseSupport:    courseSupport,
				SupplyBlankModel: *model,
			}
			supplyBlanks = append(supplyBlanks, supplyBlank)
		// 编程
		case "1000206":
			model := question.ToProgramModel(from)
			program := &questionBank.Program{
				CourseSupport: courseSupport,
				ProgramModel:  *model,
			}
			programs = append(programs, program)
		// 选择题
		case "100020101":
			model := question.ToMultipleChoiceModel(from)
			multipleChoice := &questionBank.MultipleChoice{
				CourseSupport:       courseSupport,
				MultipleChoiceModel: *model,
			}
			choices = append(choices, multipleChoice)
		}
	}
	_ = to.Create(&choices).Error
	_ = to.Create(&programs).Error
	_ = to.Create(&supplyBlanks).Error
	_ = to.Create(&judges).Error
}

// 老代码已弃用
//
//type OldProgram struct {
//	Id             string `json:"id"`
//	Title          string `json:"title"`
//	Description    string `json:"describe" gorm:"column:describe"`
//	ProblemType    int    `json:"problem_type"`
//	Stage          int    `json:"stage"`
//	Code           string `json:"code"`
//	IsProgramBlank int    `json:"is_program_blank"`
//	CanPractice    int
//	CanExam        int
//	KnowledgePoint string `json:"knowledge_point"`
//}
//
//// OldJudge 老题库的判断题
//type OldJudge struct {
//	Name   string `gorm:"column:Name"`
//	Answer string `gorm:"column:Answer"`
//	//QuestionBh   string `gorm:"QuestionBh"`
//	//CustomBh 	string `gorm:"CustomBh"`
//	Difficulty  uint   `gorm:"column:Difficulty"`
//	Stage       uint   `gorm:"column:Stage"`
//	Description string `gorm:"column:Description"`
//	KnowledgeBh string `gorm:"column:KnowledgeBh"`
//	CanPractice int    `gorm:"column:can_practice"`
//	CanExam     int    `gorm:"column:can_exam"`
//}
//
//// OldMultipleChoice 老题库的判断题
//type OldMultipleChoice struct {
//	Name   string `gorm:"column:Name"`
//	Answer string `gorm:"column:Answer"`
//	//QuestionBh   string `gorm:"QuestionBh"`
//	//CustomBh 	string `gorm:"CustomBh"`
//	Difficulty  uint   `gorm:"column:Difficulty"`
//	Stage       uint   `gorm:"column:Stage"`
//	Description string `gorm:"column:Description"`
//	KnowledgeBh string `gorm:"column:KnowledgeBh"`
//	CanPractice int    `gorm:"column:can_practice"`
//	CanExam     int    `gorm:"column:can_exam"`
//}
//
//// OldSupplyBlank 老题库的填空题
//type OldSupplyBlank struct {
//	Name string `gorm:"column:Name"`
//	//Answer int `gorm:"column:Answer"`
//	QuestionBh string `gorm:"QuestionBh"`
//	//CustomBh 	string `gorm:"CustomBh"`
//	Difficulty  uint   `gorm:"column:Difficulty"`
//	Stage       uint   `gorm:"column:Stage"`
//	Description string `gorm:"column:Description"`
//	KnowledgeBh string `gorm:"column:KnowledgeBh"`
//	CanPractice int    `gorm:"column:can_practice"`
//	CanExam     int    `gorm:"column:can_exam"`
//}
//
//type Knowledge struct {
//	KnowledgeBh   string `gorm:"column:KnowledgeBh"`
//	KnowledgeName string `gorm:"column:KnowledgeName"`
//	Description   string `gorm:"column:Description"`
//	Stage1        uint   `gorm:"column:stage1"`
//}
//
//func newPointInt(i int) *int {
//	return &i
//}
//func getKnowledgeId(knowledgeBh string, knowledgeTable map[string]uint, from, to *gorm.DB) uint {
//	if knowledgeBh == "" {
//		return 0
//	}
//	if v, ok := knowledgeTable[knowledgeBh]; ok {
//		return v
//	} else {
//		k := Knowledge{}
//		if err := from.Raw("select *,right(Stage,1) as `stage1` from knowledgepoint where KnowledgeBh=?", knowledgeBh).Find(&k).Error; err != nil {
//			return 0
//		}
//		if k.KnowledgeName == "" {
//			return 0
//		}
//		knowledge := lessondata.Knowledge{}
//		knowledge.Name = k.KnowledgeName
//		knowledge.ChapterId = uint(k.Stage1)
//		knowledge.Description = k.Description
//		if err := to.Create(&knowledge).Error; err != nil {
//			panic(err)
//		}
//		knowledgeTable[knowledgeBh] = knowledge.ID
//		return knowledge.ID
//	}
//}
//
//func transformation(from *gorm.DB, to *gorm.DB) {
//	knowledgeTable := make(map[string]uint)
//	transformationJudge(from, to, knowledgeTable)
//	transformationProgram(from, to, knowledgeTable)
//	transformationSupplyBlank(from, to, knowledgeTable)
//	transformationMultipleChoice(from, to, knowledgeTable)
//}
//
//func transformationSupplyBlank(from *gorm.DB, to *gorm.DB, knowledgeTable map[string]uint) {
//	fmt.Println("下面开始转换填空题")
//	var results []*OldSupplyBlank
//	from.Raw("SELECT `Name`,`Answer`,CASE WHEN Difficulty='1000403' THEN '2'\nWHEN Difficulty='1000402' THEN '1'\nELSE '3' END AS 'Difficulty' ,right(Stage,1) as `Stage`,`Description`,`KnowledgeBh`,Checked='100001' as 'can_exam', Score as 'can_practice',`QuestionBh` from questions where CourseID=2000301 and QuestionType=\t1000204\n").Find(&results)
//	chapterMerges := make([]*questionBank.ChapterMerge, 0, 100)
//	for i, result := range results {
//		supplyBlank := questionBank.SupplyBlank{}
//		supplyBlank.Title = result.Name
//		supplyBlank.Describe = result.Description
//		supplyBlank.CanExam = &result.CanExam
//		supplyBlank.CanPractice = &result.CanPractice
//		supplyBlank.ProblemType = int(result.Difficulty)
//		blanks := []Blank{}
//		from.Raw("select * from apfill where QuestionBh=?", result.QuestionBh).Find(&blanks)
//		var proportion []string
//		var answer []string
//		for _, blank := range blanks {
//			proportion = append(proportion, blank.Proportion)
//			answer = append(answer, blank.Answer)
//		}
//		supplyBlank.Num = len(blanks)
//		supplyBlank.Proportion = strings.Join(proportion, ",")
//		supplyBlank.Answer = strings.Join(answer, ",")
//
//		to.Create(&supplyBlank)
//		chapterMerge := &questionBank.ChapterMerge{}
//		chapterMerge.QuestionId = supplyBlank.ID
//		chapterMerge.KnowledgeId = getKnowledgeId(result.KnowledgeBh, knowledgeTable, from, to)
//		chapterMerge.ChapterId = uint(result.Stage)
//		chapterMerge.QuestionId = supplyBlank.ID
//		chapterMerge.QuestionType = questionType.SUPPLY_BLANK
//		chapterMerges = append(chapterMerges, chapterMerge)
//		fmt.Printf("处理了%d道填空题\n", i+1)
//	}
//	if err := to.CreateInBatches(chapterMerges, 50).Error; err != nil {
//		panic(err)
//	}
//}
//func transformationJudge(from *gorm.DB, to *gorm.DB, knowledgeTable map[string]uint) {
//	fmt.Println("下面开始转换判断题")
//	var results []*OldJudge
//	from.Raw("SELECT `Name`,`Answer`,CASE WHEN Difficulty='1000403' THEN '2'\nWHEN Difficulty='1000402' THEN '1'\nELSE '3' END AS 'Difficulty' ,right(Stage,1) as `Stage`,`Description` as `Description`,`KnowledgeBh`,Checked='100001' as 'can_exam', Score as 'can_practice' FROM `questions` where CourseID=2000301 and QuestionType=\t1000203").Find(&results)
//	chapterMerges := make([]*questionBank.ChapterMerge, 0, 100)
//	for i, result := range results {
//		judge := questionBank.Judge{}
//		judge.Title = result.Name
//		judge.Describe = result.Description
//		judge.CanExam = &result.CanExam
//		judge.CanPractice = &result.CanPractice
//		judge.ProblemType = int(result.Difficulty)
//		atoi, err := strconv.Atoi(strings.Trim(result.Answer, " "))
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		judge.IsRight = &atoi
//		judge.ProblemType = int(result.Difficulty)
//		if err := to.Create(&judge).Error; err != nil {
//			panic(err)
//		}
//		chapterMerge := &questionBank.ChapterMerge{}
//		chapterMerge.KnowledgeId = getKnowledgeId(result.KnowledgeBh, knowledgeTable, from, to)
//		chapterMerge.ChapterId = uint(result.Stage)
//		chapterMerge.QuestionId = judge.ID
//		chapterMerge.QuestionType = questionType.JUDGE
//		chapterMerges = append(chapterMerges, chapterMerge)
//		fmt.Printf("处理完%d个选择题了!\n", i+1)
//	}
//	if err := to.CreateInBatches(chapterMerges, 50).Error; err != nil {
//		panic(err)
//	}
//}
//
//func transformationProgram(from *gorm.DB, to *gorm.DB, knowledgeTable map[string]uint) {
//	// 编程题转化
//	fmt.Println("下面开始转换编程题")
//	results := make([]OldProgram, 0, 500)
//	sql := `select QuestionBh as id, name as 'title',Description as "describe",
//		CASE WHEN Difficulty='1000403' THEN '2' WHEN Difficulty='1000402' THEN '1' ELSE '3' END AS 'problem_type' ,
//		right(Stage,1) as stage,
//		CASE WHEN JSON_VALID(SourceCode) THEN JSON_UNQUOTE(JSON_EXTRACT(SourceCode, "$.key[0].code"))  ELSE null END as code,
//		IsProgramBlank='100001' as is_program_blank,
//		Checked='100001' as 'can_exam',
//		Score as 'can_practice' ,
//		KnowledgeBh as knowledge_point
//		from stu_system.questions
//		where CourseID=2000301 and QuestionType=1000206`
//	from.Raw(sql).Find(&results)
//	//用来记录说编程题所对应的知识点ID
//
//	cases := make([]*questionBank.ProgrammCase, 0, 2000)
//	fmt.Println("下面开始转化编程题")
//	chapterMerges := make([]*questionBank.ChapterMerge, 0, 500)
//	for i, result := range results {
//		// 创建 programm
//		programm := questionBank.Programm{}
//		programm.Title = result.Title
//		programm.Describe = result.Description
//		programm.CanExam = &result.CanExam
//		programm.CanPractice = &result.CanPractice
//		programm.ProblemType = result.ProblemType
//		if err := to.Create(&programm).Error; err != nil {
//			panic(err)
//		}
//
//		// 创建 编程题 语言支持
//		merge := questionBank.ProgrammLanguageMerge{}
//		merge.LanguageId = language.C_LANGUAGE
//		if result.IsProgramBlank == 0 {
//			merge.ReferenceAnswer = result.Code
//		} else {
//			merge.DefaultCode = result.Code
//		}
//		merge.ProgrammId = programm.ID
//		if err := to.Create(&merge).Error; err != nil {
//			panic(err)
//		}
//		var testCase []TestCase
//		if err := from.Table("testcase").Where("QuestionId", result.Id).Find(&testCase).Error; err != nil {
//			panic(err)
//		}
//
//		// 加入测试用例集合
//		for i, t := range testCase {
//			thisCase := &questionBank.ProgrammCase{}
//			thisCase.Name = fmt.Sprintf("测试用例-%d", i)
//			thisCase.Input = t.Input
//			thisCase.Output = t.Output
//			thisCase.LanguageId = language.C_LANGUAGE
//			thisCase.ProgrammId = programm.ID
//			thisCase.Score = uint(t.Score)
//			cases = append(cases, thisCase)
//		}
//
//		// 加入章节绑定
//		chapterMerge := &questionBank.ChapterMerge{}
//		chapterMerge.KnowledgeId = getKnowledgeId(result.KnowledgePoint, knowledgeTable, from, to)
//		chapterMerge.ChapterId = uint(result.Stage)
//		chapterMerge.QuestionId = programm.ID
//		chapterMerge.QuestionType = questionType.PROGRAM
//
//		chapterMerges = append(chapterMerges, chapterMerge)
//
//		fmt.Println("已经处理完", i+1, "个编程题勒！")
//	}
//
//	if err := to.CreateInBatches(cases, 50).Error; err != nil {
//		panic(err)
//	}
//	if err := to.CreateInBatches(chapterMerges, 50).Error; err != nil {
//		panic(err)
//	}
//	fmt.Println("已经处理完了", len(cases), "个编程用例")
//}
//
//func transformationMultipleChoice(from *gorm.DB, to *gorm.DB, knowledgeTable map[string]uint) {
//	fmt.Println("下面开始转换选择题")
//	var results []*OldMultipleChoice
//	from.Raw("SELECT `Name`,`Answer`,CASE WHEN Difficulty='1000403' THEN '2'\nWHEN Difficulty='1000402' THEN '1'\nELSE '3' END AS 'Difficulty' ,right(Stage,1) as `Stage`,`Description`,`KnowledgeBh`,Checked='100001' as 'can_exam', Score as 'can_practice'FROM `questions` where CourseID=2000301 and QuestionType=\t100020101").Find(&results)
//	chapterMerges := make([]*questionBank.ChapterMerge, 0, 100)
//	for i, result := range results {
//		multipleChoice := questionBank.MultipleChoice{}
//		multipleChoice.Title = result.Name
//		multipleChoice.Describe = result.Description
//		multipleChoice.CanExam = &result.CanExam
//		multipleChoice.CanPractice = &result.CanPractice
//		multipleChoice.Answer = result.Answer
//		multipleChoice.MostOptions = 1
//		multipleChoice.ProblemType = int(result.Difficulty)
//		if err := to.Create(&multipleChoice).Error; err != nil {
//			panic(err)
//		}
//		chapterMerge := &questionBank.ChapterMerge{}
//		chapterMerge.KnowledgeId = getKnowledgeId(result.KnowledgeBh, knowledgeTable, from, to)
//		chapterMerge.ChapterId = uint(result.Stage)
//		chapterMerge.QuestionId = multipleChoice.ID
//		chapterMerge.QuestionType = questionType.SINGLE_CHOICE
//		chapterMerges = append(chapterMerges, chapterMerge)
//		fmt.Printf("处理完%d个选择题了!\n", i+1)
//	}
//	if err := to.CreateInBatches(chapterMerges, 50).Error; err != nil {
//		panic(err)
//	}
//
//}
