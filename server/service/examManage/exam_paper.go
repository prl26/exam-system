package examManage

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"

	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/utils"
	"gorm.io/gorm"
	"math/rand"
	"sync"
	"time"
)

type ExamPaperService struct {
}

var wg sync.WaitGroup

//func (examPaperService *ExamPaperService) FindPlanDetail(examPaper examManage.ExamPaperDraft) (examPlan teachplan.ExamPlan, err error, count int64) {
//	err = global.GVA_DB.Where("id = ?", examPaper.PlanId).Find(&examPlan).QuestionCount(&count).Error
//	if err != nil {
//		return
//	} else if count == 0 {
//		return examPlan, err, count
//	}
//	return
//}

// CreateExamPaper 创建ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) CreateExamPaper(examPaper examManage.ExamPaper) (err error) {
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var examPlan teachplan.ExamPlan
		err = global.GVA_DB.Where("id = ?", examPaper.PlanId).Find(&examPlan).Error
		examPaper.TermId = *examPlan.TermId
		lessonId := *examPlan.LessonId
		examPaper.LessonId = uint(lessonId)
		tId := int(*examPlan.TemplateId)
		examPaper.TemplateId = &tId
		tx.Create(&examPaper)
		templateItems, err := examPaperService.GetTemplate(examPaper)
		if err != nil {
			return err
		}
		if err := examPaperService.SetPaperQuestion(templateItems, examPaper.ID); err != nil {
			return err
		}
		return err
	})
	return nil
}

func (examPaperService *ExamPaperService) FindTemplateId(examPaper examManage.ExamPaper) (tId int64, err error) {
	err = global.GVA_DB.Model(teachplan.ExamPlan{}).Select("template_id").Where("id = ?", examPaper.PlanId).Scan(&tId).Error
	return
}

// DeleteExamPaper 删除ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) DeleteExamPaper(examPaper examManage.ExamPaper) (err error) {
	err = global.GVA_DB.Where("id = ?", examPaper.ID).Delete(&examPaper).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Delete(&examManage.PaperQuestionMerge{}, "paper_id = ?", examPaper.ID).Error
	return err
}
func (examPaperService *ExamPaperService) CheckPaperIsUsed(id uint) (count int64) {
	global.GVA_DB.Model(examManage.ExamStudentPaper{}).Where("paper_id = ?", id).Count(&count)
	return
}

// DeleteExamPaperByIds 批量删除ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) DeleteExamPaperByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.ExamPaper{}, "id in ?", ids.Ids).Error
	err = global.GVA_DB.Delete(&examManage.PaperQuestionMerge{}, "paper_id in ?", ids.Ids).Error
	return err
}

// UpdateExamPaper 更新ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) UpdateExamPaper(examPaper examManage.ExamPaper) (Bool bool, err error) {
	var examPaper1 examManage.ExamPaper
	err = global.GVA_DB.Where("id = ?", examPaper.ID).Find(&examPaper1).Error
	err = global.GVA_DB.Where("id = ?", examPaper.ID).Updates(&examPaper).Error
	if examPaper1.TemplateId != examPaper.TemplateId {
		Bool = true
	}
	return
}
func (examPaperService *ExamPaperService) DeletePaperMerge(examPaper examManage.ExamPaper) error {
	var PaperMerge []examManage.PaperQuestionMerge
	err := global.GVA_DB.Table("exam_paper_question_merge").Where("paper_id = ?", examPaper.ID).Delete(&PaperMerge).Error
	return err
}

