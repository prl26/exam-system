package questionBank

import "github.com/prl26/exam-system/server/global"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/14 15:02

 * @Note:

 **/

type ChapterMerge struct {
	global.GVA_MODEL
	QuestionId   uint `json:"question_id" gorm:"question_id"`
	ChapterId    uint `json:"chapter_id" gorm:"chapter_id"`
	QuestionType int  `json:"question_type" gorm:"question_type"`
}

func (ChapterMerge) TableName() string {
	return "les_questionBank_chapter_merge"
}

type ChapterMergeView struct {
	global.GVA_MODEL
	ProblemType int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam     *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
}
