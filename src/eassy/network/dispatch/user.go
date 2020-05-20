package dispatch

import (
	"context"
	"errors"
	pb "game_framework/src/eassy/proto"
	"log"
)

func userHandle(cc pb.UserServiceClient, method string, msg interface{}) interface{} {
	//if !userFlag {
	//	initUserFuncMap(cc)
	//}
	var resp interface{}
	var err error
	switch method {
	case "Login":
		resp, err = cc.Login(context.TODO(), msg.(*pb.C2S_Login_10001))
	default:
		err = errors.New("非法的方法")
	}
	//resp, err := funcMap[method].(func(context.Context, interface{}, ...grpc.CallOption) (interface{}, error))(context.TODO(), msg)
	//resp, err := funcMap[method].(func(context.Context, *pb.C2S_Login_10001, ...grpc.CallOption) (*pb.S2C_Login_10001, error))(context.TODO(), msg.(*pb.C2S_Login_10001))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return resp
}
