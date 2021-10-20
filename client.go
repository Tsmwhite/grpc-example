package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/services"
	"io"
	"log"
	"os"
	"time"
)

const (
	address = "localhost:50051"
)

func upload(stream services.FileHandelService_UploadClient) {
	ff, err := os.Open("temp/test.txt")
	if err != nil {
		log.Fatalf("read file error %s\n", err)
	}
	buf := make([]byte, 1<<15)
	for {
		n, err := ff.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				stream.CloseAndRecv()
				log.Fatalf("read file error %s\n", err)
			}
		}
		fmt.Println("buf[:n]", buf[:n])
		stream.Send(&services.File{
			Data: buf[:n],
		})
	}
	stream.CloseAndRecv()
}

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	memClient := services.NewMemberServiceClient(conn)
	fileClient := services.NewFileHandelServiceClient(conn)

	// 设定请求超时时间 30s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := fileClient.Upload(ctx)
	upload(res)
	fmt.Println("res:", res)
	fmt.Println("err:", err)
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
