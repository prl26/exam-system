package request

import "github.com/prl26/exam-system/server/model/common/request"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/30 21:13

 * @Note:

 **/

type DetailFind struct {
	Id uint `json:"id" form:"id"`
}

type CourseSupportAdd struct {
	CourseSupports []struct {
		ChapterId    uint `json:"chapterId"`
		QuestionId   uint `json:"questionId"`
		QuestionType int  `json:"questionType"`
		KnowledgeId  uint `json:"knowledge_id"`
	} `json:"courseSupports"`
}

type QuestionsSupportFind struct {
	ChapterId    uint   `json:"chapterId"`
	QuestionType int    `json:"questionType"`
	Title        string `json:"title" form:"title" gorm:"column:title;comment:;"`
	ProblemType  int    `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice  *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam      *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	request.PageInfo
}
