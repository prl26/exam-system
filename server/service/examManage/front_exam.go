package examManage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/examType"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	"github.com/prl26/exam-system/server/model/teachplan"
	response2 "github.com/prl26/exam-system/server/model/teachplan/response"
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
	err = global.GVA_DB.Where("teach_class_id = ? and state = 2 and audit =2", teachClassId).Order("created_at desc,updated_at desc").Find(&examPlans).Error
	return
}
func (examService *ExamService) FindTargetExamPlans(teachClassId uint, sId uint) (planAndStatus []response2.ExamPlanRp1, err error) {
	var examPlans []teachplan.ExamPlan
	err = global.GVA_DB.Where("teach_class_id = ? and state = 2 and audit =2", teachClassId).Order("created_at desc,updated_at desc").Find(&examPlans).Error
	for i := 0; i < len(examPlans); i++ {
		var plan response2.ExamPlanRp1
		plan.Plan = examPlans[i]
		if examPlans[i].StartTime.Unix() > time.Now().Unix() {
			plan.Status.IsBegin = 0
		} else if examPlans[i].EndTime.Unix() < time.Now().Unix() {
			plan.Status.IsBegin = 2
		} else {
			plan.Status.IsBegin = 1
		}
		if commit, err := examService.GetPlanStatus(examPlans[i].ID, sId); err != nil {
			return nil, err
		} else if commit == true {
			plan.Status.IsCommit = 1
		} else {
			plan.Status.IsCommit = 0
		}
		if isFinishPreExam, err, _ := examService.IsFinishPreExam(examPlans[i].ID, sId); err != nil {
			return nil, err
		} else if isFinishPreExam == true {
			plan.Status.IsFinishPreExams = 1
		} else {
			plan.Status.IsFinishPreExams = 0
		}
		planAndStatus = append(planAndStatus, plan)
	}
	return
}
func (examService *ExamService) IsFinishPreExam(planId uint, studentId uint) (result bool, err error, preExamIds []string) {
	var examPlan teachplan.ExamPlan
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id = ?", planId).Find(&examPlan)
	preExamIds = strings.Split(examPlan.PrePlanId, ",")
	for _, v := range preExamIds {
		preExamId, _ := strconv.Atoi(v)
		if preExamId == 0 {
			continue
		}
		var examRecords examManage.ExamScore
		var count int64
		err = global.GVA_DB.Where("plan_id = ? and student_id = ?", preExamId, studentId).Find(&examRecords).Count(&count).Error
		if err != nil {
			return false, err, preExamIds
		}
		if count == 0 {
			return false, nil, preExamIds
		}
	}
	return true, err, preExamIds
}

func (examService *ExamService) GetPlanStatus(planId uint, sId uint) (isCommit bool, err error) {
	var status examManage.StudentPaperStatus
	var num int64
	err = global.GVA_DB.Table("student_paper_status").Where("student_id = ? and plan_id = ?", sId, planId).Find(&status).Count(&num).Error
	if err != nil {
		return false, err
	} else if num == 0 || status.IsCommit == false {
		return false, nil
	}
	return true, nil
}

func (examService *ExamService) GetExamPapers(examComing request.ExamComing, IP string) (examPaper response.ExamPaperResponse, status examManage.StudentPaperStatus, err error) {
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
	status, err = examService.CreateStatus(examComing, IP)
	if err != nil {
		return
	}
	var PlanDetail teachplan.ExamPlan
	global.GVA_DB.Model(teachplan.ExamPlan{}).Where("id =?", examComing.PlanId).Find(&PlanDetail)
	err = utils.CreateExamScore(PlanDetail, 0, examComing.StudentId)
	if err != nil {
		return
	}
	return
}

