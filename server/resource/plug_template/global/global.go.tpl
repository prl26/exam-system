package global

{{- if .HasGlobal }}

import "github.com/prl26/exam-system/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}