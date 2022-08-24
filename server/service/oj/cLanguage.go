package oj

import (
	"context"
	"fmt"
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
			Name: "stdout",
			Max:  10240},
	}
	stderr := &pb.Request_File_Pipe{
		Pipe: &pb.Request_PipeCollector{
			Name: "stderr",
			Max:  10240,
		},
	}
	cmd := &pb.Request_CmdType{
		Env:  []string{"PATH=/usr/local/bin:/usr/bin:/bin"},
		Args: []string{GCC_PATH, "a.c", "-o", "a"},
		Files: []*pb.Request_File{
			{
				File: stdio,
			}, {
				File: stout,
			}, {
				File: stderr,
			},
		},
		CpuTimeLimit: 10000000000,
		MemoryLimit:  104857600,
		ProcLimit:    50,
		CopyIn: map[string]*pb.Request_File{
			"a.c": input,
		},
		CopyOut: []*pb.Request_CmdCopyOutFile{
			{
				Name: "stdout",
			}, {
				Name: "stderr",
			},
		},
		CopyOutCached: []*pb.Request_CmdCopyOutFile{
			{
				Name: "a",
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
		return "", fmt.Errorf(string(result.Files["stderr"]))
	}
	return exec.GetResults()[0].GetFileIDs()["a"], nil
}

func (c *CLanguageService) Delete(id string) error {
	_, err := c.ExecutorClient.FileDelete(context.Background(), &pb.FileID{FileID: id})
	if err != nil {
		return err
	}
	return nil
}

func (c CLanguageService) Judge(fileId string, cases []questionBank.QuestionBankProgrammCase) error {
	//
	return nil
}
