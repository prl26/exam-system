package response

import "github.com/prl26/exam-system/server/model/oj"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 16:04

 * @Note:

 **/

type Submit struct {
	//Id string `json:"id"`		// 用例ID
	Name  string `json:"name"`  // 用例名称
	Score uint   `json:"score"` // 用例得分
	oj.ExecuteSituation
}
