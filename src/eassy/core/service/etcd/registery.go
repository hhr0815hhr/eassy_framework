package etcdSerivce

import (
	"context"
	"encoding/json"
	"errors"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type ServiceInfo struct {
	Name string
	Ip   string
}

type Service struct {
	Info    ServiceInfo
	stop    chan error
	leaseId clientv3.LeaseID
	client  *clientv3.Client
}

type IService interface {
	Start() error
	Stop()
	keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error)
	revoke() error
}

func NewService(info ServiceInfo, endpoints []string) (*Service, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Service{
		Info:   info,
		stop:   make(chan error),
		client: cli,
	}, err
}

func (s *Service) Start() error {
	ch, err := s.keepAlive()
	if err != nil {
		log.Fatal(err)
		return err
	}
	for {
		select {
		case <-s.stop:
			return err
		case <-s.client.Ctx().Done():
			return errors.New("server closed")
		case ka, ok := <-ch:
			if !ok {
				log.Println("keep alive channel closed")
				return s.revoke()
			}
			_ = ka
			//log.Printf("Recv reply from service : %s,ttl:%d", s.Info.Name, ka.TTL)
		}
	}
}

func (s *Service) Stop() {
	s.stop <- nil
}

func (s *Service) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	info := &s.Info
	key := s.Info.Name
	value, _ := json.Marshal(info)

	//min lease TTL is 5s
	resp, err := s.client.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = s.client.Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	s.leaseId = resp.ID
	return s.client.KeepAlive(context.TODO(), resp.ID)
}

func (s *Service) revoke() error {
	_, err := s.client.Revoke(context.TODO(), s.leaseId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("service:%s stop\n", s.Info.Name)
	return err
}
