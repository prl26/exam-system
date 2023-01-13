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
	global.GVA_DB.Create(&examPaper)
	return err
}
func (draftPaperService *DraftPaperService) DeleteExamPaperDraft(ids request.IdsReq) (err error) {
	return global.GVA_DB.Delete(&[]examManage.ExamPaperDraft{}, "id in ?", ids.Ids).Error

}
func (draftPaperService *DraftPaperService) UpdateExamPaperDraft(examPaper examManage.ExamPaperDraft) (err error) {
	err = global.GVA_DB.Updates(&examPaper).Error
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
func (draftPaperService *DraftPaperService) GetExamPaperDraft(id uint) (examPaper response.ExamPaperResponse, err error) {
	examPaper.BlankComponent = make([]response.BlankComponent, 0)
	examPaper.SingleChoiceComponent = make([]response.ChoiceComponent, 0)
	examPaper.MultiChoiceComponent = make([]response.ChoiceComponent, 0)
	examPaper.JudgeComponent = make([]response.JudgeComponent, 0)
	examPaper.ProgramComponent = make([]response.ProgramComponent, 0)
	var Paper []examManage.PaperQuestionMerge
	err = global.GVA_DB.Table("exam_paper_question_merge").Where("paper_id = ?", id).Find(&Paper).Error
	var singleChoiceCount, MultiChoiceCount, judgeCount, blankCount, programCount uint
	for i := 0; i < len(Paper); i++ {
		if *Paper[i].QuestionType == questionType.SINGLE_CHOICE {
			var Choice response.ChoiceComponent
			err = global.GVA_DB.Table("les_questionBank_multiple_choice").Where("id = ?", Paper[i].QuestionId).Find(&Choice.Choice).Error
			if err != nil {
				return
			}
			Choice.MergeId = Paper[i].ID
			if Choice.Choice.IsIndefinite == 1 {
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
		}
	}
	examPaper.PaperId = id
	return
}

func (draftPaperService *DraftPaperService) ConvertDraftToPaper(info request.ConvertDraft, userId uint) (PaperID uint, err error) {
	var planDetail teachplan.ExamPlan
	err = global.GVA_DB.Where("id = ?", info.PlanId).Find(&planDetail).Error
	if err != nil {
		return
	}
	var items []examManage.PaperQuestionMerge
	err = global.GVA_DB.Where("draft_paper_id = ?", info.DraftPaperId).Find(&items).Error
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
	return examPaper.ID, err
}
