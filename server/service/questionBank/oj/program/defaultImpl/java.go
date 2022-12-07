package defaultImpl

import (
	"context"
	"fmt"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	"github.com/prl26/exam-system/server/pb"
	"strconv"
)

const DEFAULT_JAVA_CODE_NAME string = "Main.java"
const DEFAULT_JAVA_FILE_NAME string = "Main"
const CLASS_SUFFIX string = ".class"

type JavaService struct {
	JAVAC_PATH                        string
	JAVA_PATH                         string
	DEFAULT_COMPILE_CPU_TIME_LIMIT    uint64
	DEFAULT_COMPILE_MEMORY_TIME_LIMIT uint64
	DEFAULT_JUDGE_CPU_TIME_LIMI       uint64
	DEFAULT_JUDGE_MEMORY_LIMIT        uint64
}

func (c *JavaService) compile(client pb.ExecutorClient, code string) (string, error) {
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
		Env:  []string{"GOCACHE=/tmp/go-build"},
		Args: []string{c.JAVAC_PATH, DEFAULT_JAVA_CODE_NAME},
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
		ProcLimit:    100,
		CopyIn: map[string]*pb.Request_File{
			DEFAULT_JAVA_CODE_NAME: input,
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
				Name: DEFAULT_JAVA_FILE_NAME + CLASS_SUFFIX,
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
		return "", fmt.Errorf("compile:%s", string(result.Files[STDERR]))
	}
	return exec.GetResults()[0].GetFileIDs()[DEFAULT_JAVA_FILE_NAME+CLASS_SUFFIX], nil
}

func (c *JavaService) makeExecuteCmd(fileId string, input string, programmLimit questionBankBo.LanguageLimit) *pb.Request_CmdType {
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
		Args: []string{c.JAVA_PATH, DEFAULT_JAVA_FILE_NAME},
		Files: []*pb.Request_File{{
			File: inputFile,
		}, {
			File: stout,
		}, {
			File: stderr,
		},
		},
		ProcLimit: 50,
		CopyIn: map[string]*pb.Request_File{
			DEFAULT_JAVA_FILE_NAME + CLASS_SUFFIX: {
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

func (c *JavaService) cmdLimit(programmLimit questionBankBo.LanguageLimit, cmd *pb.Request_CmdType) *pb.Request_CmdType {
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
