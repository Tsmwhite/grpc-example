# grpc-example

grpc学习

## 文件描述

* implement 实现grpc生成的代码接口
* php php客户端测试代码
* protobufs 存放.proto文件
* services .proto生成的rpc服务文件
* server.go grpc服务启动
* client.go go客户端测试代码

## 开始 GO

### 1、安装 protobuf 编译器 protoc
    安装包下载地址
    <https://github.com/protocolbuffers/protobuf/releases>

### 2、安装go插件库
```
    // gRPC运行时接口编解码支持库
    go get -u github.com/golang/protobuf/proto
    
    // 从 Proto文件(gRPC接口描述文件) 生成 go文件 的编译器插件
    go get -u github.com/golang/protobuf/protoc-gen-go
```
### 3、编译 proto 文件
```
    protoc --go_out=plugins=grpc:./ ./protobufs/file.proto
    
    protoc --go_out=plugins=grpc:./ ./protobufs/member.proto
   
    protoc --go_out=plugins=grpc:{生成代码输出文件} {XXX}.proto
```
### 4、根据业务实现 grpc 生成的接口
    例如当前示例中的 FileHandelServiceServer 和 MemberServiceServer
```go
    // FileHandelServiceServer is the server API for FileHandelService service.
    type FileHandelServiceServer interface {
        Upload(FileHandelService_UploadServer) error
        Download(*Request, FileHandelService_DownloadServer) error
    }

    // MemberServiceServer is the server API for MemberService service.
    type MemberServiceServer interface {
        GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
        GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
    }
```
    我们在 implement 文件夹下分别实现了这两个接口

### 5、开启服务
    server.go
    

