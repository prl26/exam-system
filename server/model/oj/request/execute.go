package request

import "github.com/prl26/exam-system/server/model/questionBank"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 1:57

 * @Note:

 **/

//
//  Execute
//  @Description:   用于是在编译在之后通过可执行文件的FileId获取输出的请求
//

type Execute struct {
	FileId                     string `json:"fileId"`     //文件ID
	LanguageId                 int    `json:"languageId"` //此是用于一些特殊语言的编译运行，例如 PY
	Input                      string `json:"input"`      //文件标准输入
	questionBank.ProgrammLimit        //各种限制
}
