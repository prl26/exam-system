package response

import (
	"github.com/prl26/exam-system/server/config"
	"github.com/prl26/exam-system/server/model/system"
)

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
type ExaFileResponse struct {
	File system.ExaFileUploadAndDownload `json:"file"`
}
