package defaultImpl

import (
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/pb"
)

// IProgram 实现 编程判题的具体实现功能的接口
type IProgram interface {
	compile(client pb.ExecutorClient, code string) (string, error)                                       // 编译
	makeCmd(fileId string, input string, programmLimit questionBankBo.LanguageLimit) *pb.Request_CmdType // 运行编译程序时构建cmd
}
