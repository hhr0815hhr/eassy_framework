package service

import (
	"context"
	"fmt"
	pb "game_framework/src/eassy/proto"
)

type GameService struct {
}

func (s *GameService) SendMail(ctx context.Context, req *pb.MailRequest) (res *pb.MailResponse, err error) {
	fmt.Printf("邮箱:%s;发送内容:%s", req.Mail, req.Text)
	return &pb.MailResponse{
		Ok: true,
	}, nil
}