// GetExamPaper 根据id获取ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) GetExamPaper(id uint) (examPaper response.ExamPaperResponse, PaperTitle examManage.ExamPaper, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&PaperTitle).Error
	examPaper.TargetComponent = make([]response.TargetComponent, 0)
	examPaper.BlankComponent = make([]response.BlankComponent, 0)
	examPaper.SingleChoiceComponent = make([]response.ChoiceComponent, 0)
	examPaper.MultiChoiceComponent = make([]response.ChoiceComponent, 0)
	examPaper.JudgeComponent = make([]response.JudgeComponent, 0)
	examPaper.ProgramComponent = make([]response.ProgramComponent, 0)
	var Paper []examManage.PaperQuestionMerge
	err = global.GVA_DB.Table("exam_paper_question_merge").Where("paper_id = ?", id).Find(&Paper).Error
	var singleChoiceCount, MultiChoiceCount, judgeCount, blankCount, programCount, targetCount uint
	for i := 0; i < len(Paper); i++ {
		if *Paper[i].QuestionType == questionType.SINGLE_CHOICE {
			var Choice response.ChoiceComponent
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Where("id = ?", Paper[i].QuestionId).Find(&Choice.Choice).Error
			if err != nil {
				return
			}
			Choice.MergeId = Paper[i].ID
			if Choice.Choice.IsIndefinite == 0 {
				examPaper.SingleChoiceComponent = append(examPaper.SingleChoiceComponent, Choice)
				examPaper.SingleChoiceComponent[singleChoiceCount].MergeId = Paper[i].ID
				singleChoiceCount++
			} else {
				examPaper.MultiChoiceComponent = append(examPaper.MultiChoiceComponent, Choice)
				examPaper.MultiChoiceComponent[MultiChoiceCount].MergeId = Paper[i].ID
				MultiChoiceCount++
			}
		} else if *Paper[i].QuestionType == questionType.JUDGE {
			var Judge response.JudgeComponent
			err = global.GVA_DB.Table("les_questionBank_judge").Where("id = ?", Paper[i].QuestionId).Find(&Judge.Judge).Error
			if err != nil {
				return
			}
			examPaper.JudgeComponent = append(examPaper.JudgeComponent, Judge)
			examPaper.JudgeComponent[judgeCount].MergeId = Paper[i].ID
			judgeCount++
		} else if *Paper[i].QuestionType == questionType.SUPPLY_BLANK {
			var Blank response.BlankComponent
			err = global.GVA_DB.Table("les_questionBank_supply_blank").Where("id = ?", Paper[i].QuestionId).Find(&Blank.Blank).Error
			if err != nil {
				return
			}
			examPaper.BlankComponent = append(examPaper.BlankComponent, Blank)
			examPaper.BlankComponent[blankCount].MergeId = Paper[i].ID
			blankCount++
		} else if *Paper[i].QuestionType == questionType.PROGRAM {
			var Program response.ProgramComponent
			var program questionBankBo.ProgramPractice
			err = global.GVA_DB.Table("les_questionBank_programm").Where("id = ?", Paper[i].QuestionId).Find(&program).Error
			if err != nil {
				return
			}
			Program.Program.Convert(&program)
			examPaper.ProgramComponent = append(examPaper.ProgramComponent, Program)
			examPaper.ProgramComponent[programCount].MergeId = Paper[i].ID
			programCount++
		} else if *Paper[i].QuestionType == questionType.Target {
			var Target response.TargetComponent
			err = global.GVA_DB.Table("les_questionBank_target").Where("id = ?", Paper[i].QuestionId).Find(&Target.Target).Error
			if err != nil {
				return
			}
			examPaper.TargetComponent = append(examPaper.TargetComponent, Target)
			examPaper.TargetComponent[targetCount].MergeId = Paper[i].ID
			targetCount++
		}
	}
	examPaper.PaperId = id
	return
}

