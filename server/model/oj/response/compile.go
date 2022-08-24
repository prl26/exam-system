package response

import "time"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 1:53

 * @Note:

 **/

//
//  compile
//  @Description: 用于编译之后的返回值，成功返回编译后的文件ID和文件的过期时间，失败返回编译失败后的错误值
//

type Compile struct {
	FileId         string    `json:"fileId"`
	ExpirationTime time.Time `json:"expirationTime"`
}
