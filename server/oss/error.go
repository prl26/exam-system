package oss

import "fmt"

// TODO 当上传文件过大时候 返回此Error
var MemoryTooLargeError = fmt.Errorf("文件过大")
