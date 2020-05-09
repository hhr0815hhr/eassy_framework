package service

import (
	"context"
	"fmt"
	pb "game_framework/src/eassy/proto"
)

type UserService struct {
}

func (s *UserService) Login(ctx context.Context, req *pb.C2S_Login_10001) (res *pb.S2C_Login_10001, err error) {
	fmt.Printf("accId:%s;token:%s", req.AccountId, req.Token)
	return &pb.S2C_Login_10001{}, nil
}
