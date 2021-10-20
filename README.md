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
[安装包下载地址](https://github.com/protocolbuffers/protobuf/releases)

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
```go
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
```

### 6、开启客户端
```go
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
```

## 开始 PHP

### 1、下载 gRPC 扩展和 protobuf 扩展
[GRPC扩展](http://pecl.php.net/package/gRPC)
[Protobuf扩展](http://pecl.php.net/package/protobuf)

### 2、编译 proto 生成php代码
    protoc --php_out=./php/ ./protobufs/member.proto 
    protoc --php_out=./php/ ./protobufs/file.proto 

### 3、使用 composer 加载类库 grpc/grpc 、google/protobuf
```json
{
    "name": "grpc-go-php",
    "require": {
        "grpc/grpc": "^v1.3.0",
        "google/protobuf": "^v3.3.0"
    },
    "autoload":{
        "psr-4":{
            "GPBMetadata\\":"GPBMetadata/",
            "Member\\":"Member/"
        }
    }
}
```

### 4、编写 php 客户端（还未支持 file（stream））
```php
//引入 composer 的自动载加
require __DIR__ . '/vendor/autoload.php';

//用于连接 服务端
$client = new \Member\MemberClient('127.0.0.1:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);

//实例化 TestRequest 请求类
$request = new \Member\GetUserInfoRequest();
$request->setId(1);

//调用远程服务
list($response,$status) = $client->getUserInfo($request)->wait();

print_r($status);

$info = $response->getInfo();

print_r($info->getName());
```