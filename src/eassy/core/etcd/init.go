package etcd

import (
	"game_framework/src/eassy/conf"
	etcdSerivce "game_framework/src/eassy/core/service/etcd"
	"log"
)

func RegEtcd(nodeType string, port string) {
	etcdSlice := make([]string, 0)
	for _, v := range conf.EtcdCfg.Etcd {
		etcdSlice = append(etcdSlice, v)
	}
	reg, err := etcdSerivce.NewService(etcdSerivce.ServiceInfo{
		Name: nodeType,
		Ip:   "127.0.0.1:" + port, //grpc服务节点ip
	}, etcdSlice) // etcd的节点ip
	if err != nil {
		log.Fatal(err)
	}
	go reg.Start()
}
