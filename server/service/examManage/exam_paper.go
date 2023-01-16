package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/examManage"
	examManageReq "github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
	"github.com/prl26/exam-system/server/model/teachplan"
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
			err = global.GVA_DB.Table("les_questionbank_target").Where("id = ?", Paper[i].QuestionId).Find(&Target.Target).Error
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
	for i := 0; i < len(studentList); i++ {
		a := rand.Intn(len(number))
		var result examManage.ExamPaper
		global.GVA_DB.Raw("INSERT INTO exam_student_paper(student_id,plan_id,question_id,score,question_type,problem_type,paper_id) SELECT student_id,tea_examplan.id,question_id,score,question_type,problem_type,paper_id from bas_student_teach_classes,exam_paper_question_merge,tea_examplan WHERE paper_id = ? and student_id = ? and tea_examplan.id = ? GROUP BY student_id,tea_examplan.id,question_id,score,question_type,problem_type,paper_id", number[a], studentList[i], PlanId).Scan(&result)
		global.GVA_DB.Raw("UPDATE exam_student_paper SET got_score = 0 where student_id = ? and plan_id = ?", studentList[i], PlanId)
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
