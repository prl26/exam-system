package utils

import (
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/service/oj"
)

var ojService oj.ServiceGroup

func ExecPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	for i := 0; i < len(examPaperCommit.JudgeCommit); i++ {
		if Bool, err := ojService.JudgeService.Check(examPaperCommit.JudgeCommit[0].QuestionId, examPaperCommit.JudgeCommit[i].Answer); err != nil {
			return err
		} else {
			if Bool == true {
				err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ?", examPaperCommit.JudgeCommit[0].MergeId).Error
				if err != nil {
					return err
				}
			}
		}
	}

	for i := 0; i < len(examPaperCommit.MultipleChoiceCommit); i++ {
		if Bool, err := ojService.MultipleChoiceService.Check(examPaperCommit.MultipleChoiceCommit[0].QuestionId, examPaperCommit.MultipleChoiceCommit[0].Answers); err != nil {
			return err
		} else {
			if Bool == true {
				err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ?", examPaperCommit.MultipleChoiceCommit[0].MergeId).Error
				if err != nil {
					return err
				}
			}
		}
	}
	for i := 0; i < len(examPaperCommit.BlankCommit); i++ {
		if Bool, err := ojService.SupplyBlankService.Check(examPaperCommit.MultipleChoiceCommit[0].QuestionId, examPaperCommit.BlankCommit[0].Answer); err != nil {
			return err
		} else {
			if Bool == true {
				err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ?", examPaperCommit.BlankCommit[0].MergeId).Error
				if err != nil {
					return err
				}
			}
		}
	}
	return
}