func (examService *ExamService) GetExamPapersAndScores(examComing request.ExamComing, IP string) (examPaper response.ExamPaperResponse2, status examManage.StudentPaperStatus, err error) {
	examPaper.BlankComponent = make([]response.BlankComponent2, 0)
	examPaper.SingleChoiceComponent = make([]response.ChoiceComponent2, 0)
	examPaper.MultiChoiceComponent = make([]response.ChoiceComponent2, 0)
	examPaper.JudgeComponent = make([]response.JudgeComponent2, 0)
	examPaper.ProgramComponent = make([]response.ProgramComponent2, 0)
	var studentPaper []examManage.ExamStudentPaper
	err = global.GVA_DB.Where("student_id = ? and plan_id = ?", examComing.StudentId, examComing.PlanId).Find(&studentPaper).Error
	var singleChoiceCount, MultiChoiceCount, judgeCount, blankCount, programCount uint
	for i := 0; i < len(studentPaper); i++ {
		if *studentPaper[i].QuestionType == questionType.SINGLE_CHOICE {
			var Choice response.ChoiceComponent2
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Where("id = ?", studentPaper[i].QuestionId).Find(&Choice.Choice).Error
			if err != nil {
				return
			}
			//Choice.MergeId = studentPaper[i].ID
			if Choice.Choice.IsIndefinite == 0 {
				examPaper.SingleChoiceComponent = append(examPaper.SingleChoiceComponent, Choice)
				examPaper.SingleChoiceComponent[singleChoiceCount].MergeId = studentPaper[i].ID
				examPaper.SingleChoiceComponent[singleChoiceCount].Score = studentPaper[i].Score
				examPaper.SingleChoiceComponent[singleChoiceCount].Answer = studentPaper[i].Answer
				examPaper.SingleChoiceComponent[singleChoiceCount].GotScore = studentPaper[i].GotScore
				singleChoiceCount++
			} else {
				examPaper.MultiChoiceComponent = append(examPaper.MultiChoiceComponent, Choice)
				examPaper.MultiChoiceComponent[MultiChoiceCount].MergeId = studentPaper[i].ID
				MultiChoiceCount++
			}
		} else if *studentPaper[i].QuestionType == questionType.JUDGE {
			var Judge response.JudgeComponent2
			err = global.GVA_DB.Table("les_questionBank_judge").Where("id = ?", studentPaper[i].QuestionId).Find(&Judge.Judge).Error
			if err != nil {
				return
			}
			examPaper.JudgeComponent = append(examPaper.JudgeComponent, Judge)
			examPaper.JudgeComponent[judgeCount].MergeId = studentPaper[i].ID
			examPaper.JudgeComponent[judgeCount].Score = studentPaper[i].Score
			examPaper.JudgeComponent[judgeCount].GotScore = studentPaper[i].GotScore
			examPaper.JudgeComponent[judgeCount].Answer = studentPaper[i].Answer
			judgeCount++
		} else if *studentPaper[i].QuestionType == questionType.SUPPLY_BLANK {
			var Blank response.BlankComponent2
			err = global.GVA_DB.Table("les_questionBank_supply_blank").Where("id = ?", studentPaper[i].QuestionId).Find(&Blank.Blank).Error
			if err != nil {
				return
			}
			examPaper.BlankComponent = append(examPaper.BlankComponent, Blank)
			examPaper.BlankComponent[blankCount].MergeId = studentPaper[i].ID
			examPaper.BlankComponent[blankCount].Score = studentPaper[i].Score
			examPaper.BlankComponent[blankCount].GotScore = studentPaper[i].GotScore
			examPaper.BlankComponent[blankCount].Answer = studentPaper[i].Answer
			blankCount++
		} else if *studentPaper[i].QuestionType == questionType.PROGRAM {
			var Program response.ProgramComponent2
			var program questionBankBo.ProgramPractice
			err = global.GVA_DB.Table("les_questionBank_programm").Where("id = ?", studentPaper[i].QuestionId).Find(&program).Error
			if err != nil {
				return
			}
			Program.Program.Convert(&program)
			examPaper.ProgramComponent = append(examPaper.ProgramComponent, Program)
			examPaper.ProgramComponent[programCount].MergeId = studentPaper[i].ID
			examPaper.ProgramComponent[programCount].Score = studentPaper[i].Score
			examPaper.ProgramComponent[programCount].Answer = studentPaper[i].Answer
			examPaper.ProgramComponent[programCount].GotScore = studentPaper[i].GotScore
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
	status, err = examService.CreateStatus(examComing, IP)
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

func (examService *ExamService) CreateStatus(examComing request.ExamComing, IP string) (status examManage.StudentPaperStatus, err error) {
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
			EndTime:   time.Now(),
			IsCommit:  false,
			Ip:        IP,
		}
		global.GVA_DB.Create(&status)
	}
	return
}

//保存试卷
func (examService *ExamService) SaveExamPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	var optionCommit = examPaperCommit.MultipleChoiceCommit
	var JudgeCommit = examPaperCommit.JudgeCommit
	var BlankCommit = examPaperCommit.BlankCommit
	for j := 0; j < len(optionCommit); j++ {
		answers := strings.Join(optionCommit[j].Answer, ",")
		global.GVA_REDIS.Set(context.Background(), fmt.Sprintf("examRecord:%d:%d:%d", examPaperCommit.StudentId, examPaperCommit.PlanId, optionCommit[j].MergeId), answers, 7*24*time.Hour)
	}
	for j := 0; j < len(JudgeCommit); j++ {
		s := strconv.FormatBool(examPaperCommit.JudgeCommit[0].Answer)
		global.GVA_REDIS.Set(context.Background(), fmt.Sprintf("examRecord:%d:%d:%d", examPaperCommit.StudentId, examPaperCommit.PlanId, JudgeCommit[j].MergeId), s, 7*24*time.Hour)
	}
	for j := 0; j < len(BlankCommit); j++ {
		blankAnswer := utils.StringArrayToString(BlankCommit[j].Answer)
		global.GVA_REDIS.Set(context.Background(), fmt.Sprintf("examRecord:%d:%d:%d", examPaperCommit.StudentId, examPaperCommit.PlanId, BlankCommit[j].MergeId), blankAnswer, 7*24*time.Hour)
	}
	return
}

