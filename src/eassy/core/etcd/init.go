package etcd

import (
	"game_framework/src/eassy/conf"
	etcdSerivce "game_framework/src/eassy/core/service/etcd"
	"game_framework/src/eassy/util"
	"log"
)

func RegEtcd(nodeType string, port string) {
	etcdSlice := make([]string, 0)
	for _, v := range conf.EtcdCfg.Etcd {
		etcdSlice = append(etcdSlice, v)
	}
	ok, serverIP := util.ServerIP()
	if !ok {
		serverIP = "127.0.0.1"
	}
	reg, err := etcdSerivce.NewService(etcdSerivce.ServiceInfo{
		Name: nodeType,
		Ip:   serverIP + ":" + port, //grpc服务节点ip
	}, etcdSlice) // etcd的节点ip
	if err != nil {
		log.Fatal(err)
	}
	go reg.Start()
}
