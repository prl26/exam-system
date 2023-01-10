package initialize

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/service/questionBank/oj/target"
)

func Target() {
	config := global.GVA_CONFIG.TargetConfig
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		return
	}
	//fmt.Println(client)
	target.DeployerAddr = config.DeployAddress
	target.Client = client
}