//提交试卷
func (examService *ExamService) CommitExamPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	var optionCommit = examPaperCommit.MultipleChoiceCommit
	var JudgeCommit = examPaperCommit.JudgeCommit
	var BlankCommit = examPaperCommit.BlankCommit
	//err = global.GVA_DB.Table("student_paper_status").Where("student_id = ? and plan_id =?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("is_commit", 1).Error
	err = global.GVA_DB.Model(examManage.StudentPaperStatus{}).Where("student_id = ? and plan_id =?", examPaperCommit.StudentId, examPaperCommit.PlanId).Updates(examManage.StudentPaperStatus{IsCommit: true, EndTime: time.Now()}).Error
	if err != nil {
		return
	}
	for j := 0; j < len(optionCommit); j++ {
		answers := strings.Join(optionCommit[j].Answer, ",")
		global.GVA_REDIS.Set(context.Background(), fmt.Sprintf("examRecord:%d:%d:%d", examPaperCommit.StudentId, examPaperCommit.PlanId, optionCommit[j].MergeId), answers, 7*24*time.Hour)
	}
	for j := 0; j < len(JudgeCommit); j++ {
		s := strconv.FormatBool(examPaperCommit.JudgeCommit[0].Answer)
		global.GVA_REDIS.Set(context.Background(), fmt.Sprintf("examRecord:%d:%d:%d", examPaperCommit.StudentId, examPaperCommit.PlanId, JudgeCommit[j].MergeId), s, 7*24*time.Hour)
	}
	for j := 0; j < len(BlankCommit); j++ {
		blankAnswer := utils.StringArrayToString(BlankCommit[j].Answer)
		global.GVA_REDIS.Set(context.Background(), fmt.Sprintf("examRecord:%d:%d:%d", examPaperCommit.StudentId, examPaperCommit.PlanId, BlankCommit[j].MergeId), blankAnswer, 7*24*time.Hour)
	}
	return
}

func (examService *ExamService) QueryExamPapers(studentId uint, planId uint, mergeId uint) (string, bool) {
	answer, err := global.GVA_REDIS.Get(context.Background(), fmt.Sprintf("examRecord:%d:%d:%d", studentId, planId, mergeId)).Result()
	if err != nil {
		return "", false
	}
	return answer, true
}
func (examService *ExamService) UpdateExamPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	var optionCommit = examPaperCommit.MultipleChoiceCommit
	var JudgeCommit = examPaperCommit.JudgeCommit
	var BlankCommit = examPaperCommit.BlankCommit
	for j := 0; j < len(optionCommit); j++ {
		answers, isCommit := examService.QueryExamPapers(examPaperCommit.StudentId, examPaperCommit.PlanId, optionCommit[j].MergeId)
		if isCommit == false {
			return errors.New("找不到答题记录")
		} else {
			err = global.GVA_DB.Table("exam_student_paper").Select("answer").
				Where("id = ?", optionCommit[j].MergeId).
				Updates(&examManage.ExamStudentPaper{Answer: answers}).
				Error
			if err != nil {
				return
			}
		}
	}
	for j := 0; j < len(JudgeCommit); j++ {
		answers, isCommit := examService.QueryExamPapers(examPaperCommit.StudentId, examPaperCommit.PlanId, JudgeCommit[j].MergeId)
		if isCommit == false {
			return errors.New("找不到答题记录")
		} else {
			err = global.GVA_DB.Table("exam_student_paper").Select("answer").
				Where("id = ?", JudgeCommit[j].MergeId).
				Updates(&examManage.ExamStudentPaper{Answer: answers}).
				Error
			if err != nil {
				return
			}
		}
	}
	for j := 0; j < len(BlankCommit); j++ {
		answers, isCommit := examService.QueryExamPapers(examPaperCommit.StudentId, examPaperCommit.PlanId, BlankCommit[j].MergeId)
		if isCommit == false {
			return errors.New("找不到答题记录")
		} else {
			err = global.GVA_DB.Table("exam_student_paper").Select("answer").
				Where("id = ?", BlankCommit[j].MergeId).
				Updates(&examManage.ExamStudentPaper{Answer: answers}).
				Error
			if err != nil {
				return
			}
		}
	}
	return
}

