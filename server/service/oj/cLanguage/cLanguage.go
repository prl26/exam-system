package cLanguage

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oj"
	ojResp "github.com/flipped-aurora/gin-vue-admin/server/model/oj/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
	"github.com/flipped-aurora/gin-vue-admin/server/pb"
	"strconv"
	"time"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 12:21

 * @Note:

 **/

type CService struct {
	ExecutorClient pb.ExecutorClient
}

const GCC_PATH = "/usr/bin/gcc"
const STDOUT = "stdout"
const STDERR = "stderr"

const DEFAULT_COMPILE_CPU_TIME_LIMIT uint64 = 10000000000
const DEFAULT_COMPILE_MEMORY_TIME_LIMIT uint64 = 104857600
const DEFAULT_JUDGE_CPU_TIME_LIMIT uint64 = 10000000000
const DEFAULT_JUDGE_MEMORY_LIMIT uint64 = 104857600
const DEFAULT_CODE_NAME string = "a.c"
const DEFAULT_FILE_NAME string = "a"
const FILE_FAILED_DURATION time.Duration = 5 * time.Second

//
// Compile
//  @Description:  虚假的编译接口，给上层进行使用的，会在默认的文件过期时间过后删除文件
//  @receiver c*CService
//  @param code
//  @return string
//  @return *time.Time
//  @return error
//
func (c *CService) Compile(code string) (string, *time.Time, error) {
	fileID, err := c.compile(code)
	if err != nil {
		return "", nil, err
	}
	failedTime := time.Now().Add(FILE_FAILED_DURATION)
	go func() {
		after := time.After(FILE_FAILED_DURATION)
		<-after
		err := c.Delete(fileID)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return
		}
	}()
	return fileID, &failedTime, nil
}

//
// compile
//  @Description:  真正的编译方法，返回编译之后 FileID
//  @receiver c
//  @param code
//  @return string
//  @return error
//
func (c *CService) compile(code string) (string, error) {
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
	result := exec.GetResults()[0]
	if result.Status != pb.Response_Result_Accepted {
		//说明出现了错误
		//此数应该打日志
		return "", fmt.Errorf(string(result.Files[STDERR]))
	}
	return exec.GetResults()[0].GetFileIDs()[DEFAULT_FILE_NAME], nil
}

//
// Delete
//  @Description: 手动删除 FileId
//  @receiver c
//  @param id
//  @return error
//
func (c *CService) Delete(id string) error {
	_, err := c.ExecutorClient.FileDelete(context.Background(), &pb.FileID{FileID: id})
	if err != nil {
		return err
	}
	return nil
}

//
// Judge
//  @Description: 进行程序判断，输入FileID与所有用例，返回分数结果
//  @receiver c
//  @param fileId
//  @param cases
//  @return []*ojResp.Submit
//  @return error
//
func (c *CService) Judge(fileId string, cases []*questionBank.ProgrammCase) ([]*ojResp.Submit, error) {
	n := len(cases)
	submits := make([]*ojResp.Submit, n)
	cmds := make([]*pb.Request_CmdType, n)
	for i, programmCase := range cases {
		cmds[i] = makeCmd(fileId, programmCase.Input, &programmCase.ProgrammLimit)
	}
	exec, err := c.ExecutorClient.Exec(context.Background(), &pb.Request{
		Cmd: cmds,
	})
	if err != nil {
		return nil, err
	}
	results := exec.GetResults()
	for i, result := range results {
		submits[i] = &ojResp.Submit{Name: cases[i].Name, Score: 0, ExecuteSituation: oj.ExecuteSituation{
			ResultStatus: result.Status.String(), ExitStatus: int(result.ExitStatus), Time: uint(result.Time), Memory: uint(result.Memory), Runtime: uint(result.RunTime)},
		}
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

//
// Execute
//  @Description: 通过FileID 和程序输入并携带着限制结果进行程序的执行
//  @receiver c
//  @param fileId
//  @param input
//  @param programmLimit
//  @return string
//  @return *oj.ExecuteSituation
//  @return error
//
func (c *CService) Execute(fileId string, input string, programmLimit *questionBank.ProgrammLimit) (string, *oj.ExecuteSituation, error) {
	cmd := makeCmd(fileId, input, programmLimit)
	result, err := c.ExecutorClient.Exec(context.Background(), &pb.Request{
		Cmd: []*pb.Request_CmdType{
			cmd,
		},
	})
	if err != nil {
		return "", nil, err
	}
	response := result.Results[0]
	var out string
	var executeSituation *oj.ExecuteSituation = &oj.ExecuteSituation{ResultStatus: response.Status.String(), ExitStatus: int(response.ExitStatus), Time: uint(response.Time), Memory: uint(response.Memory), Runtime: uint(response.RunTime)}
	if response.Status == pb.Response_Result_Accepted {
		out = string(response.Files[STDOUT])
	}
	return out, executeSituation, nil
}

func makeCmd(fileId string, input string, programmLimit *questionBank.ProgrammLimit) *pb.Request_CmdType {
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
		Args: []string{DEFAULT_FILE_NAME},
		Files: []*pb.Request_File{{
			File: inputFile,
		}, {
			File: stout,
		}, {
			File: stderr,
		},
		},
		CopyIn: map[string]*pb.Request_File{
			DEFAULT_FILE_NAME: {
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
	if programmLimit != nil {
		cmd = cmdLimit(programmLimit, cmd)
	}
	return cmd
}

func cmdLimit(programmLimit *questionBank.ProgrammLimit, cmd *pb.Request_CmdType) *pb.Request_CmdType {
	if programmLimit.CpuLimit != nil {
		cmd.CpuTimeLimit = uint64(*programmLimit.CpuLimit)
	} else {
		cmd.CpuTimeLimit = DEFAULT_JUDGE_CPU_TIME_LIMIT
	}
	if programmLimit.MemoryLimit != nil {
		cmd.MemoryLimit = uint64(*programmLimit.MemoryLimit)
	} else {
		cmd.MemoryLimit = DEFAULT_JUDGE_MEMORY_LIMIT
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
