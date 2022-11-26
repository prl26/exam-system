package _go

import (
	"github.com/prl26/exam-system/server/initialize"
	"github.com/prl26/exam-system/server/pb"
	"google.golang.org/grpc"
	"log"
	"testing"
)

var GoOj *GoLanguageService

var RightCode1 = `func main() {
	rand.Seed(uint64(time.Now().Unix()))
	for i := 0; i < 9; i++ {
		a := rand.Intn(9)
		fmt.Println(a)
	}
}
`
var WrongCode1 = `func main() {
	timeStr := "2019-04-07 15:15:05"
	var layout = "2006-01-02 15:04:05" //转换的时间字符串带秒则 为 2006-01-02 15:04:05
	timeVal, errByTimeConver := time.ParseInLocation(layout, timeStr, time.Local)
	if errByTimeConver != nil {
		fmt.Println("TimeStr To Time Error.....", errByTimeConver)
	}
	d := Date{
		date: &timeVal,
		str:  timeStr,
	}
}
`

type CodeStruct struct {
	Name    string
	Code    string
	success bool
}

func NewClient() pb.ExecutorClient {
	rpcClient, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		panic(err)
		return nil
	}
	client := pb.NewExecutorClient(rpcClient)
	return client
}

func TestGOCompile(t *testing.T) {
	initialize.GoJudge()
	cases := []CodeStruct{
		{Name: "test1", Code: RightCode1, success: true},
		{Name: "test2", Code: WrongCode1, success: false},
	}
	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			fileId, err := GoOj.compile(v.Code)
			if err != nil && v.success {
				log.Fatalf("compile(%q),err(%v)", v.Code, err)
			} else if err == nil && !v.success {
				log.Fatalf("compile(%q) return true, but can't get wrong answer", v.Code)
			}
			defer func() {
				err := GoOj.Delete(fileId)
				if err != nil {
					log.Printf("无法删除ID为%q的文件\n", fileId)
					return
				}
			}()
		})
	}
}
