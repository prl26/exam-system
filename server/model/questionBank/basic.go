package questionBank

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 17:34

 * @Note:

 **/

type BasicModel struct {
	ProblemType *int   `json:"problemType" form:"problemType" gorm:"column:problem_type;comment:试卷难度;"`
	CanPractice *int   `json:"canPractice" form:"canPractice" gorm:"column:can_practice;comment:是否训练题目"`
	CanExam     *int   `json:"canExam" form:"canExam" gorm:"column:can_exam;comment:是否为考试题目"`
	Describe    string `json:"describe" form:"describe" gorm:"column:describe;comment:;"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;"`
}
