package defaultImpl

import (
	"context"
	"github.com/prl26/exam-system/server/global"
	questionBankBo "github.com/prl26/exam-system/server/model/questionBank/bo"
	exception "github.com/prl26/exam-system/server/model/questionBank/error"
	"github.com/prl26/exam-system/server/pb"
	"github.com/prl26/exam-system/server/service/questionBank/oj/program"
	"strings"
	"time"
)

func BuildDefaultImpl(client pb.ExecutorClient, iProgram IProgram) program.IProgramService {
	return &baseService{
		ExecutorClient: client,
		IProgram:       iProgram,
	}
}

type baseService struct {
	ExecutorClient pb.ExecutorClient
	IProgram
}

const STDOUT = "stdout"
const STDERR = "stderr"

const FILE_FAILED_DURATION time.Duration = 3 * time.Minute

func (c *baseService) Compile(code string) (string, *time.Time, error) {
	fileID, err := c.compile(c.ExecutorClient, code)
	if err != nil {
		return "", nil, err
	}
	failedTime := time.Now().Add(FILE_FAILED_DURATION)
	go func() {
		after := time.After(FILE_FAILED_DURATION)
		<-after
		err := c.delete(fileID)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return
		}
	}()
	return fileID, &failedTime, nil
}

func (c *baseService) delete(id string) error {
	_, err := c.ExecutorClient.FileDelete(context.Background(), &pb.FileID{FileID: id})
	if err != nil {
		return err
	}
	return nil
}

func (c *baseService) Execute(fileId string, input string, programmLimit questionBankBo.LanguageLimit) (string, *questionBankBo.ExecuteSituation, error) {
	cmd := c.makeExecuteCmd(fileId, input, programmLimit)
	result, err := c.ExecutorClient.Exec(context.Background(), &pb.Request{
		Cmd: []*pb.Request_CmdType{
			cmd,
		},
	})
	if err != nil {
		return "", nil, exception.SandboxError
	}
	response := result.Results[0]
	var out string
	var executeSituation = &questionBankBo.ExecuteSituation{ResultStatus: uint(response.Status), ExitStatus: int(response.ExitStatus), Time: uint(response.Time), Memory: uint(response.Memory), Runtime: uint(response.RunTime)}
	if response.Status == pb.Response_Result_Accepted {
		out = string(response.Files[STDOUT])
	}
	return out, executeSituation, nil
}

func (c *baseService) Check(code string, limit questionBankBo.LanguageLimit, cases questionBankBo.ProgramCases) ([]*questionBankBo.Submit, uint, error) {
	fileID, err := c.compile(c.ExecutorClient, code)
	if err != nil {
		return nil, 0, err
	}
	go func() {
		after := time.After(FILE_FAILED_DURATION)
		<-after
		err := c.delete(fileID)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return
		}
	}()
	return c.judge(fileID, limit, cases)
}

func (c *baseService) judge(fileId string, limit questionBankBo.LanguageLimit, cases questionBankBo.ProgramCases) ([]*questionBankBo.Submit, uint, error) {
	n := len(cases)
	submits := make([]*questionBankBo.Submit, n)
	cmds := make([]*pb.Request_CmdType, n)
	for i, programmCase := range cases {
		cmds[i] = c.makeExecuteCmd(fileId, programmCase.Input, limit)
	}
	exec, err := c.ExecutorClient.Exec(context.Background(), &pb.Request{
		Cmd: cmds,
	})
	if err != nil {
		return nil, 0, err
	}
	results := exec.GetResults()
	var sum uint
	for i, result := range results {
		var score uint
		standardAnswer := strings.ReplaceAll(string(result.Files[STDOUT]), "\r\n", "\n")
		actualAnswer := strings.ReplaceAll(cases[i].Output, "\r\n", "\n")
		if result.Status == pb.Response_Result_Accepted {
			if standardAnswer != actualAnswer {
				if replacer.Replace(standardAnswer) == replacer.Replace(actualAnswer) {
					result.Status = pb.Response_Result_PartiallyCorrect
				} else {
					result.Status = pb.Response_Result_WrongAnswer
				}
			} else {
				score = cases[i].Score
				sum += cases[i].Score
			}
		}
		submits[i] = &questionBankBo.Submit{Name: cases[i].Name, Score: score, ExecuteSituation: questionBankBo.ExecuteSituation{
			ResultStatus: uint(result.Status), ExitStatus: int(result.ExitStatus), Time: uint(result.Time), Memory: uint(result.Memory), Runtime: uint(result.RunTime)},
		}
		submits[i].ActualOutput = actualAnswer
		if i == 0 {
			submits[i].AnswerOutput = standardAnswer
		} else {
			submits[i].AnswerOutput = "标准答案暂不可知"
		}
	}
	return submits, sum, nil
}
