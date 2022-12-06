package oj

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 2:41

 * @Note:

 **/

type ExecuteSituation struct {
	ResultStatus uint `json:"resultStatus"` //  结果码
	//ResultStatusStr string `json:"resultStatusStr"` // 结果状态
	ExitStatus int  `json:"exitStatus"` // 程序返回值
	Time       uint `json:"time"`       // 程序运行 CPU 时间，单位纳秒
	Memory     uint `json:"memory"`     // 程序运行内存，单位 byte
	Runtime    uint `json:"runtime"`    // 程序运行现实时间，单位纳秒
}
