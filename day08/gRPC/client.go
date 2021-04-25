package main

import (
	"context"
	"fmt"

	pb "github.com/studygo/day08/gRPC/proto"
	"google.golang.org/grpc"
)

// 1、连接服务端
// 2、实例化 grpc 客户端
// 3、调用

func main() {
	// 连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println("连接失败")
		return
	}

	// 组装请求参数
	cli := pb.NewUserInfoServiceClient(conn)
	req := new(pb.UserRequest)
	req.Name = "zs"

	// 调用接口
	resp, err := cli.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常：%s\n", err)
	}
	fmt.Printf("响应结果: %v\n", resp)

}
