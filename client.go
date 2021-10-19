package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/services"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	memClient := services.NewMemberServiceClient(conn)

	// 设定请求超时时间 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	getInfoResponse, err := memClient.GetUserInfo(ctx, &services.GetUserInfoRequest{Id: 1})
	if err != nil {
		log.Printf("getInfoResponse Error: %v", err)
	}

	if "" == getInfoResponse.Err {
		log.Printf("user index success: %s", getInfoResponse.Info)
	} else {
		log.Printf("user index error: %d", getInfoResponse.Err)
	}

	getMemberList, err := memClient.GetUserList(ctx, &services.GetUserListRequest{Page: 1, Size: 20})
	if "" == getMemberList.Err {
		log.Printf("user index success: %s", getMemberList.Data)
	} else {
		log.Printf("user index error: %d", getMemberList.Err)
	}
}
