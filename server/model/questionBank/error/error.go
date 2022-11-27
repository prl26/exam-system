package exception

import "fmt"

var NotLanguageSupportError = fmt.Errorf("没有该语言支持")
var ScoreError = fmt.Errorf("用例总分不足100")
