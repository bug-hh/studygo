package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/studygo/day08/gRPC/proto"
	"google.golang.org/grpc"
)

// 定义空接口
type userInfoServiceServer struct {
	pb.UnimplementedUserInfoServiceServer
}

// GetUserInfo ...
func (s *userInfoServiceServer) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name

	resp = &pb.UserResponse{
		Id:    1,
		Name:  name,
		Age:   22,
		Hobby: []string{"sing", "run"},
	}
	return resp, nil

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("监听异常：%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", "8080")
	// 实例化 gRPC
	s := grpc.NewServer()
	// 在 rpc 上注册微服务
	pb.RegisterUserInfoServiceServer(s, &userInfoServiceServer{})
	// 启动服务端
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
