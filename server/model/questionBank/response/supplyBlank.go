package response

import "github.com/prl26/exam-system/server/model/questionBank"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/30 22:06

 * @Note:

 **/

type SupplyBlankDetail struct {
	SupplyBlank   questionBank.SupplyBlank `json:"supplyBlank"`
	CourseSupport []CourseSupport          `json:"courseSupport"`
}
