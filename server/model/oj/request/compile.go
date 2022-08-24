package request

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 1:50

 * @Note:

 **/

//
//  compile
//  @Description: 用于选定编译语言并书写代码之后所进行的编译操作请求
//

type Compile struct {
	Code       string `json:"code"`
	LanguageId int    `json:"languageId"`
}
