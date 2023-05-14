package examManage

import "encoding/json"

type ProgramAnswer struct {
	Code         string `json:"code"`
	LanguageType string `json:"languageType"`
}

func (p *ProgramAnswer) Encode() string {
	marshal, _ := json.Marshal(p)
	return string(marshal)
}

func (p *ProgramAnswer) Decode(s string) {
	json.Unmarshal([]byte(s), p)
}
