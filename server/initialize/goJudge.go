package initialize

import (
	"github.com/prl26/exam-system/server/config"
	"github.com/prl26/exam-system/server/global"
	questionBankEnum "github.com/prl26/exam-system/server/model/questionBank/enum/languageType"
	"github.com/prl26/exam-system/server/pb"
	"github.com/prl26/exam-system/server/service/questionBank/oj/program"
	defaultImpl2 "github.com/prl26/exam-system/server/service/questionBank/oj/program/defaultImpl"
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
			CLanguage(client, goJudgeConfig.CLanguage)
		}
		if goJudgeConfig.GoLanguage.Enable {
			GoLanguage(client, goJudgeConfig.GoLanguage)
		}
		if goJudgeConfig.JavaLanguage.Enable {
			JavaLanguage(client, goJudgeConfig.JavaLanguage)
		}
	} else {
		global.GVA_LOG.Sugar().Info("goJudge-enable 配置属性为false,将不会连接 goJudge")
	}
}

func CLanguage(client pb.ExecutorClient, cLanguageConfig config.CLanguage) {
	program.Register(questionBankEnum.C_LANGUAGE, defaultImpl2.BuildDefaultImpl(client, &defaultImpl2.CLanguageService{
		GCC_PATH:                          cLanguageConfig.GCC_PATH,
		DEFAULT_COMPILE_MEMORY_TIME_LIMIT: cLanguageConfig.DEFAULT_COMPILE_CPU_TIME_LIMIT,
		DEFAULT_COMPILE_CPU_TIME_LIMIT:    cLanguageConfig.DEFAULT_COMPILE_CPU_TIME_LIMIT,
		DEFAULT_JUDGE_CPU_TIME_LIMI:       cLanguageConfig.DEFAULT_JUDGE_CPU_TIME_LIMIT,
		DEFAULT_JUDGE_MEMORY_LIMIT:        cLanguageConfig.DEFAULT_JUDGE_MEMORY_LIMIT,
	}))
}

func GoLanguage(client pb.ExecutorClient, goLanguageConfig config.GoLanguage) {
	program.Register(questionBankEnum.GO_LANGUAGE, defaultImpl2.BuildDefaultImpl(client, &defaultImpl2.GoLanguageService{
		GC_PATH:                           goLanguageConfig.GC_PATH,
		DEFAULT_COMPILE_MEMORY_TIME_LIMIT: goLanguageConfig.DEFAULT_COMPILE_CPU_TIME_LIMIT,
		DEFAULT_COMPILE_CPU_TIME_LIMIT:    goLanguageConfig.DEFAULT_COMPILE_CPU_TIME_LIMIT,
		DEFAULT_JUDGE_CPU_TIME_LIMI:       goLanguageConfig.DEFAULT_JUDGE_CPU_TIME_LIMIT,
		DEFAULT_JUDGE_MEMORY_LIMIT:        goLanguageConfig.DEFAULT_JUDGE_MEMORY_LIMIT,
	}))
}

func JavaLanguage(client pb.ExecutorClient, goLanguageConfig config.JavaLanguage) {
	program.Register(questionBankEnum.JAVA, defaultImpl2.BuildDefaultImpl(client, &defaultImpl2.JavaService{
		JAVA_PATH:                         goLanguageConfig.JAVA_PATH,
		JAVAC_PATH:                        goLanguageConfig.JAVAC_PATH,
		DEFAULT_COMPILE_MEMORY_TIME_LIMIT: goLanguageConfig.DEFAULT_COMPILE_CPU_TIME_LIMIT,
		DEFAULT_COMPILE_CPU_TIME_LIMIT:    goLanguageConfig.DEFAULT_COMPILE_CPU_TIME_LIMIT,
		DEFAULT_JUDGE_CPU_TIME_LIMI:       goLanguageConfig.DEFAULT_JUDGE_CPU_TIME_LIMIT,
		DEFAULT_JUDGE_MEMORY_LIMIT:        goLanguageConfig.DEFAULT_JUDGE_MEMORY_LIMIT,
	}))
}
