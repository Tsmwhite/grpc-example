package main

import (
	"google.golang.org/grpc"
	"grpc/implement"
	"grpc/services"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	implement.CreateTestFile()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建服务器
	grpcServer := grpc.NewServer()
	// 注册业务实现
	services.RegisterMemberServiceServer(grpcServer, &implement.Member{})
	services.RegisterFileHandelServiceServer(grpcServer, &implement.FileRealize{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