func (examPaperService *ExamPaperService) GetExamPaper1(id uint) (examPaper response.ExamPaperResponse2, PaperTitle examManage.ExamPaper, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&PaperTitle).Error
	examPaper.TargetComponent = make([]response.STargetComponent, 0)
	examPaper.BlankComponent = make([]response.BlankComponent2, 0)
	examPaper.SingleChoiceComponent = make([]response.ChoiceComponent2, 0)
	examPaper.MultiChoiceComponent = make([]response.ChoiceComponent2, 0)
	examPaper.JudgeComponent = make([]response.JudgeComponent2, 0)
	examPaper.ProgramComponent = make([]response.ProgramComponent2, 0)
	var studentPaper []examManage.PaperQuestionMerge
	err = global.GVA_DB.Table("exam_paper_question_merge").Where("paper_id = ?", id).Find(&studentPaper).Error
	var singleChoiceCount, MultiChoiceCount, judgeCount, blankCount, programCount, targetCount uint
	var singleChoiceOrder, MultiChoiceOrder, judgeOrder, blankOrder, programOrder, targetOrder uint
	for i := 0; i < len(studentPaper); i++ {
		if *studentPaper[i].QuestionType == questionType.SINGLE_CHOICE {
			var Choice response.ChoiceComponent2
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Where("id = ?", studentPaper[i].QuestionId).Find(&Choice.Choice).Error
			var answer string
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Select("answer").Where("id = ?", studentPaper[i].QuestionId).Scan(&answer).Error
			if err != nil {
				return
			}
			//Choice.MergeId = studentPaper[i].ID
			if Choice.Choice.IsIndefinite == 0 {
				singleChoiceOrder++
				examPaper.SingleChoiceComponent = append(examPaper.SingleChoiceComponent, Choice)
				examPaper.SingleChoiceComponent[singleChoiceCount].MergeId = studentPaper[i].ID
				examPaper.SingleChoiceComponent[singleChoiceCount].Order = fmt.Sprintf("%d.", singleChoiceOrder)
				var a float64
				a = float64(*studentPaper[i].Score)
				examPaper.SingleChoiceComponent[singleChoiceCount].Score = &a
				examPaper.SingleChoiceComponent[singleChoiceCount].CorrectAnswer = answer
				singleChoiceCount++
			} else {
				MultiChoiceOrder++
				examPaper.MultiChoiceComponent = append(examPaper.MultiChoiceComponent, Choice)
				examPaper.MultiChoiceComponent[MultiChoiceCount].MergeId = studentPaper[i].ID
				examPaper.MultiChoiceComponent[MultiChoiceCount].Order = fmt.Sprintf("%d.", MultiChoiceOrder)
				var a float64
				a = float64(*studentPaper[i].Score)
				examPaper.MultiChoiceComponent[MultiChoiceCount].Score = &a
				examPaper.MultiChoiceComponent[MultiChoiceCount].CorrectAnswer = answer
				MultiChoiceCount++
			}
		} else if *studentPaper[i].QuestionType == questionType.JUDGE {
			var Judge response.JudgeComponent2
			err = global.GVA_DB.Table("les_questionBank_judge").Where("id = ?", studentPaper[i].QuestionId).Find(&Judge.Judge).Error
			if err != nil {
				return
			}
			var answer string
			err = global.GVA_DB.Table("les_questionBank_judge").Select("is_right").Where("id = ?", studentPaper[i].QuestionId).Scan(&answer).Error
			if err != nil {
				return
			}
			judgeOrder++
			examPaper.JudgeComponent = append(examPaper.JudgeComponent, Judge)
			examPaper.JudgeComponent[judgeCount].MergeId = studentPaper[i].ID
			examPaper.JudgeComponent[judgeCount].Order = fmt.Sprintf("%d.", judgeOrder)
			var a float64
			a = float64(*studentPaper[i].Score)
			examPaper.JudgeComponent[judgeCount].Score = &a
			examPaper.JudgeComponent[judgeCount].CorrectAnswer = answer
			judgeCount++
		} else if *studentPaper[i].QuestionType == questionType.SUPPLY_BLANK {
			var Blank response.BlankComponent2
			err = global.GVA_DB.Table("les_questionBank_supply_blank").Where("id = ?", studentPaper[i].QuestionId).Find(&Blank.Blank).Error
			if err != nil {
				return
			}
			var answer string
			err = global.GVA_DB.Table("les_questionBank_supply_blank").Select("answer").Where("id = ?", studentPaper[i].QuestionId).Scan(&answer).Error
			if err != nil {
				return
			}
			blankOrder++
			examPaper.BlankComponent = append(examPaper.BlankComponent, Blank)
			examPaper.BlankComponent[blankCount].MergeId = studentPaper[i].ID
			examPaper.BlankComponent[blankCount].Order = fmt.Sprintf("%d.", blankOrder)
			var a float64
			a = float64(*studentPaper[i].Score)
			examPaper.BlankComponent[blankCount].Score = &a
			examPaper.BlankComponent[blankCount].CorrectAnswer = answer
			blankCount++
		} else if *studentPaper[i].QuestionType == questionType.PROGRAM {
			var Program response.ProgramComponent2
			var pr questionBankPo.Program
			err = global.GVA_DB.Table("les_questionBank_programm").Where("id = ?", studentPaper[i].QuestionId).Find(&pr).Error
			if err != nil {
				return
			}
			Program.Program.PracticeModel.GVA_MODEL = pr.GVA_MODEL
			Program.Program.PracticeModel.SerNo = pr.SerNo
			Program.Program.PracticeModel.Describe = pr.Describe
			Program.Program.PracticeModel.Title = pr.Title
			Program.Program.PracticeModel.ProblemType = int(pr.ProblemType)
			programOrder++
			examPaper.ProgramComponent = append(examPaper.ProgramComponent, Program)
			examPaper.ProgramComponent[programCount].MergeId = studentPaper[i].ID
			examPaper.ProgramComponent[programCount].Order = fmt.Sprintf("%d.", programOrder)
			var a float64
			a = float64(*studentPaper[i].Score)
			examPaper.ProgramComponent[programCount].Score = &a
			if pr.ReferenceAnswers != nil {
				examPaper.ProgramComponent[programCount].CorrectAnswer = *pr.ReferenceAnswers
			}
			programCount++
		} else if *studentPaper[i].QuestionType == questionType.Target {
			var target response.STargetComponent
			err = global.GVA_DB.Table("les_questionBank_target").Where("id = ?", studentPaper[i].QuestionId).Find(&target.Target).Error
			if err != nil {
				return
			}
			if err != nil {
				return
			}
			targetOrder++
			examPaper.TargetComponent = append(examPaper.TargetComponent, target)
			examPaper.TargetComponent[targetCount].MergeId = studentPaper[i].ID
			examPaper.TargetComponent[targetCount].Order = fmt.Sprintf("%d.", targetOrder)
			var a float64
			a = float64(*studentPaper[i].Score)
			examPaper.TargetComponent[targetCount].Score = &a
			targetCount++
		}
	}
	examPaper.PaperId = id
	return
}

