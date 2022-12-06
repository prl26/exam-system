package exception

import "fmt"

var NotLanguageSupportError = fmt.Errorf("没有该语言支持")
var ScoreError = fmt.Errorf("用例总分不足100")

type CompileError struct {
	Msg string
}

func (c CompileError) Error() string {
	//panic("implement me")
	return c.Msg
}
