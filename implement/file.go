package implement

import (
	"bufio"
	"fmt"
	"grpc/services"
	"io"
	"log"
	"os"
	"strconv"
)

type FileRealize struct{}

func CreateTestFile() {
	f, err := os.OpenFile("temp/test.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("open file failed, error:%s\n", err)
	}
	defer func() {
		f.Close()
	}()
	for i := 1; i < 1000; i++ {
		f.WriteString(strconv.Itoa(i) + "\n")
	}
}

func (f FileRealize) Upload(stream services.FileHandelService_UploadServer) error {
	fmt.Println("open Upload")
	ff, err := os.OpenFile("temp/save-test.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("open file failed, error %s\n", err)
	}
	buf := bufio.NewWriter(ff)
	defer func() {
		buf.Flush()
		ff.Close()
	}()
	fmt.Println("--------")
	for {
		file, err := stream.Recv()
		if err != nil {
			fmt.Println("err-1", err)
			if err == io.EOF {
				fmt.Println("err-2", err)
				stream.SendAndClose(&services.Response{
					Code: 400,
					Msg:  "send file failed",
				})
				break
			} else {
				fmt.Println("err-3", err)
				stream.SendAndClose(&services.Response{
					Code: 200,
					Msg:  "upload success",
				})
				break
			}
		}
		content := file.GetData()
		fmt.Println("content:", content)
		n, err := ff.Write(content)
		fmt.Println("n", n)
		fmt.Println("err", err)
	}
	return nil
}

func (f FileRealize) Download(request *services.Request, stream services.FileHandelService_DownloadServer) error {
	panic("implement me")
}
