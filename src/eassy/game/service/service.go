package service

import (
	"context"
	"fmt"
	"game_framework/src/eassy/game"
	pb "game_framework/src/eassy/proto"
)

type GameService struct {
}

func (s *GameService) GetRoomList(ctx context.Context, req *pb.C2S_GetRoomList_20000) (res *pb.S2C_GetRoomList_20000, err error) {
	m := game.RoomMgr.GetRooms()
	rooms := make([]*pb.RoomBriefInfo, 0)
	for k, _ := range m {
		rooms = append(rooms, &pb.RoomBriefInfo{
			RoomId:    int32(m[k].RoomId),
			PlayerNum: int32(m[k].PlayerNums),
		})
	}
	return &pb.S2C_GetRoomList_20000{
		Ret:  0,
		Room: rooms,
	}, nil
}

func (s *GameService) CreateRoom(ctx context.Context, req *pb.C2S_CreateRoom_20001) (res *pb.S2C_CreateRoom_20001, err error) {
	roomId := game.RoomMgr.CreateRoom("landlord")
	game.RoomMgr.Join(req.Uid, roomId)
	return &pb.S2C_CreateRoom_20001{
		Ret:    0,
		RoomId: int32(roomId),
	}, nil
}

func (s *GameService) JoinRoom(ctx context.Context, req *pb.C2S_JoinRoom_20002) (res *pb.S2C_JoinRoom_20002, err error) {
	game.RoomMgr.Join(req.Uid, int(req.RoomId))
	return &pb.S2C_JoinRoom_20002{
		Ret:    0,
		RoomId: req.RoomId,
	}, nil
}

func (s *GameService) LeaveRoom(ctx context.Context, req *pb.C2S_LeaveRoom_20003) (res *pb.S2C_LeaveRoom_20003, err error) {
	//todo
	return &pb.S2C_LeaveRoom_20003{Ret: 0}, nil
}

func (s *GameService) SendMail(ctx context.Context, req *pb.MailRequest) (res *pb.MailResponse, err error) {
	fmt.Printf("邮箱:%s;发送内容:%s", req.Mail, req.Text)
	return &pb.MailResponse{
		Ok: true,
	}, nil
}
