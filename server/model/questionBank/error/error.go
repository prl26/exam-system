package exception

import "fmt"

var NotLanguageSupportError = fmt.Errorf("没有该语言支持")
var ScoreError = fmt.Errorf("用例总分不足100")
var SandboxError = fmt.Errorf("沙盒评判系统连接异常，请联系管理员修复")

type CompileError struct {
	Msg string
}

func (c CompileError) Error() string {
	//panic("implement me")
	return c.Msg
}
