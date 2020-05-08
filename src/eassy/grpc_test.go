package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"

	//"google.golang.org/grpc/balancer/roundrobin"
	etcdSerivce "game_framework/src/eassy/core/service/etcd"
	"google.golang.org/grpc/resolver"
	"log"
	"testing"
	//"time"
	pb "game_framework/src/eassy/proto/game"
)

func TestService_SendMail(t *testing.T) {
	r := etcdSerivce.NewResolver([]string{
		"192.168.0.115:2379",
		"192.168.0.115:22379",
		"192.168.0.115:32379",
	}, "game")
	resolver.Register(r)

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// https://github.com/grpc/grpc/blob/master/doc/naming.md
	// The gRPC client library will use the specified scheme to pick the right resolver plugin and pass it the fully qualified name string.

	addr := fmt.Sprintf("%s:///%s", r.Scheme(), "game" /*g.srv.mail经测试，这个可以随便写，底层只是取scheme对应的Build对象*/)

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(),

		// grpc.WithBalancerName(roundrobin.Name),
		//指定初始化round_robin => balancer (后续可以自行定制balancer和 register、resolver 同样的方式)
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),

		grpc.WithBlock())

	// 这种方式也行
	//conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))

	//conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	/*conn, err := grpc.Dial(
	      fmt.Sprintf("%s://%s/%s", "consul", GetConsulHost(), s.Name),
	      //不能block => blockkingPicker打开，在调用轮询时picker_wrapper => picker时若block则不进行robin操作直接返回失败
	      //grpc.WithBlock(),
	      grpc.WithInsecure(),
	      //指定初始化round_robin => balancer (后续可以自行定制balancer和 register、resolver 同样的方式)
	      grpc.WithBalancerName(roundrobin.Name),
	      //grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	  )
	  //原文链接：https://blog.csdn.net/qq_35916684/article/details/104055246*/

	if err != nil {
		panic(err)
	}

	c := pb.NewGameServiceClient(conn)

	resp, err := c.SendMail(context.TODO(), &pb.MailRequest{
		Mail: "qq@mail.com",
		Text: "test,test",
	})
	log.Print(resp)
}
