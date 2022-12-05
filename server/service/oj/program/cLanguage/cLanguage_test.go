package cLanguage

//
//import (
//	"fmt"
//	"github.com/prl26/exam-system/server/global"
//	"os"
//
//	"github.com/prl26/exam-system/server/pb"
//	"google.golang.org/grpc"
//	"log"
//	"testing"
//)
//
///**
//
// * @Author: AloneAtWar
//
// * @Date:   2022/8/24 12:33
//
// * @Note:
//
// **/
//var obj *CLanguageService
//
//var errorCode = `
//	#include<stdio.h>
//
//	int main()
//
//	{
//
//		printf("helloworld!\n");
//		return 0
//	}`
//
//// 测试编译
//func TestCompile(t *testing.T) {
//	cases := []struct {
//		name    string
//		code    string
//		success bool
//	}{
//		{
//			name: "correctCode",
//			code: `
//					#include<stdio.h>
//					int main(){
//						printf("hello,world!\n");
//						return 0;
//					}`,
//			success: true,
//		}, {
//			name: "errorCode",
//			code: `
//					#include<stdio.h>
//					int main(){
//						printf("helloworld!\n");
//						return 0
//					}`,
//			success: false,
//		},
//	}
//	for _, s := range cases {
//		t.Run(s.name, func(t *testing.T) {
//			fileId, err := obj.compile(s.code)
//			if s.success && err != nil {
//				log.Fatalf("compile(%q) err: %v", s.code, err)
//			}
//			if !s.success && err == nil {
//				log.Fatalf("compile(%q) return true, can not get want false", s.code)
//			}
//			defer func() {
//				err := obj.Delete(fileId)
//				if err != nil {
//					log.Printf("无法删除ID为%q的文件\n", fileId)
//					return
//				}
//			}()
//		})
//	}
//}
//
//func TestCLanguageService_Judge(t *testing.T) {
//	cases := []struct {
//		name    string
//		code    string
//		cases   []*questionBank.ProgrammCase
//		success bool
//	}{
//		{
//			name: "代码成功案例",
//			code: `
//					#include<stdio.h>
//					int main(){
//						printf("hello,world!\n");
//						return 0;
//					}`,
//			cases: []*questionBank.ProgrammCase{
//				{
//					Name:   "你好世界!",
//					Score:  12,
//					Output: "hello,world!\n",
//				},
//			},
//			success: true,
//		}, {
//			name: "非0返回",
//			code: `
//					#include<stdio.h>
//					int main(){
//						printf("hello,world!\n");
//						return 1;
//					}`,
//			cases: []*questionBank.ProgrammCase{
//				{
//					Name:   "你好世界!",
//					Score:  12,
//					Output: "hello,world!\n",
//				},
//			},
//			success: false,
//		},
//	}
//	for _, s := range cases {
//		t.Run(s.name, func(t *testing.T) {
//			fileId, _ := obj.compile(s.code)
//			defer func() {
//				err := obj.Delete(fileId)
//				if err != nil {
//					log.Printf("无法删除ID为%q的文件\n", fileId)
//					return
//				}
//			}()
//			judge, err := obj.Judge(fileId, s.cases)
//			if err != nil {
//				return
//			}
//			fmt.Println(judge)
//		})
//
//	}
//}
//func NewClient() pb.ExecutorClient {
//	rpcClient, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
//	client := pb.NewExecutorClient(rpcClient)
//	if err != nil {
//		global.GVA_LOG.Error("无法正常连接 goJudge")
//		panic(err)
//	}
//	return client
//}
//func TestMain(m *testing.M) {
//	//client := NewClient()
//	//obj = &CLanguageService{
//	//	ExecutorClient:                    client,
//	//	GC_PATH:                          "/usr/bin/gcc",
//	//	DEFAULT_COMPILE_CPU_TIME_LIMIT:    100000000000,
//	//	DEFAULT_COMPILE_MEMORY_TIME_LIMIT: 100000000000,
//	//	DEFAULT_JUDGE_CPU_TIME_LIMI:       100000000000,
//	//	DEFAULT_JUDGE_MEMORY_LIMIT:        100000000000,
//	//}
//	//var s *pb.FileID
//	//s, err := client.FileAdd(context.Background(), &pb.FileContent{
//	//	Name:    "a.c",
//	//	Content: []byte("我是aaa"),
//	//})
//	//if err != nil {
//	//	panic(err)
//	//	return
//	//}
//	//s, err = client.FileAdd(context.Background(), &pb.FileContent{
//	//	Name:    "a.c",
//	//	Content: []byte("我是bb"),
//	//})
//	//if err != nil {
//	//	panic(err)
//	//	panic(s)
//	//	return
//	//}
//	//
//	//client.FileAdd(context.Background(), &pb.FileContent{
//	//	Name:    "a.c",
//	//	Content: []byte("我是ExecutorClient"),
//	//})
//	//if err != nil {
//	//	panic(err)
//	//	return
//	//}
//	//list, err := client.FileList(context.Background(), &emptypb.Empty{})
//	//
//	//if err != nil {
//	//	panic(err)
//	//	return
//	//}
//	//log.Println(list)
//	exitcoe := m.Run()
//	os.Exit(exitcoe)
//}
