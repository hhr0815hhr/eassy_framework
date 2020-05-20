package dispatch

import (
	"context"
	"fmt"
	"game_framework/src/eassy/conf"
	etcdSerivce "game_framework/src/eassy/core/service/etcd"
	pb "game_framework/src/eassy/proto"
	"game_framework/src/eassy/service/msgService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

func Dispatch(protoId int, msg interface{}) (resp interface{}) {
	_, ok := msgService.GetMsgService().GetMsgByRouteId(protoId)
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
	switch protoPrefix {
	case 1:
		c := pb.NewUserServiceClient(conn)

		resp = userHandle(c, msgService.GetMsgService().GetMethodByRouteId(protoId), msg)
	case 2:
		c := pb.NewGameServiceClient(conn)
		resp = gameHandle(c, msgService.GetMsgService().GetMethodByRouteId(protoId), msg)
	default:
	}
	return
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
