package defaultImpl

import (
	"context"
	"fmt"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/pb"
	"strconv"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 12:21

 * @Note:

 **/

type CLanguageService struct {
	GCC_PATH                          string
	DEFAULT_COMPILE_CPU_TIME_LIMIT    uint64
	DEFAULT_COMPILE_MEMORY_TIME_LIMIT uint64
	DEFAULT_JUDGE_CPU_TIME_LIMI       uint64
	DEFAULT_JUDGE_MEMORY_LIMIT        uint64
}

const DEFAULT_C_CODE_NAME string = "a.c"
const DEFAULT_C_FILE_NAME string = "a"

var replacer = strings.NewReplacer("\n", "", " ", "", "\t", "")

// 注意此处并不是 服务启动的真实值
// 此为服务启动的默认值   根据 config文件的配置 之后会进行依赖注入 来修改上面的值

func (c *CLanguageService) compile(client pb.ExecutorClient, code string) (string, error) {
	input := &pb.Request_File{
		File: &pb.Request_File_Memory{
			Memory: &pb.Request_MemoryFile{
				Content: []byte(code)},
		},
	}
	stdio := &pb.Request_File_Memory{
		Memory: &pb.Request_MemoryFile{
			Content: []byte("")},
	}
	stout := &pb.Request_File_Pipe{
		Pipe: &pb.Request_PipeCollector{
			Name: STDOUT,
			Max:  10240},
	}
	stderr := &pb.Request_File_Pipe{
		Pipe: &pb.Request_PipeCollector{
			Name: STDERR,
			Max:  10240,
		},
	}
	cmd := &pb.Request_CmdType{
		Env:  []string{"PATH=/usr/local/bin:/usr/bin:/bin"},
		Args: []string{c.GCC_PATH, DEFAULT_C_CODE_NAME, "-o", DEFAULT_C_FILE_NAME},
		Files: []*pb.Request_File{
			{
				File: stdio,
			}, {
				File: stout,
			}, {
				File: stderr,
			},
		},
		CpuTimeLimit: c.DEFAULT_COMPILE_CPU_TIME_LIMIT,
		MemoryLimit:  c.DEFAULT_COMPILE_MEMORY_TIME_LIMIT,
		ProcLimit:    50,
		CopyIn: map[string]*pb.Request_File{
			DEFAULT_C_CODE_NAME: input,
		},
		CopyOut: []*pb.Request_CmdCopyOutFile{
			{
				Name: STDOUT,
			}, {
				Name: STDERR,
			},
		},
		CopyOutCached: []*pb.Request_CmdCopyOutFile{
			{
				Name: DEFAULT_C_FILE_NAME,
			},
		},
	}
	exec, err := client.Exec(context.Background(), &pb.Request{
		Cmd: []*pb.Request_CmdType{
			cmd,
		},
	})
	if err != nil {
		return "", err
	}
	result := exec.GetResults()[0]
	if result.Status != pb.Response_Result_Accepted {
		//说明出现了错误
		//此数应该打日志
		return "", fmt.Errorf(string(result.Files[STDERR]))
	}
	return exec.GetResults()[0].GetFileIDs()[DEFAULT_C_FILE_NAME], nil
}

func (c *CLanguageService) makeCmd(fileId string, input string, programmLimit questionBankBo.LanguageLimit) *pb.Request_CmdType {
	inputFile := &pb.Request_File_Memory{
		Memory: &pb.Request_MemoryFile{
			Content: []byte(input),
		},
	}
	stout := &pb.Request_File_Pipe{
		Pipe: &pb.Request_PipeCollector{
			Name: STDOUT,
			Max:  10240},
	}
	stderr := &pb.Request_File_Pipe{
		Pipe: &pb.Request_PipeCollector{
			Name: STDERR,
			Max:  10240,
		},
	}
	cmd := &pb.Request_CmdType{
		Env:  []string{"PATH=/usr/local/bin:/usr/bin:/bin"},
		Args: []string{DEFAULT_C_FILE_NAME},
		Files: []*pb.Request_File{{
			File: inputFile,
		}, {
			File: stout,
		}, {
			File: stderr,
		},
		},
		CopyIn: map[string]*pb.Request_File{
			DEFAULT_C_FILE_NAME: {
				File: &pb.Request_File_Cached{
					Cached: &pb.Request_CachedFile{
						FileID: fileId,
					},
				},
			},
		},
		CopyOut: []*pb.Request_CmdCopyOutFile{
			{
				Name: STDOUT,
			}, {
				Name: STDERR,
			},
		},
	}

	cmd = c.cmdLimit(programmLimit, cmd)

	return cmd
}

func (c *CLanguageService) cmdLimit(programmLimit questionBankBo.LanguageLimit, cmd *pb.Request_CmdType) *pb.Request_CmdType {
	if programmLimit.CpuLimit != nil {
		cmd.CpuTimeLimit = uint64(*programmLimit.CpuLimit)
	} else {
		cmd.CpuTimeLimit = c.DEFAULT_COMPILE_CPU_TIME_LIMIT
	}
	if programmLimit.MemoryLimit != nil {
		cmd.MemoryLimit = uint64(*programmLimit.MemoryLimit)
	} else {
		cmd.MemoryLimit = c.DEFAULT_JUDGE_MEMORY_LIMIT
	}
	if programmLimit.ProcLimit != nil {
		cmd.ProcLimit = uint64(*programmLimit.ProcLimit)
	}
	if programmLimit.CpuSetLimit != nil {
		cmd.CpuSetLimit = strconv.Itoa(*programmLimit.CpuSetLimit)
	}
	if programmLimit.StackLimit != nil {
		cmd.StackLimit = uint64(*programmLimit.StackLimit)
	}
	if programmLimit.CpuRateLimit != nil {
		cmd.CpuRateLimit = uint64(*programmLimit.CpuRateLimit)
	}
	if programmLimit.ClockLimit != nil {
		cmd.ClockTimeLimit = uint64(*programmLimit.ClockLimit)
	}
	if programmLimit.StrictMemoryLimit != nil && *programmLimit.StackLimit == 1 {
		cmd.StrictMemoryLimit = true
	}
	return cmd
}
