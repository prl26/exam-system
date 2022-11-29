package initialize

import (
	"github.com/prl26/exam-system/server/config"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/pb"
	"github.com/prl26/exam-system/server/service"
	"google.golang.org/grpc"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/10/25 10:25

 * @Note:

 **/

func GoJudge() {
	goJudgeConfig := global.GVA_CONFIG.GoJudge
	if goJudgeConfig.Enable {
		global.GVA_LOG.Sugar().Infof("goJudge-enable 配置属性为true,将尝试正常连接 goJudge , address 地址为 %s", global.GVA_CONFIG.GoJudge.Address)
		rpcClient, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
		client := pb.NewExecutorClient(rpcClient)
		if err != nil {
			global.GVA_LOG.Error("无法正常连接 goJudge")
			panic(err)
		}
		// 依赖注入
		if goJudgeConfig.CLanguage.Enable {
			CLanguage(client, goJudgeConfig)
		}

	} else {
		global.GVA_LOG.Sugar().Info("goJudge-enable 配置属性为false,将不会连接 goJudge")
	}
}

func CLanguage(client pb.ExecutorClient, goJudgeConfig config.GoJudge) {
	service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService.ExecutorClient = client
	service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService.DEFAULT_COMPILE_CPU_TIME_LIMIT = goJudgeConfig.CLanguage.DEFAULT_JUDGE_CPU_TIME_LIMIT
	service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService.DEFAULT_JUDGE_CPU_TIME_LIMI = goJudgeConfig.CLanguage.DEFAULT_JUDGE_CPU_TIME_LIMIT
	service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService.DEFAULT_COMPILE_MEMORY_TIME_LIMIT = goJudgeConfig.CLanguage.DEFAULT_COMPILE_MEMORY_TIME_LIMIT
	service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService.DEFAULT_JUDGE_MEMORY_LIMIT = goJudgeConfig.CLanguage.DEFAULT_JUDGE_MEMORY_LIMIT
	service.ServiceGroupApp.OjServiceServiceGroup.CLanguageService.GCC_PATH = goJudgeConfig.CLanguage.GCC_PATH
}
