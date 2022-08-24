package oj

import (
	"github.com/flipped-aurora/gin-vue-admin/server/pb"
	"google.golang.org/grpc"
	"log"
	"testing"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 12:33

 * @Note:

 **/
var obj *CLanguageService

var errorCode = `
	#include<stdio.h>

	int main()

	{

		printf("helloworld!\n");
		return 0 
	}`

// 测试编译
func TestCompile(t *testing.T) {
	cases := []struct {
		name    string
		code    string
		success bool
	}{
		{
			name: "correctCode",
			code: `
					#include<stdio.h>
					int main(){
						printf("hello,world!\n");
						return 0;
					}`,
			success: true,
		}, {
			name: "errorCode",
			code: `
					#include<stdio.h>
					int main(){	
						printf("helloworld!\n");
						return 0 
					}`,
			success: false,
		},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			fileId, err := obj.Compile(s.code)
			if s.success && err != nil {
				log.Fatalf("Compile(%q) err: %v", s.code, err)
			}
			if !s.success && err == nil {
				log.Fatalf("Compile(%q) return true, can not get want false", s.code)
			}
			defer func() {
				err := obj.Delete(fileId)
				if err != nil {
					log.Printf("无法删除ID为%q的文件\n", fileId)
					return
				}
			}()
		})
	}
}

func TestMain(m *testing.M) {
	rpcClient, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		//panic(err)
		panic(err)
		return
	}
	client := pb.NewExecutorClient(rpcClient)
	obj = &CLanguageService{client}
	m.Run()

}
