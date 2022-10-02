package response

import "github.com/prl26/exam-system/server/model/questionBank"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/30 22:14

 * @Note:

 **/
type JudgeDetail struct {
	Judge         questionBank.Judge
	CourseSupport []CourseSupport
}
