package response

import "exam-system/model/oj"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 2:02

 * @Note:

 **/

type Execute struct {
	Output string `json:"output"` // 标准输出
	oj.ExecuteSituation
}
