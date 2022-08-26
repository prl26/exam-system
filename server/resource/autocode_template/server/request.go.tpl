package request

import (
	"github.com/prl26/exam-system/server/model/{{.Package}}"
	"github.com/prl26/exam-system/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
