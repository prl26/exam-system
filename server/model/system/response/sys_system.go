package response

import "exam-system/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
