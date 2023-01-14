package examManage

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/examManage/request"
	"github.com/prl26/exam-system/server/model/examManage/response"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/model/questionBank/enum/questionType"
	"github.com/prl26/exam-system/server/model/teachplan"
)

type DraftPaperService struct {
}

func (draftPaperService *DraftPaperService) CreateExamPaperDraft(examPaper examManage.ExamPaperDraft) (err error) {
	err = global.GVA_DB.Create(&examPaper).Error
	return err
}
func (draftPaperService *DraftPaperService) DeleteExamPaperDraft(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]examManage.ExamPaperDraft{}, "id in ?", ids.Ids).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Delete(&[]examManage.DraftPaperQuestionMerge{}, "draft_paper_id in ?", ids.Ids).Error
	return err
}
func (draftPaperService *DraftPaperService) UpdateExamPaperDraft(examPaper examManage.ExamPaperDraft) (err error) {
	err = global.GVA_DB.Where("id = ?", examPaper.ID).Updates(&examPaper).Error
	for i := 0; i < len(examPaper.PaperItem); i++ {
		global.GVA_DB.Save(&examPaper.PaperItem[i])
	}
	var IdOfItems []uint
	global.GVA_DB.Model(&examManage.DraftPaperQuestionMerge{}).Select("id").Where("draft_paper_id  = ?", examPaper.ID).Find(&IdOfItems)
	set := make(map[uint]bool)
	for _, v := range examPaper.PaperItem {
		set[v.ID] = true
	}
	for _, v := range IdOfItems {
		_, ok := set[v]
		if !ok {
			global.GVA_DB.Where("id = ?", v).Delete(&examManage.DraftPaperQuestionMerge{})
		}
	}
	return err
}
func (draftPaperService *DraftPaperService) GetExamPaperDraft(id uint) (examPaper response.ExamPaperResponse1, err error) {
	examPaper.BlankComponent = make([]response.BlankComponent1, 0)
	examPaper.SingleChoiceComponent = make([]response.ChoiceComponent1, 0)
	examPaper.MultiChoiceComponent = make([]response.ChoiceComponent1, 0)
	examPaper.JudgeComponent = make([]response.JudgeComponent1, 0)
	examPaper.ProgramComponent = make([]response.ProgramComponent1, 0)
	examPaper.TargetComponent = make([]response.TargetComponent1, 0)
	var Paper []examManage.DraftPaperQuestionMerge
	err = global.GVA_DB.Table("exam_draft_paper_merge").Where("draft_paper_id = ?", id).Find(&Paper).Error
	var singleChoiceCount, MultiChoiceCount, judgeCount, blankCount, programCount, targetCount uint
	for i := 0; i < len(Paper); i++ {
		if *Paper[i].QuestionType == questionType.SINGLE_CHOICE {
			var Choice response.ChoiceComponent1
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Where("id = ?", Paper[i].QuestionId).Find(&Choice.Choice).Error
			if err != nil {
				return
			}
			Choice.MergeId = Paper[i].ID
			Choice.Score = Paper[i].Score
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
			var Judge response.JudgeComponent1
			err = global.GVA_DB.Table("les_questionBank_judge").Where("id = ?", Paper[i].QuestionId).Find(&Judge.Judge).Error
			if err != nil {
				return
			}
			examPaper.JudgeComponent = append(examPaper.JudgeComponent, Judge)
			examPaper.JudgeComponent[judgeCount].MergeId = Paper[i].ID
			examPaper.JudgeComponent[judgeCount].Score = Paper[i].Score
			judgeCount++
		} else if *Paper[i].QuestionType == questionType.SUPPLY_BLANK {
			var Blank response.BlankComponent1
			err = global.GVA_DB.Table("les_questionBank_supply_blank").Where("id = ?", Paper[i].QuestionId).Find(&Blank.Blank).Error
			if err != nil {
				return
			}
			examPaper.BlankComponent = append(examPaper.BlankComponent, Blank)
			examPaper.BlankComponent[blankCount].MergeId = Paper[i].ID
			examPaper.BlankComponent[blankCount].Score = Paper[i].Score
			blankCount++
		} else if *Paper[i].QuestionType == questionType.PROGRAM {
			var Program response.ProgramComponent1
			var program questionBankBo.ProgramPractice
			err = global.GVA_DB.Table("les_questionBank_programm").Where("id = ?", Paper[i].QuestionId).Find(&program).Error
			if err != nil {
				return
			}
			Program.Program.Convert(&program)
			examPaper.ProgramComponent = append(examPaper.ProgramComponent, Program)
			examPaper.ProgramComponent[programCount].MergeId = Paper[i].ID
			examPaper.ProgramComponent[programCount].Score = Paper[i].Score
			programCount++
		} else if *Paper[i].QuestionType == questionType.Target {
			var Target response.TargetComponent1
			err = global.GVA_DB.Table("les_questionbank_target").Where("id = ?", Paper[i].QuestionId).Find(&Target.Target).Error
			if err != nil {
				return
			}
			examPaper.TargetComponent = append(examPaper.TargetComponent, Target)
			examPaper.TargetComponent[targetCount].MergeId = Paper[i].ID
			examPaper.TargetComponent[targetCount].Score = Paper[i].Score
			targetCount++
		}
	}
	examPaper.PaperId = id
	return
}
func (draftPaperService *DraftPaperService) GetPaperDraftInfoList(info request.DraftPaperSearch, userId uint, authorityID uint) (list []examManage.ExamPaperDraft1, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&examManage.ExamPaperDraft1{})
	db = db.Where("user_id = ?", userId)
	var examPapers []examManage.ExamPaperDraft1
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
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
func (draftPaperService *DraftPaperService) ConvertDraftToPaper(info request.ConvertDraft, userId uint) (PaperID uint, err error) {
	var planDetail teachplan.ExamPlan
	err = global.GVA_DB.Where("id = ?", info.PlanId).Find(&planDetail).Error
	if err != nil {
		return
	}
	var items []examManage.PaperQuestionMerge
	err = global.GVA_DB.Table("exam_draft_paper_merge").Where("draft_paper_id = ?", info.DraftPaperId).Find(&items).Error
	if err != nil {
		return
	}
	examPaper := examManage.ExamPaper1{
		GVA_MODEL:  global.GVA_MODEL{},
		PlanId:     &info.PlanId,
		Name:       info.Name,
		TemplateId: nil,
		TermId:     *planDetail.TermId,
		LessonId:   uint(*planDetail.LessonId),
		UserId:     &userId,
		PaperItem:  items,
	}
	global.GVA_DB.Create(&examPaper)
	return examPaper.ID, err
}
func (draftPaperService *DraftPaperService) ConvertDraftCheck(info request.ConvertDraft) (IsOk bool, err error) {
	var totalScore int64
	err = global.GVA_DB.Raw("SELECT COUNT(score) FROM `exam_draft_paper_merge` where draft_paper_id = ? and deleted_at is null", info.DraftPaperId).Scan(&totalScore).Error
	if err != nil {
		return false, err
	}
	if totalScore == 100 {
		return true, nil
	}
	return false, err
}
