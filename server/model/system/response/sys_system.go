package response

import "github.com/prl26/exam-system/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
