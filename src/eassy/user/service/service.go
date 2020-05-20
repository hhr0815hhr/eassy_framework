package service

import (
	"context"
	"fmt"
	pb "game_framework/src/eassy/proto"
)

type UserService struct {
}

func (s *UserService) Login(ctx context.Context, req *pb.C2S_Login_10001) (res *pb.S2C_Login_10001, err error) {
	fmt.Printf("accId:%d;token:%s", req.AccountId, req.Token)
	return &pb.S2C_Login_10001{
		Code: 1,
		Info: &pb.Player_Info{
			Uid:       0,
			AccountId: req.AccountId,
			Nick:      "asd",
			Level:     10,
			Coin:      110,
			Diamond:   110,
			HeadImg:   "asd",
		},
	}, nil
}
