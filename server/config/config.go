package config

type Server struct {
	JWT              JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap              Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis            Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email            Email   `mapstructure:"email" json:"email" yaml:"email"`
	System           System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha          Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	FrontDeskAddress string  `mapstructure:"frontDeskAddress"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	HuaWeiObs  HuaWeiObs  `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	AwsS3      AwsS3      `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`
	Minio      Minio      `mapstructure:"minio" json:"minio" yaml:"minio"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	HTML  Html  `mapstructure:"html" json:"html" yaml:"html"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	// Judge
	GoJudge GoJudge `mapstructure:"goJudge" json:"goJudge" yaml:"goJudge"`

	//靶场设置
	TargetConfig TargetConfig `mapstructure:"target"`
}
