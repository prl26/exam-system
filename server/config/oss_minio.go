package config

type Minio struct {
	EndPoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	SecretAccessKey string `mapstructure:"secret-access-key" json:"secret-access-key" yaml:"secret-access-key"`
	Secure          bool   `mapstructure:"secure" json:"secure" yaml:"secure"`
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
}
