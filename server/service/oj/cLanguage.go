package oj

import (
	"context"
	"fmt"
	ojResp "github.com/flipped-aurora/gin-vue-admin/server/model/oj/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	"github.com/flipped-aurora/gin-vue-admin/server/pb"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 12:21

 * @Note:

 **/

type CLanguageService struct {
	ExecutorClient pb.ExecutorClient
}

const GCC_PATH = "/usr/bin/gcc"
const STDOUT = "stdout"
const STDERR = "stderr"

const DEFAULT_COMPILE_CPU_TIME_LIMIT uint64 = 10000000000
const DEFAULT_COMPILE_MEMORY_TIME_LIMIT uint64 = 104857600
const DEFAULT_JUDGE_CPU_TIME_LIMIT uint64 = 10000000000
const DEFAULT_JUDGE_MEMORY_TIME_LIMI uint64 = 104857600
const DEFAULT_CODE_NAME string = "a.c"
const DEFAULT_FILE_NAME string = "a"

func (c *CLanguageService) Compile(code string) (string, error) {
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
		Args: []string{GCC_PATH, DEFAULT_CODE_NAME, "-o", DEFAULT_FILE_NAME},
		Files: []*pb.Request_File{
			{
				File: stdio,
			}, {
				File: stout,
			}, {
				File: stderr,
			},
		},
		CpuTimeLimit: DEFAULT_COMPILE_CPU_TIME_LIMIT,
		MemoryLimit:  DEFAULT_COMPILE_MEMORY_TIME_LIMIT,
		ProcLimit:    50,
		CopyIn: map[string]*pb.Request_File{
			DEFAULT_CODE_NAME: input,
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
				Name: DEFAULT_FILE_NAME,
			},
		},
	}
	exec, err := c.ExecutorClient.Exec(context.Background(), &pb.Request{
		Cmd: []*pb.Request_CmdType{
			cmd,
		},
	})
	if err != nil {
		return "", err
	}
	exec.GetError()
	result := exec.GetResults()[0]
	if result.Status != pb.Response_Result_Accepted {
		//说明出现了错误
		//此数应该打日志
		return "", fmt.Errorf(string(result.Files[STDERR]))
	}
	return exec.GetResults()[0].GetFileIDs()[DEFAULT_FILE_NAME], nil
}

func (c *CLanguageService) Delete(id string) error {
	_, err := c.ExecutorClient.FileDelete(context.Background(), &pb.FileID{FileID: id})
	if err != nil {
		return err
	}
	return nil
}

func (c *CLanguageService) Judge(fileId string, cases []*questionBank.ProgrammCase) ([]*ojResp.Submit, error) {
	n := len(cases)
	submits := make([]*ojResp.Submit, n)
	cmds := make([]*pb.Request_CmdType, n)
	for i, programmCase := range cases {
		cmds[i] = makeCmd(fileId, programmCase)
	}
	exec, err := c.ExecutorClient.Exec(context.Background(), &pb.Request{
		Cmd: cmds,
	})
	if err != nil {
		return nil, err
	}
	results := exec.GetResults()
	for i, result := range results {
		submits[i] = &ojResp.Submit{Name: cases[i].Name, Score: 0, ResultStatus: result.Status.String(), ExitStatus: int(result.ExitStatus), Time: uint(result.Time), Memory: uint(result.Memory), Runtime: uint(result.RunTime)}
		if result.Status == pb.Response_Result_Accepted {
			if string(result.Files[STDOUT]) != cases[i].Output {
				result.Status = pb.Response_Result_WrongAnswer
			} else {
				submits[i].Score = *cases[i].Score
			}
		}
	}
	return submits, nil
}

func makeCmd(fileId string, programmCase *questionBank.ProgrammCase) *pb.Request_CmdType {
	inputFile := &pb.Request_File_Memory{
		Memory: &pb.Request_MemoryFile{
			Content: []byte(programmCase.Input),
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
		Args: []string{"a"},
		Files: []*pb.Request_File{{
			File: inputFile,
		}, {
			File: stout,
		}, {
			File: stderr,
		},
		},
		CopyIn: map[string]*pb.Request_File{
			"a": {
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
				Name: "stderr",
			},
		},
	}
	cmd.CpuTimeLimit = DEFAULT_JUDGE_CPU_TIME_LIMIT
	cmd.MemoryLimit = DEFAULT_JUDGE_MEMORY_TIME_LIMI
	cmd.ProcLimit = 50
	return cmd
}
