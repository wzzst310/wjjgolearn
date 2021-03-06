package main

import (
	"context"
	"fmt"
	pb "github.com/wzzst310/wjjgolearn/02liwenzhou/121grpc/proto"
	"google.golang.org/grpc"
	"testing"
)

// 1.连接服务端
// 2.实例gRPC客户端
// 3.调用

func TestClient(t *testing.T) {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常： %s\n", err)
	}
	defer conn.Close()
	// 2. 实例化gRPC客户端
	client := pb.NewUserInfoServiceClient(conn)
	// 3.组装请求参数
	req := new(pb.UserRequest)
	req.Name = "zs"
	// 4. 调用接口
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常  %s\n", err)
	}
	fmt.Printf("响应结果： %v\n", response)
}