// GetExamPaperInfoList 分页获取ExamPaper记录
// Author [piexlmax](https://github.com/piexlmax)
func (examPaperService *ExamPaperService) GetExamPaperInfoList(info examManageReq.ExamPaperSearch, userId uint, authorityID uint) (list []examManage.ExamPaper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamPaper{})
	if authorityID != 888 {
		db = db.Where("user_id = ?", userId)
	}
	var examPapers []examManage.ExamPaper
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.TermId != 0 {
		db = db.Where("term_id = ?", info.TermId)
	}
	if info.LessonId != 0 {
		db = db.Where("lesson_id = ?", info.LessonId)
	}
	if info.PlanId != 0 {
		db = db.Where("plan_id = ?", info.PlanId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc,updated_at desc ").Limit(limit).Offset(offset).Find(&examPapers).Error
	return examPapers, total, err
}
func (examPaperService *ExamPaperService) GetTemplate(info examManage.ExamPaper) (list []examManage.PaperTemplateItem, err error) {
	err = global.GVA_DB.Where("template_id = ?", info.TemplateId).Find(&list).Error
	return
}

// 该考试计划是否已经分发试卷
func (examPaperService *ExamPaperService) GetPlanStatus(PlanId uint) (status bool, err error) {
	err = global.GVA_DB.Table("tea_examplan").Select("is_distributed").Where("id = ?", PlanId).Scan(&status).Error
	return
}
func (examPaperService *ExamPaperService) GetPaperNum(PlanId uint) (number []int64, err error) {
	err = global.GVA_DB.Table("exam_paper").Select("id").Where("plan_id = ?", PlanId).Scan(&number).Error
	if err != nil {
		return nil, err
	}
	return
}
func (examPaperService *ExamPaperService) FindLateJoinStd(pid uint) (diffArray []int64, err error) {
	var students []int64
	err = global.GVA_DB.Raw("SELECT student_id FROM bas_student_teach_classes as b\njoin tea_examplan as t on t.teach_class_id = b.teach_class_id and t.id =?", pid).Scan(&students).Error
	var nowStds []int64
	err = global.GVA_DB.Raw("SELECT student_id from exam_student_paper where plan_id = ? group by student_id", pid).Scan(&nowStds).Error
	diffArray = utils.DiffArray(students, nowStds)
	return
}
func (examPaperService *ExamPaperService) LateStdsDistribution(PlanId uint, studentList []int64, number []int64) (err error) {
	global.GVA_DB.Table("tea_examplan").Where("id = ?", PlanId).Update("is_distributed", 1)
	rand.Seed(time.Now().UnixNano())
	for _, v := range studentList {
		a := rand.Intn(len(number))
		var result examManage.ExamPaper
		global.GVA_DB.Raw("INSERT INTO exam_student_paper(student_id,plan_id,question_id,score,question_type,problem_type,paper_id) SELECT student_id,tea_examplan.id,question_id,score,question_type,problem_type,paper_id from bas_student_teach_classes,exam_paper_question_merge,tea_examplan WHERE paper_id = ? and student_id = ? and tea_examplan.id = ? and exam_paper_question_merge.deleted_at is null GROUP BY student_id,tea_examplan.id,question_id,score,question_type,problem_type,paper_id", number[a], v, PlanId).Scan(&result)
		var res1 []examManage.ExamStudentPaper
		global.GVA_DB.Raw("UPDATE exam_student_paper SET got_score = 0 and created_at = NOW() and updated_at =NOW() where student_id = ? and plan_id = ?", v, PlanId).Scan(&res1)
	}
	return
}
func (examPaperService *ExamPaperService) PaperDistribution(PlanId uint, number []int64) (err error) {
	var studentList []int64
	global.GVA_DB.Table("tea_examplan").Where("id = ?", PlanId).Update("is_distributed", 1)
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = global.GVA_DB.Raw("SELECT student_id FROM bas_student_teach_classes join tea_examplan on  tea_examplan.teach_class_id = bas_student_teach_classes.teach_class_id and tea_examplan.id = ?  GROUP BY student_id ", PlanId).
			Scan(&studentList).Error
		if err != nil {
			return err
		}
		return nil
	})
	rand.Seed(time.Now().UnixNano())
	for _, v := range studentList {
		a := rand.Intn(len(number))
		var result examManage.ExamPaper
		global.GVA_DB.Raw("INSERT INTO exam_student_paper(student_id,plan_id,question_id,score,question_type,problem_type,paper_id) SELECT student_id,tea_examplan.id,question_id,score,question_type,problem_type,paper_id from bas_student_teach_classes,exam_paper_question_merge,tea_examplan WHERE paper_id = ? and student_id = ? and tea_examplan.id = ? and exam_paper_question_merge.deleted_at is null GROUP BY student_id,tea_examplan.id,question_id,score,question_type,problem_type,paper_id", number[a], v, PlanId).Scan(&result)
		var res1 []examManage.ExamStudentPaper
		global.GVA_DB.Raw("UPDATE exam_student_paper SET got_score = 0 and created_at = NOW() and updated_at =NOW() where student_id = ? and plan_id = ?", v, PlanId).Scan(&res1)
	}
	return
}

