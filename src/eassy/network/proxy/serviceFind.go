package proxy

import (
	etcdSerivce "game_framework/src/eassy/core/service/etcd"
	"google.golang.org/grpc/resolver"
)

func serviceFind() {
	r := etcdSerivce.NewResolver([]string{
		"192.168.0.115:2379",
		"192.168.0.115:22379",
		"192.168.0.115:32379",
	}, "game")
	resolver.Register(r)
}
