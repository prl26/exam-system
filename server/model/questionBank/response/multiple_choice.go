package response

import "github.com/prl26/exam-system/server/model/questionBank"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/30 14:17

 * @Note:

 **/

type MultipleChoiceDetail struct {
	MultipleChoice questionBank.MultipleChoice `json:"multipleChoice"`
	CourseSupport  []CourseSupport             `json:"courseSupport"`
}
