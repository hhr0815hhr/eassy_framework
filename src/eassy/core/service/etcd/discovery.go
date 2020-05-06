package etcdSerivce

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type Master struct {
	Path   string
	Nodes  map[string]*Node
	Client *clientv3.Client
}

//node is a client
type Node struct {
	State bool
	Key   string
	Info  map[string]*ServiceInfo
}

func NewMaster(endpoints []string, watchPath string) (*Master, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	master := &Master{
		Path:   watchPath,
		Nodes:  make(map[string]*Node),
		Client: cli,
	}

	go master.WatchNodes()
	return master, err
}

func GetServiceInfo(ev *clientv3.Event) *ServiceInfo {
	info := &ServiceInfo{}
	err := json.Unmarshal([]byte(ev.Kv.Value), info)
	if err != nil {
		log.Println(err)
	}
	return info
}

func (m *Master) AddNode(key string, info *ServiceInfo) {
	node, ok := m.Nodes[key]
	if !ok {
		node = &Node{
			State: true,
			Key:   key,
			Info:  map[string]*ServiceInfo{info.Name: info},
		}
		m.Nodes[node.Key] = node
	} else {
		node.Info[info.Name] = info
	}
}

func (m *Master) DeleteNode(key string, info *ServiceInfo) {
	node, ok := m.Nodes[key]
	if !ok {
		return
	}
	delete(node.Info, info.Name)
}

func (m *Master) WatchNodes() {
	rch := m.Client.Watch(context.Background(), m.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				info := GetServiceInfo(ev)
				m.AddNode(string(ev.Kv.Key), info)
			case clientv3.EventTypeDelete:
				fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				info := GetServiceInfo(ev)
				m.DeleteNode(string(ev.Kv.Key), info)
			}
		}
	}
}
