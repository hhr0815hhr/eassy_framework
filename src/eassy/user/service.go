package user

import (
	g "game_framework/src/eassy/core/grpc"
	pb "game_framework/src/eassy/proto"
	"google.golang.org/grpc/reflection"
)

func registerService(service pb.UserServiceServer) {
	// 在gRPC服务端注册服务
	pb.RegisterUserServiceServer(g.GRPC, service)
	//在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(g.GRPC)
}
