package config

type Html struct {
	Dir      string `mapstructure:"dir" json:"dir" yaml:"dir"`
	Template string `mapstructure:"template" json:"template" yaml:"template"`
	OutPut   string `mapstructure:"outPut" json:"outPut" yaml:"outPut"`
}
