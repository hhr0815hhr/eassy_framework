package center

import (
	"fmt"
	"game_framework/src/eassy/core/etcd"
	g "game_framework/src/eassy/core/grpc"
	"net"
)

func Run(nodeType string, port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
		return
	}

	//registerService(&service.GameService{})

	//etcd服务注册
	etcd.RegEtcd(nodeType, port)

	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	if err := g.GRPC.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
