<?php
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
