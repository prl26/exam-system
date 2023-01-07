package examManage

import (
	"bytes"
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/utils"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
	"io"
	"strconv"
	"strings"
	"time"
)

type ExamService struct {
}

func (examService *ExamService) FindExamPlans(teachClassId uint) (examPlans []teachplan.ExamPlan, err error) {
	err = global.GVA_DB.Where("teach_class_id = ? and state = 1 and audit =1", teachClassId).Order("created_at desc,updated_at desc ").Find(&examPlans).Error
	return
}

func (examService *ExamService) GetExamPapers(examComing request.ExamComing) (examPaper response.ExamPaperResponse, status examManage.StudentPaperStatus, err error) {
	examPaper.BlankComponent = make([]response.BlankComponent, 0)
	examPaper.SingleChoiceComponent = make([]response.ChoiceComponent, 0)
	examPaper.MultiChoiceComponent = make([]response.ChoiceComponent, 0)
	examPaper.JudgeComponent = make([]response.JudgeComponent, 0)
	examPaper.ProgramComponent = make([]response.ProgramComponent, 0)
	var studentPaper []examManage.ExamStudentPaper
	err = global.GVA_DB.Where("student_id = ? and plan_id = ?", examComing.StudentId, examComing.PlanId).Find(&studentPaper).Error
	var singleChoiceCount, MultiChoiceCount, judgeCount, blankCount, programCount uint
	for i := 0; i < len(studentPaper); i++ {
		if *studentPaper[i].QuestionType == questionType.SINGLE_CHOICE {
			var Choice response.ChoiceComponent
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Where("id = ?", studentPaper[i].QuestionId).Find(&Choice.Choice).Error
			if err != nil {
				return
			}
			Choice.MergeId = studentPaper[i].ID
			if Choice.Choice.IsIndefinite == 0 {
				examPaper.SingleChoiceComponent = append(examPaper.SingleChoiceComponent, Choice)
				examPaper.SingleChoiceComponent[singleChoiceCount].MergeId = studentPaper[i].ID
				singleChoiceCount++
			} else {
				examPaper.MultiChoiceComponent = append(examPaper.MultiChoiceComponent, Choice)
				examPaper.MultiChoiceComponent[MultiChoiceCount].MergeId = studentPaper[i].ID
				MultiChoiceCount++
			}
		} else if *studentPaper[i].QuestionType == questionType.JUDGE {
			var Judge response.JudgeComponent
			err = global.GVA_DB.Table("les_questionBank_judge").Where("id = ?", studentPaper[i].QuestionId).Find(&Judge.Judge).Error
			if err != nil {
				return
			}
			examPaper.JudgeComponent = append(examPaper.JudgeComponent, Judge)
			examPaper.JudgeComponent[judgeCount].MergeId = studentPaper[i].ID
			judgeCount++
		} else if *studentPaper[i].QuestionType == questionType.SUPPLY_BLANK {
			var Blank response.BlankComponent
			err = global.GVA_DB.Table("les_questionBank_supply_blank").Where("id = ?", studentPaper[i].QuestionId).Find(&Blank.Blank).Error
			if err != nil {
				return
			}
			examPaper.BlankComponent = append(examPaper.BlankComponent, Blank)
			examPaper.BlankComponent[blankCount].MergeId = studentPaper[i].ID
			blankCount++
		} else if *studentPaper[i].QuestionType == questionType.PROGRAM {
			var Program response.ProgramComponent
			var program questionBankBo.ProgramPractice
			err = global.GVA_DB.Table("les_questionBank_programm").Where("id = ?", studentPaper[i].QuestionId).Find(&program).Error
			if err != nil {
				return
			}
			Program.Program.Convert(&program)
			examPaper.ProgramComponent = append(examPaper.ProgramComponent, Program)
			examPaper.ProgramComponent[programCount].MergeId = studentPaper[i].ID
			programCount++
		}
	}
	var PaperId int64
	err = global.GVA_DB.Table("exam_student_paper").Select("paper_id").Where("student_id = ? and plan_id =?", examComing.StudentId, examComing.PlanId).Scan(&PaperId).Error
	//PaperId, err := examService.GetStudentPaperId(examComing)
	if err != nil {
		return
	}
	examPaper.PaperId = uint(PaperId)
	status, err = examService.CreateStatus(examComing)
	fmt.Println(status)
	if err != nil {
		return
	}
	return
}
func (examService *ExamService) GetStudentPaperId(examComing request.ExamComing) (Id int64, err error) {
	err = global.GVA_DB.Table("exam_student_paper").Select("paper_id").Where("student_id = ? and plan_id =?", examComing.StudentId, examComing.PlanId).First(&Id).Error
	return
}

func (examService *ExamService) CreateStatus(examComing request.ExamComing) (status examManage.StudentPaperStatus, err error) {
	var num int64
	err = global.GVA_DB.Table("student_paper_status").Where("student_id = ? and plan_id = ?", examComing.StudentId, examComing.PlanId).Find(&status).Count(&num).Error
	if err != nil {
		return
	} else if num == 0 {
		status = examManage.StudentPaperStatus{
			GVA_MODEL: global.GVA_MODEL{},
			StudentId: examComing.StudentId,
			PlanId:    examComing.PlanId,
			EnterTime: time.Now(),
			IsCommit:  false,
		}
		global.GVA_DB.Create(&status)
	}
	return
}