func (examPaperService *ExamPaperService) SetPaperChoiceQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.MultipleChoice
	num := info.Num
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Raw("SELECT * FROM les_questionbank_multiple_choice where problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null ORDER BY RAND()", info.ProblemType, 1, info.ChapterId).Limit(*num).Find(&list).Error
		if err != nil {
			return err
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &Id,
						QuestionId:   &list[j].ID,
						Score:        info.Score,
						QuestionType: info.QuestionType,
						ProblemType:  info.ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	wg.Done()
	return
}
func (examPaperService *ExamPaperService) SetPaperJudgeQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.Judge
	num := info.Num
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Raw("SELECT * FROM les_questionbank_judge  where problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null ORDER BY RAND()", info.ProblemType, 1, info.ChapterId).Limit(*num).Find(&list).Error
		if err != nil {
			return err
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &Id,
						QuestionId:   &list[j].ID,
						Score:        info.Score,
						QuestionType: info.QuestionType,
						ProblemType:  info.ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	wg.Done()
	return
}
func (examPaperService *ExamPaperService) SetPaperBlankQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.SupplyBlank
	num := info.Num
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Raw("SELECT * FROM les_questionbank_supply_blank  where problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null ORDER BY RAND()", info.ProblemType, 1, info.ChapterId).Limit(*num).Find(&list).Error
		if err != nil {
			return err
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &Id,
						QuestionId:   &list[j].ID,
						Score:        info.Score,
						QuestionType: info.QuestionType,
						ProblemType:  info.ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	wg.Done()
	return
}
func (examPaperService *ExamPaperService) SetPaperProgramQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.Program
	num := info.Num
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Raw("SELECT * FROM les_questionbank_programm  where problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null ORDER BY RAND()", info.ProblemType, 1, info.ChapterId).Limit(*num).Find(&list).Error
		if err != nil {
			return err
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &Id,
						QuestionId:   &list[j].ID,
						Score:        info.Score,
						QuestionType: info.QuestionType,
						ProblemType:  info.ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	wg.Done()
	return
}
func (examPaperService *ExamPaperService) SetPaperTargetQuestion(info examManage.PaperTemplateItem, Id uint) (err error) {
	var list []questionBank.Program
	num := info.Num
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Raw("SELECT * FROM les_questionbank_target  where problem_type = ? and can_exam = ? and chapter_id =? and deleted_at is null ORDER BY RAND()", info.ProblemType, 1, info.ChapterId).Limit(*num).Find(&list).Error
		if err != nil {
			return err
		} else {
			if len(list) > 0 {
				for j := 0; j < *num; j++ {
					questionMerge := examManage.PaperQuestionMerge{
						GVA_MODEL:    global.GVA_MODEL{},
						PaperId:      &Id,
						QuestionId:   &list[j].ID,
						Score:        info.Score,
						QuestionType: info.QuestionType,
						ProblemType:  info.ProblemType,
					}
					err = global.GVA_DB.Create(&questionMerge).Error
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	wg.Done()
	return
}

func (examPaperService *ExamPaperService) SetPaperQuestion(info []examManage.PaperTemplateItem, Id uint) (err error) {
	for _, v := range info {
		if *v.QuestionType == questionType.SINGLE_CHOICE {
			wg.Add(1)
			go examPaperService.SetPaperChoiceQuestion(v, Id)
		} else if *v.QuestionType == questionType.JUDGE {
			wg.Add(1)
			go examPaperService.SetPaperJudgeQuestion(v, Id)
		} else if *v.QuestionType == questionType.SUPPLY_BLANK {
			wg.Add(1)
			go examPaperService.SetPaperBlankQuestion(v, Id)

		} else if *v.QuestionType == questionType.PROGRAM {
			wg.Add(1)
			go examPaperService.SetPaperProgramQuestion(v, Id)
		} else if *v.QuestionType == questionType.Target {
			wg.Add(1)
			go examPaperService.SetPaperTargetQuestion(v, Id)
		}
	}
	wg.Wait()
	return
}
