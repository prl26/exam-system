package utils

import (
	"fmt"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/examManage"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service/oj"
)

var ojService oj.ServiceGroup

func ExecPapers(examPaperCommit examManage.CommitExamPaper) (err error) {
	//判断题处理
	for i := 0; i < len(examPaperCommit.JudgeCommit); i++ {
		if Bool, err := ojService.JudgeService.Check(examPaperCommit.JudgeCommit[i].QuestionId, examPaperCommit.JudgeCommit[i].Answer); err != nil {
			return err
		} else {
			if Bool == true {
				var result examManage.ExamStudentPaper
				err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score where id = ?", examPaperCommit.JudgeCommit[i].MergeId).Scan(&result).Error
				if err != nil {
					return err
				}
			}
		}
	}
	//选择题处理
	for i := 0; i < len(examPaperCommit.MultipleChoiceCommit); i++ {
		if Bool, err := ojService.MultipleChoiceService.Check(examPaperCommit.MultipleChoiceCommit[i].QuestionId, examPaperCommit.MultipleChoiceCommit[i].Answer); err != nil {
			return err
		} else {
			if Bool == true {
				var result examManage.ExamStudentPaper
				err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score  where id = ?", examPaperCommit.MultipleChoiceCommit[i].MergeId).Scan(&result).Error
				if err != nil {
					return err
				}
			}
		}
	}
	//填空题处理
	for i := 0; i < len(examPaperCommit.BlankCommit); i++ {
		if _, num, err := ojService.SupplyBlankService.Check(examPaperCommit.BlankCommit[i].QuestionId, examPaperCommit.BlankCommit[i].Answer); err != nil {
			return err
		} else {
			if num != 0 {
				var result examManage.ExamStudentPaper
				err = global.GVA_DB.Raw("UPDATE exam_student_paper SET exam_student_paper.got_score = exam_student_paper.score*"+fmt.Sprintf("%f", float64(num)/100.0)+" where id = ?", examPaperCommit.BlankCommit[i].MergeId).Scan(&result).Error
				if err != nil {
					return err
				}
			}
		}
	}
	//总分
	var sum uint
	global.GVA_DB.Raw("SELECT SUM(score) FROM exam_student_paper where student_id = ? and plan_id = ? ", examPaperCommit.StudentId, examPaperCommit.PlanId).Scan(&sum)
	global.GVA_DB.Model(&teachplan.Score{}).Where("student_id = ? and planId = ?", examPaperCommit.StudentId, examPaperCommit.PlanId).Update("exam_score", sum)
	return
}
