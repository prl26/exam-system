package main

import (
	"encoding/json"
	"fmt"
	"github.com/prl26/exam-system/server/model/basicdata"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"strings"
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

type OldProgram struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"describe" gorm:"column:describe"`
	ProblemType    int    `json:"problem_type"`
	Stage          int    `json:"stage"`
	Code           string `json:"code"`
	IsProgramBlank int    `json:"is_program_blank"`
	CanPractice    int
	CanExam        int
	KnowledgePoint string `json:"knowledge_point"`
}

// OldJudge 老题库的判断题
type OldJudge struct {
	Name   string `gorm:"column:Name"`
	Answer string `gorm:"column:Answer"`
	//QuestionBh   string `gorm:"QuestionBh"`
	//CustomBh 	string `gorm:"CustomBh"`
	Difficulty  uint   `gorm:"column:Difficulty"`
	Stage       uint   `gorm:"column:Stage"`
	Description string `gorm:"column:Description"`
	KnowledgeBh string `gorm:"column:KnowledgeBh"`
	CanPractice int    `gorm:"column:can_practice"`
	CanExam     int    `gorm:"column:can_exam"`
}

// OldMultipleChoice 老题库的判断题
type OldMultipleChoice struct {
	Name   string `gorm:"column:Name"`
	Answer string `gorm:"column:Answer"`
	//QuestionBh   string `gorm:"QuestionBh"`
	//CustomBh 	string `gorm:"CustomBh"`
	Difficulty  uint   `gorm:"column:Difficulty"`
	Stage       uint   `gorm:"column:Stage"`
	Description string `gorm:"column:Description"`
	KnowledgeBh string `gorm:"column:KnowledgeBh"`
	CanPractice int    `gorm:"column:can_practice"`
	CanExam     int    `gorm:"column:can_exam"`
}

// OldSupplyBlank 老题库的填空题
type OldSupplyBlank struct {
	Name string `gorm:"column:Name"`
	//Answer int `gorm:"column:Answer"`
	QuestionBh string `gorm:"QuestionBh"`
	//CustomBh 	string `gorm:"CustomBh"`
	Difficulty  uint   `gorm:"column:Difficulty"`
	Stage       uint   `gorm:"column:Stage"`
	Description string `gorm:"column:Description"`
	KnowledgeBh string `gorm:"column:KnowledgeBh"`
	CanPractice int    `gorm:"column:can_practice"`
	CanExam     int    `gorm:"column:can_exam"`
}
type TestCase struct {
	Input  string `json:"input" gorm:"column:TestCaseInput"`
	Output string `json:"output" gorm:"column:TestCaseOutput"`
	Score  int    `json:"score" gorm:"column:ScoreWeight"`
}
type Knowledge struct {
	KnowledgeBh   string `gorm:"column:KnowledgeBh"`
	KnowledgeName string `gorm:"column:KnowledgeName"`
	Description   string `gorm:"column:Description"`
	Stage1        uint   `gorm:"column:stage1"`
}

type Blank struct {
	Answer     string
	Proportion string
}

func main() {
	from := getDB("root:cuitexamloopmysql123@tcp(exam.cuit.edu.cn:3306)/stu_system?charset=utf8")
	to := getDB("root:cuit@123456@tcp(139.9.249.149:3306)/gva?charset=utf8")
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
func getKnowledgeId(knowledgeBh string, knowledgeTable map[string]uint, from, to *gorm.DB) uint {
	if knowledgeBh == "" {
		return 0
	}
	if v, ok := knowledgeTable[knowledgeBh]; ok {
		return v
	} else {
		k := Knowledge{}
		if err := from.Raw("select *,right(Stage,1) as `stage1` from knowledgepoint where KnowledgeBh=?", knowledgeBh).Find(&k).Error; err != nil {
			//fmt.Println(err)
			//panic(err)
			return 0
		}
		if k.KnowledgeName == "" {
			return 0
		}
		knowledge := basicdata.Knowledge{}
		knowledge.Name = k.KnowledgeName
		knowledge.ChapterId = uint(k.Stage1)
		knowledge.Description = k.Description
		if err := to.Create(&knowledge).Error; err != nil {
			panic(err)
		}
		knowledgeTable[knowledgeBh] = knowledge.Id
		return knowledge.Id
	}
}

func transformation(from *gorm.DB, to *gorm.DB) {
	knowledgeTable := make(map[string]uint)
	transformationProgram(from, to, knowledgeTable)
	transformationJudge(from, to, knowledgeTable)
	transformationSupplyBlank(from, to, knowledgeTable)
	transformationMultipleChoice(from, to, knowledgeTable)
}

func transformationSupplyBlank(from *gorm.DB, to *gorm.DB, knowledgeTable map[string]uint) {
	fmt.Println("下面开始转换填空题")
	var results []*OldSupplyBlank
	from.Raw("SELECT `Name`,`Answer`,CASE WHEN Difficulty='1000403' THEN '2'\nWHEN Difficulty='1000402' THEN '1'\nELSE '3' END AS 'Difficulty' ,right(Stage,1) as `Stage`,`Description`,`KnowledgeBh`,Checked='100001' as 'can_exam', Score as 'can_practice',`QuestionBh` from questions where CourseID=2000301 and QuestionType=\t1000204\n").Find(&results)
	for i, result := range results {
		supplyBlank := po.SupplyBlank{}
		supplyBlank.Title = result.Name
		supplyBlank.Describe = result.Description
		supplyBlank.CanExam = &result.CanExam
		supplyBlank.CanPractice = &result.CanPractice
		supplyBlank.ProblemType = int(result.Difficulty)
		blanks := []Blank{}
		from.Raw("select * from apfill where QuestionBh=?", result.QuestionBh).Find(&blanks)
		var proportion []string
		var answer []string
		for _, blank := range blanks {
			proportion = append(proportion, blank.Proportion)
			answer = append(answer, blank.Answer)
		}
		supplyBlank.Num = len(blanks)
		supplyBlank.Proportion = strings.Join(proportion, ",")
		supplyBlank.Answer = strings.Join(answer, ",")
		supplyBlank.ChapterId = result.Stage
		supplyBlank.KnowledgeId = getKnowledgeId(result.KnowledgeBh, knowledgeTable, from, to)
		to.Create(&supplyBlank)
		fmt.Printf("处理了%d道填空题\n", i+1)
	}

}
func transformationJudge(from *gorm.DB, to *gorm.DB, knowledgeTable map[string]uint) {
	fmt.Println("下面开始转换判断题")
	var results []*OldJudge
	from.Raw("SELECT `Name`,`Answer`,CASE WHEN Difficulty='1000403' THEN '2'\nWHEN Difficulty='1000402' THEN '1'\nELSE '3' END AS 'Difficulty' ,right(Stage,1) as `Stage`,`Description` as `Description`,`KnowledgeBh`,Checked='100001' as 'can_exam', Score as 'can_practice' FROM `questions` where CourseID=2000301 and QuestionType=\t1000203").Find(&results)
	for i, result := range results {
		judge := po.Judge{}
		judge.Title = result.Name
		judge.Describe = result.Description
		judge.CanExam = &result.CanExam
		judge.CanPractice = &result.CanPractice
		judge.ProblemType = int(result.Difficulty)
		atoi, err := strconv.Atoi(strings.Trim(result.Answer, " "))
		if err != nil {
			fmt.Println(err)
			continue
		}
		judge.IsRight = atoi > 0
		judge.ProblemType = int(result.Difficulty)
		judge.ChapterId = result.Stage
		judge.KnowledgeId = getKnowledgeId(result.KnowledgeBh, knowledgeTable, from, to)
		if err := to.Create(&judge).Error; err != nil {
			panic(err)
		}
		fmt.Printf("处理完%d个判断题了!\n", i+1)
	}
}

func transformationProgram(from *gorm.DB, to *gorm.DB, knowledgeTable map[string]uint) {
	// 编程题转化
	fmt.Println("下面开始转换编程题")
	results := make([]OldProgram, 0, 500)
	from.Raw("select `QuestionBh` as \"id\"," +
		"`name` as 'title'," +
		"`Description` as `describe`," +
		"CASE WHEN Difficulty='1000403' THEN '2'\nWHEN Difficulty='1000402' THEN '1'\nELSE '3' END AS 'problem_type' , " +
		"right(Stage,1) as `stage`," +
		"CASE WHEN JSON_VALID(SourceCode) THEN JSON_UNQUOTE(JSON_EXTRACT(SourceCode, \"$.key[0].code\"))  ELSE null END as `code`," +
		"IsProgramBlank='100001' as `is_program_blank`," +
		"Checked='100001' as 'can_exam', " +
		"Score as 'can_practice' ," +
		"KnowledgeBh as `knowledge_point`" +
		" from stu_system.questions  " +
		" where CourseID=2000301 and QuestionType=1000206").Find(&results)
	//用来记录说编程题所对应的知识点ID

	fmt.Println("下面开始转化编程题")

	for i, result := range results {
		// 创建 programm
		programm := po.Program{}
		programm.Title = result.Title
		programm.Describe = result.Description
		programm.CanExam = &result.CanExam
		programm.CanPractice = &result.CanPractice
		programm.ProblemType = result.ProblemType
		if result.Code != "" {
			marshal, err := json.Marshal(map[string]string{
				"c": result.Code,
			})
			if err != nil {
				panic(err)
				return
			}
			if result.IsProgramBlank == 0 {
				programm.ReferenceAnswers = string(marshal)
			} else {
				programm.DefaultCodes = string(marshal)
			}
		}
		supports := questionBankBo.LanguageSupports{}
		supports = append(supports, &questionBankBo.LanguageSupport{LanguageId: 1})
		s, brief, err := supports.Serialize()
		programm.LanguageSupportsBrief = brief
		if err != nil {
			panic(err)
			return
		}
		programm.LanguageSupports = s
		var testCase []TestCase
		if err := from.Table("testcase").Where("QuestionId", result.Id).Find(&testCase).Error; err != nil {
			panic(err)
		}
		cases := questionBankBo.ProgramCases{}
		// 加入测试用例集合
		for i, t := range testCase {
			thisCase := questionBankBo.ProgramCase{}
			thisCase.Name = fmt.Sprintf("测试用例-%d", i)
			thisCase.Input = t.Input
			thisCase.Output = t.Output
			thisCase.Score = uint(t.Score)
			cases = append(cases, &thisCase)
		}
		serialize, err := cases.Serialize()
		if err != nil {
			return
		}
		programm.ProgramCases = serialize
		programm.ChapterId = uint(result.Stage) + 1
		programm.KnowledgeId = getKnowledgeId(result.KnowledgePoint, knowledgeTable, from, to)
		fmt.Println("已经处理完", i+1, "个编程题勒！")
		if err := to.Create(&programm).Error; err != nil {
			panic(err)
		}
		if err := to.Model(&po.PublicProgram{}).Create(&programm).Error; err != nil {
			panic(err)
		}
	}

}

func transformationMultipleChoice(from *gorm.DB, to *gorm.DB, knowledgeTable map[string]uint) {
	fmt.Println("下面开始转换选择题")
	var results []*OldMultipleChoice
	from.Raw("SELECT `Name`,`Answer`,CASE WHEN Difficulty='1000403' THEN '2'\nWHEN Difficulty='1000402' THEN '1'\nELSE '3' END AS 'Difficulty' ,right(Stage,1) as `Stage`,`Description`,`KnowledgeBh`,Checked='100001' as 'can_exam', Score as 'can_practice'FROM `questions` where CourseID=2000301 and QuestionType=\t100020101").Find(&results)

	for i, result := range results {
		multipleChoice := po.MultipleChoice{}
		multipleChoice.Title = result.Name
		multipleChoice.Describe = result.Description
		multipleChoice.CanExam = &result.CanExam
		multipleChoice.CanPractice = &result.CanPractice
		multipleChoice.Answer = result.Answer
		multipleChoice.MostOptions = 1
		multipleChoice.ProblemType = int(result.Difficulty)
		multipleChoice.KnowledgeId = getKnowledgeId(result.KnowledgeBh, knowledgeTable, from, to)
		multipleChoice.ChapterId = uint(result.Stage)
		if err := to.Create(&multipleChoice).Error; err != nil {
			panic(err)
		}
		fmt.Printf("处理完%d个选择题了!\n", i+1)
	}
}
