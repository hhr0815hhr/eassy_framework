package game

import (
	"context"
	"fmt"
	"game_framework/src/eassy/conf"
	etcdSerivce "game_framework/src/eassy/core/service/etcd"
	pb "game_framework/src/eassy/proto/game"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type gameService struct {
}

func (s *gameService) SendMail(ctx context.Context, req *pb.MailRequest) (res *pb.MailResponse, err error) {
	fmt.Printf("邮箱:%s;发送内容:%s", req.Mail, req.Text)
	return &pb.MailResponse{
		Ok: true,
	}, nil
}

func Run(port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
		return
	}
	// 创建gRPC服务器
	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &gameService{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务

	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	//etcd服务注册
	etcdSlice := make([]string, 0)
	for _, v := range conf.EtcdCfg.Etcd {
		etcdSlice = append(etcdSlice, v)
	}
	reg, err := etcdSerivce.NewService(etcdSerivce.ServiceInfo{
		Name: "game.mail",
		Ip:   "127.0.0.1:" + port, //grpc服务节点ip
	}, etcdSlice) // etcd的节点ip
	if err != nil {
		log.Fatal(err)
	}
	go reg.Start()

	if err := s.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
