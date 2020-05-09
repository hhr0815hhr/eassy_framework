package dispatch

import (
	"context"
	"fmt"
	"game_framework/src/eassy/conf"
	etcdSerivce "game_framework/src/eassy/core/service/etcd"
	pb "game_framework/src/eassy/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

var ServiceRoute map[int]string

func Dispatch(protoId int, msg []byte) {
	call, ok := ServiceRoute[protoId]
	if !ok {
		log.Fatal("非法protoID")
		return
	}
	protoPrefix := protoId / 10000
	var nodeTypeList = []string{"", "user", "game"}
	conn, err := getServiceNode(nodeTypeList[protoPrefix])
	if err != nil {
		log.Fatal(err)
		return
	}

	//实例化service客户端
	var c interface{}
	switch protoPrefix {
	case 1:
		c = pb.NewUserServiceClient(conn)
	case 2:
		c = pb.NewGameServiceClient(conn)
	default:

	}

	resp, err := c.call(context.TODO(), &pb.MailRequest{
		Mail: "qq@mail.com",
		Text: "test,test",
	})
	log.Print(resp)
}

func getServiceNode(nodeType string) (*grpc.ClientConn, error) {
	etcdSlice := make([]string, 0)
	for _, v := range conf.EtcdCfg.Etcd {
		etcdSlice = append(etcdSlice, v)
	}
	r := etcdSerivce.NewResolver(etcdSlice, nodeType)
	resolver.Register(r)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	addr := fmt.Sprintf("%s:///%s", r.Scheme(), nodeType)
	return grpc.DialContext(ctx, addr, grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithBlock())
}
