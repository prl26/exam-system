package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/oj"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 16:04

 * @Note:

 **/

type Submit struct {
	//Id string `json:"id"`		// 用例ID
	Name  string `json:"name"`  // 用例名称
	Score int    `json:"score"` // 用例得分
	oj.ExecuteSituation
}