//已废弃
func (examService *ExamService) CommitExamPapers1(examPaperCommit examManage.CommitExamPaper) (err error) {
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
	err = db.Where("student_id = ? and is_report = 1", studentId).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Where("student_id = ? and is_report = 1", studentId).Order("created_at desc,updated_at desc ").Limit(limit).Offset(offset).Find(&studentScore).Error
	return studentScore, total, err
}
func (ExamService *ExamService) ExportScore(infoList []teachplan.Score, filePath string) (err error) {
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
func (ExamService *ExamService) ExportPaperScore(infoList []examManage.ExamScore, filePath string) (err error) {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"学号", "姓名", "考试名称", "学期", "课程名", "分数"})
	for i, paper := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		studentId := strconv.Itoa(int(*paper.StudentId))
		score := strconv.Itoa(int(*paper.Score))
		var studentInfo basicdata.Student
		global.GVA_DB.Model(basicdata.Student{}).Where("id = ?", paper.StudentId).Find(&studentInfo)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			studentId,
			studentInfo.Name,
			paper.Name,
			paper.TermName,
			paper.CourseName,
			score,
		})
	}
	err = excel.SaveAs(filePath)
	return err
}
func (ExamService *ExamService) GetMultiExamScoreToExcel(id uint) (infoList []examManage.ExamScore, err error) {
	err = global.GVA_DB.Where("plan_id = ?", id).Find(&infoList).Error
	return
}
func (ExamService *ExamService) GetPlanList(id uint) (infoList []uint, err error) {
	err = global.GVA_DB.Model(&teachplan.ExamPlan{}).Select("id").Where("teach_class_id = ?", id).Find(&infoList).Error
	return
}
func (ExamService *ExamService) ExportMultiPaperScore(planList []uint, filePath string) (err error) {
	excel := excelize.NewFile()
	//构建excel
	var s1 = []string{"学号", "姓名"}
	for k, v := range planList {
		var plan teachplan.ExamPlan
		global.GVA_DB.Where("id = ?", v).Find(&plan)
		if plan.Type == examType.ProceduralExam {
			s1 = append(s1, fmt.Sprintf("第%d次平时考试成绩", k), fmt.Sprintf("第%d次平时考试占比", k))
		}
	}
	s1 = append(s1, "平时成绩分数", "期末考试分数")
	// 获取学生列表
	var studentList []basicdata.Student
	var plandetail []teachplan.ExamPlan
	global.GVA_DB.Where("id in (?)", planList).Find(&plandetail)
	global.GVA_DB.Where("id in (?)", global.GVA_DB.Table("bas_student_teach_classes").Where("teach_class_id = ?", plandetail[0].TeachClassId)).Find(&studentList)
	//获取学生成绩
	var scoreList [][]interface{}
	for i := 0; i < len(studentList); i++ {
		var list1 = make([]interface{}, 20)
		list1 = append(list1, studentList[i].ID, studentList[i].Name)
		var infoList []examManage.ExamScore
		global.GVA_DB.Where("id in ?", planList).Order("exam_type DESC,start_time").Find(&infoList)
		var sum float64
		for j := 0; j < len(infoList); j++ {
			if *infoList[j].ExamType == examType.ProceduralExam {
				list1 = append(list1, infoList[j].Score, infoList[j].Weight)
				temp1 := *infoList[j].Score
				temp2 := float64(*infoList[j].Weight) / 100.0
				sum += temp2 * temp1
			}
			if *infoList[j].ExamType == examType.FinalExam {
				list1 = append(list1, fmt.Sprintf("%2.f", sum))
				list1 = append(list1, infoList[j].Score)
			}
		}
		scoreList = append(scoreList, list1)
	}
	excel.SetSheetRow("Sheet1", "A1", &s1)
	for i := 0; i < len(scoreList); i++ {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, scoreList[i])
	}
	err = excel.SaveAs(filePath)
	return err
}
func (ExamService *ExamService) GetTeachScore(id uint) (infoList []teachplan.Score, err error) {
	err = global.GVA_DB.Where("teach_class_id = ?", id).Find(&infoList).Error
	return
}
func (ExamService *ExamService) GetExamScoreToExcel(id uint) (infoList []examManage.ExamScore, err error) {
	err = global.GVA_DB.Where("plan_id = ?", id).Find(&infoList).Error
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
