package config

type TargetConfig struct {
	RpcURL        string `mapstructure:"rpc-url" `
	DeployAddress string `mapstructure:"deploy-address"`
}
