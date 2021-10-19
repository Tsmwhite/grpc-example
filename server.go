package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/services"
	"log"
	"net"
	"strconv"
)

const (
	port = ":50051"
)

type Member struct{}

func (m *Member) GetUserInfo(ctx context.Context, req *services.GetUserInfoRequest) (res *services.GetUserInfoResponse, err error) {
	fmt.Println("open GetUserInfo")
	res = &services.GetUserInfoResponse{
		Code: 200,
		Info: &services.Member{
			Id:   req.Id,
			Name: "小黑",
			Age:  18,
			Sex:  services.MemberSexEnum_SEX_BOY,
		},
	}
	return
}

func (m *Member) GetUserList(context.Context, *services.GetUserListRequest) (res *services.GetUserListResponse, err error) {
	fmt.Println("open GetUserList")
	var members []*services.Member
	for a := 1; a < 10; a++ {
		members = append(members, &services.Member{
			Id:   int64(a),
			Name: "小黑" + strconv.Itoa(a),
			Age:  18,
			Sex:  services.MemberSexEnum_SEX_BOY,
		})
	}
	res = &services.GetUserListResponse{
		Code: 200,
		Data: members,
	}
	return
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建服务器
	grpcServer := grpc.NewServer()
	// 注册业务实现
	services.RegisterMemberServiceServer(grpcServer, &Member{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
