package grpc

import "google.golang.org/grpc"

var GRPC *grpc.Server

func init() {
	// 创建gRPC服务器
	GRPC = grpc.NewServer()

}
