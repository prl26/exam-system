package response

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 16:04

 * @Note:

 **/

type Submit struct {
	//Id string `json:"id"`		// 用例ID
	Name       string `json:"name"`       // 用例名称
	Score      int    `json:"score"`      // 用例得分
	ErrorStr   string `json:"errorStr"`   // 详细错误信息
	ExitStatus int    `json:"exitStatus"` // 程序返回值
	Time       uint   `json:"time"`       // 程序运行 CPU 时间，单位纳秒
	Memory     uint   `json:"memory"`     // 程序运行内存，单位 byte
	Runtime    uint   `json:"runtime"`    // 程序运行现实时间，单位纳秒
}