func (examService *ExamService) CommitExamPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	var optionCommit = examPaperCommit.MultipleChoiceCommit
	var JudgeCommit = examPaperCommit.JudgeCommit
	var BlankCommit = examPaperCommit.BlankCommit
	for j := 0; j < len(optionCommit); j++ {
		answers := strings.Join(optionCommit[j].Answer, ",")
		err = global.GVA_DB.Table("exam_student_paper").Select("answer").
			Where("id = ?", optionCommit[j].MergeId).
			Updates(&examManage.ExamStudentPaper{Answer: answers}).
			Error
		if err != nil {
			return
		}
	}
	for j := 0; j < len(JudgeCommit); j++ {
		s := strconv.FormatBool(examPaperCommit.JudgeCommit[0].Answer)
		err = global.GVA_DB.Table("exam_student_paper").Select("answer").
			Where("id = ?", JudgeCommit[j].MergeId).
			Updates(examManage.ExamStudentPaper{Answer: s}).
			Error
		if err != nil {
			return
		}
	}
	for j := 0; j < len(BlankCommit); j++ {
		blankAnswer := utils.StringArrayToString(BlankCommit[j].Answer)
		err = global.GVA_DB.Table("exam_student_paper").Select("answer").
			Where("id = ?", BlankCommit[j].MergeId).
			Updates(examManage.ExamStudentPaper{Answer: blankAnswer}).
			Error
		if err != nil {
			return
		}
	}
	err = global.GVA_DB.Table("student_paper_status").Where("student_id = ? and plan_id =?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("is_commit", 1).Error
	if err != nil {
		return
	}
	return
}
func (examService *ExamService) CommitProgram(program examManage.CommitProgram) (err error) {
	err = global.GVA_DB.Table("exam_student_paper").Select("answer").
		Where("id = ?", program.MergeId).
		Updates(&examManage.ExamStudentPaper{Answer: program.Code}).
		Error
	return
}
func (examService *ExamService) GetExamScore(info request.ExamStudentScore, studentId uint) (studentScore []response.ExamScoreResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamScore{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if *info.TermId != 0 {
		db = db.Where("term_id = ?", info.TermId)
	}
	if *info.LessonId != 0 {
		db = db.Where("lesson_id = ?", info.LessonId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Where("student_id = ?", studentId).Order("created_at desc,updated_at desc ").Limit(limit).Offset(offset).Find(&studentScore).Error
	return studentScore, total, err
}
func (ExamService *ExamService) ExportPaperScore(infoList []teachplan.Score, filePath string) (err error) {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"学号", "课程名称", "教学班名称",
		"期末考试成绩", "期末考试占比", "过程化考核得分", "过程化考核占比",
		"考勤得分", "考勤占比", "学习资源得分", "学习资源占比"})
	for i, paper := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			paper.ID,
			paper.CourseName,
			paper.TeachClassName,
			paper.ExamScrore,
			paper.ExamProporation,
			paper.ProcedureScore,
			paper.ProcedureProportion,
			paper.AttendanceScore,
			paper.AttendanceProportion,
			paper.LearnResourcesScore,
			paper.LearnResourcesProportion,
		})
	}
	err = excel.SaveAs(filePath)
	return err
}
func (ExamService *ExamService) GetTeachScore(id uint) (infoList []teachplan.Score, err error) {
	err = global.GVA_DB.Where("teach_class_id = ?", id).Find(&infoList).Error
	return
}
func (ExamService *ExamService) ExportPaperScore1(infoList []teachplan.Score) (content io.ReadSeeker, err error) {
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleList := []string{"学号", "课程名称", "教学班名称", "考勤得分", "考勤占比", "学习资源得分", "学习资源占比",
		"过程化考核得分", "过程化考核占比", "期末考试成绩", "期末考试占比", "总分",
	}
	titleRow := sheet.AddRow()
	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
	}
	// 插入内容
	var dataList []interface{}
	for _, role := range infoList {
		dataList = append(dataList, &teachplan.Score{
			StudentId:                role.StudentId,
			CourseName:               role.CourseName,
			TeachClassName:           role.TeachClassName,
			AttendanceProportion:     role.AttendanceProportion,
			AttendanceScore:          role.AttendanceScore,
			LearnResourcesProportion: role.LearnResourcesProportion,
			LearnResourcesScore:      role.LearnResourcesScore,
			ProcedureScore:           role.ProcedureScore,
			ProcedureProportion:      role.ProcedureProportion,
			ExamScrore:               role.ExamScrore,
			ExamProporation:          role.ExamProporation,
			TotalScore:               role.TotalScore,
		})
	}
	for _, v := range dataList {
		row := sheet.AddRow()
		row.WriteStruct(v, -1)
	}
	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content = bytes.NewReader(buffer.Bytes())
	return
}
